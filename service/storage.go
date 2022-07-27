package service

import (
	"encoding/json"
	"fmt"
	"mime/multipart"
	"multi-chain-storage/common/constants"
	"multi-chain-storage/database"
	"net/url"
	"strings"
	"time"

	"multi-chain-storage/config"
	"multi-chain-storage/models"
	"multi-chain-storage/service/scheduler"
	"os"
	"path/filepath"
	"strconv"
	"sync"

	"github.com/filswan/go-swan-lib/client/ipfs"
	"github.com/filswan/go-swan-lib/client/web"
	libconstants "github.com/filswan/go-swan-lib/constants"
	"github.com/filswan/go-swan-lib/logs"
	libutils "github.com/filswan/go-swan-lib/utils"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/shopspring/decimal"
)

var uploadMutext sync.Mutex

type UploadResult struct {
	SourceFileUploadId int64  `json:"source_file_upload_id"`
	PayloadCid         string `json:"payload_cid"`
	IpfsUrl            string `json:"ipfs_url"`
	FileSize           int64  `json:"file_size"`
	WCid               string `json:"w_cid"`
}

func SaveFile(c *gin.Context, srcFile *multipart.FileHeader, duration, fileType int, walletAddress string) (*UploadResult, error) {
	wallet, err := models.GetWalletByAddress(walletAddress, constants.WALLET_TYPE_META_MASK)
	if err != nil {
		logs.GetLogger().Error(err)
		return nil, err
	}

	srcDir := scheduler.GetSrcDir()

	uploadMutext.Lock()
	filename := srcFile.Filename
	if libutils.IsFileExists(srcDir, filename) {
		for i := 0; ; i++ {
			filename = srcFile.Filename + "_" + strconv.Itoa(i)
			if !libutils.IsFileExists(srcDir, filename) {
				break
			}
		}
	}

	srcFilepath := filepath.Join(srcDir, filename)
	logs.GetLogger().Info("saving source file to ", srcFilepath)
	err = c.SaveUploadedFile(srcFile, srcFilepath)
	if err != nil {
		logs.GetLogger().Error(err)
		uploadMutext.Unlock()
		return nil, err
	}
	logs.GetLogger().Info("source file saved to ", srcFilepath)
	uploadMutext.Unlock()

	logs.GetLogger().Info("uploading source file ", srcFilepath, " to ", config.GetConfig().IpfsServer.UploadUrlPrefix)
	uploadUrl := libutils.UrlJoin(config.GetConfig().IpfsServer.UploadUrlPrefix, "api/v0/add?stream-channels=true&pin=true")
	ipfsFileHash, err := ipfs.IpfsUploadFileByWebApi(uploadUrl, srcFilepath)
	if err != nil {
		logs.GetLogger().Error(err)
		return nil, err
	}
	logs.GetLogger().Info("source file ", srcFilepath, " uploaded to ", config.GetConfig().IpfsServer.UploadUrlPrefix)

	ipfsUrl := libutils.UrlJoin(config.GetConfig().IpfsServer.DownloadUrlPrefix, constants.IPFS_URL_PREFIX_BEFORE_HASH, *ipfsFileHash)

	sourceFile, err := models.GetSourceFileByPayloadCid(*ipfsFileHash)
	if err != nil {
		logs.GetLogger().Error(err)
		return nil, err
	}

	currentUtcMilliSec := libutils.GetCurrentUtcSecond()

	if sourceFile == nil {
		sourceFile = &models.SourceFile{
			FileSize:    srcFile.Size,
			ResourceUri: srcFilepath,
			IpfsUrl:     ipfsUrl,
			PinStatus:   constants.IPFS_File_PINNED_STATUS,
			PayloadCid:  *ipfsFileHash,
			CreateAt:    currentUtcMilliSec,
			UpdateAt:    currentUtcMilliSec,
		}

		sourceFile, err = models.CreateSourceFile(sourceFile)
		if err != nil {
			logs.GetLogger().Error(err)
			err = os.Remove(srcFilepath)
			if err != nil {
				logs.GetLogger().Error(err)
				return nil, err
			}
			return nil, err
		}
	} else {
		if !libutils.IsFileExistsFullPath(sourceFile.ResourceUri) {
			sourceFile.ResourceUri = srcFilepath
			sourceFile.PinStatus = constants.IPFS_File_PINNED_STATUS
			sourceFile.UpdateAt = currentUtcMilliSec
			err := database.SaveOne(sourceFile)
			if err != nil {
				logs.GetLogger().Error(err)
				return nil, err
			}
		} else {
			sourceFile.PinStatus = constants.IPFS_File_PINNED_STATUS
			sourceFile.UpdateAt = currentUtcMilliSec
			err := database.SaveOne(sourceFile)
			if err != nil {
				logs.GetLogger().Error(err)
				return nil, err
			}

			if !strings.EqualFold(sourceFile.ResourceUri, srcFilepath) {
				// remove the current copy of file
				err = os.Remove(srcFilepath)
				if err != nil {
					logs.GetLogger().Error(err)
					return nil, err
				}
			}
		}
	}

	sourceFileUploadUuid := uuid.NewString()
	sourceFileUpload := &models.SourceFileUpload{
		SourceFileId: sourceFile.ID,
		FileType:     fileType,
		FileName:     srcFile.Filename,
		Uuid:         sourceFileUploadUuid,
		WalletId:     wallet.ID,
		Status:       constants.SOURCE_FILE_UPLOAD_STATUS_PENDING,
		Duration:     duration,
		PinStatus:    constants.IPFS_File_PINNED_STATUS,
		CreateAt:     currentUtcMilliSec,
		UpdateAt:     currentUtcMilliSec,
	}

	sourceFileUpload, err = models.CreateSourceFileUpload(sourceFileUpload)
	if err != nil {
		logs.GetLogger().Error(err)
		err = os.Remove(srcFilepath)
		if err != nil {
			logs.GetLogger().Error(err)
			return nil, err
		}
		return nil, err
	}

	uploadResult := &UploadResult{
		SourceFileUploadId: sourceFileUpload.Id,
		PayloadCid:         *ipfsFileHash,
		IpfsUrl:            ipfsUrl,
		FileSize:           sourceFile.FileSize,
		WCid:               sourceFileUploadUuid + *ipfsFileHash,
	}

	return uploadResult, nil
}

func GetSourceFileUploads(walletAddress string, status, fileName, orderBy, isMinted *string, isAscend bool, limit, offset *int, uploadAtStart, uploadAtEnd *int64) ([]*models.SourceFileUploadResult, *int, error) {
	wallet, err := models.GetWalletByAddress(walletAddress, constants.WALLET_TYPE_META_MASK)
	if err != nil {
		logs.GetLogger().Error(err)
		return nil, nil, err
	}

	srcFileUploads, totalRecordCount, err := models.GetSourceFileUploads(wallet.ID, status, fileName, orderBy, isMinted, isAscend, limit, offset, uploadAtStart, uploadAtEnd)
	if err != nil {
		logs.GetLogger().Error(err)
		return nil, nil, err
	}

	for _, srcFileUpload := range srcFileUploads {
		offlineDeals, err := models.GetOfflineDealOutsBySourceFileUploadId(srcFileUpload.SourceFileUploadId)
		if err != nil {
			logs.GetLogger().Error(err)
			return nil, nil, err
		}
		srcFileUpload.OfflineDeals = offlineDeals

		if srcFileUpload.Status != constants.SOURCE_FILE_UPLOAD_STATUS_PENDING &&
			srcFileUpload.Status != constants.SOURCE_FILE_UPLOAD_STATUS_REFUNDABLE &&
			srcFileUpload.Status != constants.SOURCE_FILE_UPLOAD_STATUS_REFUNDED &&
			srcFileUpload.Status != constants.SOURCE_FILE_UPLOAD_STATUS_SUCCESS {
			srcFileUpload.Status = constants.SOURCE_FILE_UPLOAD_STATUS_PROCESSING
		}
	}

	return srcFileUploads, totalRecordCount, nil
}

func DownloadSourceFileUploads(locationStr, walletAddress string, uploadAtStart, uploadAtEnd *int64) (*string, error) {
	srcFileUploads, _, err := GetSourceFileUploads(walletAddress, nil, nil, nil, nil, true, nil, nil, uploadAtStart, uploadAtEnd)
	if err != nil {
		logs.GetLogger().Error(err)
		return nil, err
	}

	contentStr := "No.,File Name,File Size,Status,Pin Status,Price,Storage Providers,Upload Time,"
	contentStr = contentStr + "Payment,Mint\n"

	e18 := decimal.NewFromFloat32(libconstants.LOTUS_PRICE_MULTIPLE_1E18)
	for i, srcFileUpload := range srcFileUploads {
		status := srcFileUpload.Status
		minerFids := ""
		for _, offlineDeal := range srcFileUpload.OfflineDeals {
			minerFids = minerFids + offlineDeal.MinerFid + ","
			if offlineDeal.Status == constants.OFFLINE_DEAL_STATUS_SUCCESS {
				status = constants.SOURCE_FILE_UPLOAD_STATUS_SUCCESS
			}
		}
		minerFids = strings.Trim(minerFids, ",")

		contentStr = contentStr + strconv.Itoa(i+1) + ","

		contentStr = contentStr + srcFileUpload.FileName + ","

		fileSize := srcFileUpload.FileSize / constants.BYTES_1GB
		if fileSize > 0 {
			contentStr = contentStr + strconv.FormatInt(fileSize, 10) + " GB,"
		} else {
			fileSize = srcFileUpload.FileSize / constants.BYTES_1MB
			if fileSize > 0 {
				contentStr = contentStr + strconv.FormatInt(fileSize, 10) + " MB,"
			} else {
				fileSize = srcFileUpload.FileSize / constants.BYTES_1KB
				if fileSize > 0 {
					contentStr = contentStr + strconv.FormatInt(fileSize, 10) + " KB,"
				} else {
					contentStr = contentStr + strconv.FormatInt(srcFileUpload.FileSize, 10) + " B,"
				}
			}
		}

		contentStr = contentStr + status + ","
		contentStr = contentStr + srcFileUpload.PinStatus + ","

		if libutils.IsStrEmpty(&srcFileUpload.PayAmount) {
			contentStr = contentStr + ","
		} else {
			price, err := decimal.NewFromString(srcFileUpload.PayAmount)
			if err != nil {
				logs.GetLogger().Error(err)
				return nil, err
			}
			price = price.Div(e18)
			contentStr = contentStr + price.String() + " USDC,"
		}

		contentStr = contentStr + "\"" + minerFids + "\"" + ","

		uploadAt := time.Unix(srcFileUpload.UploadAt, 0)
		location, err := time.LoadLocation(locationStr)
		if err != nil {
			logs.GetLogger().Error(err)
			return nil, err
		}

		contentStr = contentStr + uploadAt.In(location).Format("2006-01-02 15:04:05") + ","

		if srcFileUpload.Status == constants.SOURCE_FILE_UPLOAD_STATUS_PENDING {
			contentStr = contentStr + constants.SOURCE_FILE_UPLOAD_STATUS_PENDING
		} else {
			contentStr = contentStr + constants.SOURCE_FILE_UPLOAD_STATUS_PAID
		}
		contentStr = contentStr + ","

		if srcFileUpload.IsMinted {
			contentStr = contentStr + "Minted"
		}

		contentStr = contentStr + "\n"
	}

	return &contentStr, nil
}

type SourceFileUpload struct {
	WCid string `json:"w_cid"`
}

func GetSourceFileUpload(sourceFileUploadId int64) (*SourceFileUpload, error) {
	sourceFileUpload, err := models.GetSourceFileUploadById(sourceFileUploadId)
	if err != nil {
		logs.GetLogger().Error(err)
		return nil, err
	}

	if sourceFileUpload == nil {
		err := fmt.Errorf("source file upload:%d not exists", sourceFileUploadId)
		logs.GetLogger().Error(err)
		return nil, err
	}

	sourceFile, err := models.GetSourceFileById(sourceFileUpload.SourceFileId)
	if err != nil {
		logs.GetLogger().Error(err)
		return nil, err
	}

	sourceFileUploadOut := &SourceFileUpload{
		WCid: sourceFileUpload.Uuid + sourceFile.PayloadCid,
	}

	return sourceFileUploadOut, nil
}

type SourceFileUploadDeal struct {
	DealID                   *int    `json:"deal_id"`
	DealCid                  *string `json:"deal_cid"`
	MessageCid               *string `json:"message_cid"`
	Height                   *int    `json:"height"`
	PieceCid                 *string `json:"piece_cid"`
	VerifiedDeal             *bool   `json:"verified_deal"`
	StoragePricePerEpoch     *int    `json:"storage_price_per_epoch"`
	Signature                *string `json:"signature"`
	SignatureType            *string `json:"signature_type"`
	CreatedAt                *int    `json:"created_at"`
	PieceSizeFormat          *string `json:"piece_size_format"`
	StartHeight              *int    `json:"start_height"`
	EndHeight                *int    `json:"end_height"`
	Client                   *string `json:"client"`
	ClientCollateralFormat   *string `json:"client_collateral_format"`
	Provider                 *string `json:"provider"`
	ProviderTag              *string `json:"provider_tag"`
	VerifiedProvider         *int    `json:"verified_provider"`
	ProviderCollateralFormat *string `json:"provider_collateral_format"`
	Status                   *int    `json:"status"`
	NetworkName              *string `json:"network_name"`
	StoragePrice             *int    `json:"storage_price"`
	IpfsUrl                  string  `json:"ipfs_url"`
	FileName                 string  `json:"file_name"`
	WCid                     string  `json:"w_cid"`
	CarFilePayloadCid        string  `json:"car_file_payload_cid"`
	LockedAt                 int64   `json:"locked_at"`
	LockedFee                string  `json:"locked_fee"`
	Unlocked                 bool    `json:"unlocked"`
}
type FlinkDealResult struct {
	JobRunID string `json:"jobRunID"`
	Data     struct {
		Deal SourceFileUploadDeal `json:"deal"`
	} `json:"data"`
}

func GetSourceFileUploadDeal(sourceFileUploadId int64, dealId int64) (*SourceFileUploadDeal, []*models.DaoSignatureOut, error) {
	flinkDealResult := FlinkDealResult{}
	sourceFileUploadDeal := &flinkDealResult.Data.Deal
	if dealId > 0 {
		flinkUrl := libutils.UrlJoin(config.GetConfig().FlinkUrl, strconv.FormatInt(dealId, 10))
		flinkUrl = flinkUrl + "?network=" + config.GetConfig().FilecoinNetwork
		params := url.Values{}
		response, err := web.HttpGetNoToken(flinkUrl, strings.NewReader(params.Encode()))
		if err != nil {
			logs.GetLogger().Error(err)
		} else {
			err = json.Unmarshal(response, &flinkDealResult)
			if err != nil {
				logs.GetLogger().Error(err)
			}
		}

		offlineDeal, err := models.GetOfflineDealByDealId(dealId)
		if err != nil {
			logs.GetLogger().Error(err)
			return nil, nil, err
		}

		if offlineDeal != nil {
			if libutils.IsStrEmpty(sourceFileUploadDeal.DealCid) {
				sourceFileUploadDeal.DealCid = &offlineDeal.DealCid
			}

			sourceFileUploadDeal.Unlocked = offlineDeal.Status == constants.OFFLINE_DEAL_STATUS_SUCCESS
		}
	}

	sourceFileUpload, err := models.GetSourceFileUploadById(sourceFileUploadId)
	if err != nil {
		logs.GetLogger().Error(err)
		return nil, nil, err
	}

	if sourceFileUpload == nil {
		err := fmt.Errorf("source file upload:%d not exists", sourceFileUploadId)
		logs.GetLogger().Error(err)
		return nil, nil, err
	}

	sourceFile, err := models.GetSourceFileById(sourceFileUpload.SourceFileId)
	if err != nil {
		logs.GetLogger().Error(err)
		return nil, nil, err
	}

	transactionPay, err := models.GetTransactionBySourceFileUploadId(sourceFileUploadId)
	if err != nil {
		logs.GetLogger().Error(err)
		return nil, nil, err
	}

	if sourceFileUpload.PinStatus == constants.IPFS_File_PINNED_STATUS {
		sourceFileUploadDeal.IpfsUrl = sourceFile.IpfsUrl
	} else {
		sourceFileUploadDeal.IpfsUrl = ""
	}

	sourceFileUploadDeal.FileName = sourceFileUpload.FileName
	sourceFileUploadDeal.WCid = sourceFileUpload.Uuid + sourceFile.PayloadCid

	if transactionPay != nil {
		sourceFileUploadDeal.LockedAt = transactionPay.CreateAt
		sourceFileUploadDeal.LockedFee = transactionPay.PayAmount
	}

	carFile, err := models.GetCarFileBySourceFileUploadId(sourceFileUploadId)
	if err != nil {
		logs.GetLogger().Error(err)
		return nil, nil, err
	}

	if carFile != nil {
		sourceFileUploadDeal.CarFilePayloadCid = carFile.PayloadCid
	}

	daoSignatures, err := models.GetDaoSignaturesByDealId(dealId)
	if err != nil {
		logs.GetLogger().Error(err)
		return nil, nil, err
	}

	return sourceFileUploadDeal, daoSignatures, nil
}

func RecordMintInfo(sourceFileIploadId int64, txHash string, tokenId int64, mintAddress string) (*models.SourceFileMint, error) {
	/*
		ethClient, _, err := client.GetEthClient()
		if err != nil {
			logs.GetLogger().Error(err)
			return nil, err
		}*/

	mintInfo, err := models.GetSourceFileMintBySourceFileUploadId(sourceFileIploadId)
	if err != nil {
		logs.GetLogger().Error(err)
		return nil, err
	}

	if mintInfo != nil {
		err := fmt.Errorf("mint info exists for source file upload id:%d", sourceFileIploadId)
		logs.GetLogger().Error(err)
		return nil, err
	}

	currentUtcSecond := libutils.GetCurrentUtcSecond()
	sourceFileMint := &models.SourceFileMint{
		SourceFileUploadId: sourceFileIploadId,
		NftTxHash:          txHash,
		TokenId:            tokenId,
		MintAddress:        mintAddress,
		CreateAt:           currentUtcSecond,
		UpdateAt:           currentUtcSecond,
	}

	sourceFileMint, err = models.CreateSourceFileMint(sourceFileMint)
	if err != nil {
		logs.GetLogger().Error(err)
		return nil, err
	}

	return sourceFileMint, nil
}

func UnpinSourceFile(sourceFileUploadId int64) error {
	err := models.UpdateSourceFileUploadPinStatus(sourceFileUploadId, constants.IPFS_File_UNPINNED_STATUS)
	if err != nil {
		logs.GetLogger().Error(err)
		return err
	}

	sourceFileUpload, err := models.GetSourceFileUploadById(sourceFileUploadId)
	if err != nil {
		logs.GetLogger().Error(err)
		return err
	}

	sourceFile, err := models.GetSourceFileById(sourceFileUpload.SourceFileId)
	if err != nil {
		logs.GetLogger().Error(err)
		return err
	}

	sourceFileUploadsPinned, err := models.GetSourceFileUploadsBySourceFileIdPinStatus(sourceFileUpload.SourceFileId, constants.IPFS_File_PINNED_STATUS)
	if err != nil {
		logs.GetLogger().Error(err)
		return err
	}

	if len(sourceFileUploadsPinned) == 0 {
		unpinUrl := libutils.UrlJoin(config.GetConfig().IpfsServer.UploadUrlPrefix, "api/v0/pin/rm")
		unpinUrl = unpinUrl + "?arg=" + sourceFile.PayloadCid
		params := url.Values{}
		_, err = web.HttpPostNoToken(unpinUrl, strings.NewReader(params.Encode()))
		if err != nil {
			logs.GetLogger().Info(err)
			return err
		}

		err = models.UpdateSourceFilePinStatus(sourceFileUpload.SourceFileId, constants.IPFS_File_UNPINNED_STATUS)
		if err != nil {
			logs.GetLogger().Error(err)
			return err
		}
	}

	return nil
}
