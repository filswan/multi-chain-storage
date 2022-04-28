package service

import (
	"encoding/json"
	"fmt"
	"mime/multipart"
	"multi-chain-storage/common/constants"

	"multi-chain-storage/config"
	"multi-chain-storage/database"
	"multi-chain-storage/models"
	"multi-chain-storage/service/scheduler"
	"os"
	"path/filepath"
	"strconv"
	"sync"

	"github.com/filswan/go-swan-lib/client/ipfs"
	"github.com/filswan/go-swan-lib/client/web"
	"github.com/filswan/go-swan-lib/logs"
	libutils "github.com/filswan/go-swan-lib/utils"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
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
			filename = srcFile.Filename + strconv.Itoa(i)
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
		return nil, err
	}
	logs.GetLogger().Info("source file saved to ", srcFilepath)
	uploadMutext.Unlock()

	logs.GetLogger().Info("uploading source file to ", config.GetConfig().IpfsServer.UploadUrlPrefix)
	uploadUrl := libutils.UrlJoin(config.GetConfig().IpfsServer.UploadUrlPrefix, "api/v0/add?stream-channels=true&pin=true")
	ipfsFileHash, err := ipfs.IpfsUploadFileByWebApi(uploadUrl, srcFilepath)
	if err != nil {
		logs.GetLogger().Error(err)
		return nil, err
	}
	logs.GetLogger().Info("source file uploaded to ", config.GetConfig().IpfsServer.UploadUrlPrefix)

	ipfsUrl := libutils.UrlJoin(config.GetConfig().IpfsServer.DownloadUrlPrefix, constants.IPFS_URL_PREFIX_BEFORE_HASH, *ipfsFileHash)

	sourceFile, err := models.GetSourceFileByPayloadCid(*ipfsFileHash)
	if err != nil {
		logs.GetLogger().Error(err)
		return nil, err
	}

	currentUtcMilliSec := libutils.GetCurrentUtcSecond()
	// not uploaded by anyone yet
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
		// remove the current copy of file
		err = os.Remove(srcFilepath)
		if err != nil {
			logs.GetLogger().Error(err)
			return nil, err
		}
	}

	sourceFileUploadUuid := uuid.NewString()
	sourceFileUpload := &models.SourceFileUpload{
		SourceFileId: sourceFile.ID,
		FileType:     fileType,
		FileName:     srcFile.Filename,
		Uuid:         sourceFileUploadUuid,
		WalletId:     wallet.ID,
		Status:       constants.SOURCE_FILE_UPLOAD_STATUS_CREATED,
		DurationDay:  duration,
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

func GetSourceFileUploads(walletAddress string, fileName, orderBy string, isAscend bool, limit, offset int) ([]*models.SourceFileUploadResult, *int, error) {
	wallet, err := models.GetWalletByAddress(walletAddress, constants.WALLET_TYPE_META_MASK)
	if err != nil {
		logs.GetLogger().Error(err)
		return nil, nil, err
	}

	srcFileUploads, totalRecordCount, err := models.GetSourceFileUploads(wallet.ID, fileName, orderBy, isAscend, limit, offset)
	if err != nil {
		logs.GetLogger().Error(err)
		return nil, nil, err
	}

	for _, srcFileUpload := range srcFileUploads {
		offlineDeals, err := models.GetOfflineDealsByCarFileId(srcFileUpload.CarFileId)
		if err != nil {
			logs.GetLogger().Error(err)
			return nil, nil, err
		}
		srcFileUpload.OfflineDeals = offlineDeals

		if srcFileUpload.SourceFileUploadStatus == constants.SOURCE_FILE_UPLOAD_STATUS_CREATED {
			srcFileUpload.Status = constants.PROCESS_STATUS_WAITING_PAYMENT
		} else {
			srcFileUpload.Status = constants.PROCESS_STATUS_PROCESSING
		}
	}

	return srcFileUploads, totalRecordCount, nil
}

type SourceFileUploadDeal struct {
	DealID                   int    `json:"deal_id"`
	DealCid                  string `json:"deal_cid"`
	MessageCid               string `json:"message_cid"`
	Height                   int    `json:"height"`
	PieceCid                 string `json:"piece_cid"`
	VerifiedDeal             bool   `json:"verified_deal"`
	StoragePricePerEpoch     int    `json:"storage_price_per_epoch"`
	Signature                string `json:"signature"`
	SignatureType            string `json:"signature_type"`
	CreatedAt                int    `json:"created_at"`
	PieceSizeFormat          string `json:"piece_size_format"`
	StartHeight              int    `json:"start_height"`
	EndHeight                int    `json:"end_height"`
	Client                   string `json:"client"`
	ClientCollateralFormat   string `json:"client_collateral_format"`
	Provider                 string `json:"provider"`
	ProviderTag              string `json:"provider_tag"`
	VerifiedProvider         int    `json:"verified_provider"`
	ProviderCollateralFormat string `json:"provider_collateral_format"`
	Status                   int    `json:"status"`
	NetworkName              string `json:"network_name"`
	StoragePrice             int    `json:"storage_price"`
	IpfsUrl                  string `json:"ipfs_url"`
	FileName                 string `json:"file_name"`
	WCid                     string `json:"w_cid"`
	LockedAt                 int64  `json:"locked_at"`
	LockedFee                string `json:"locked_fee"`
	Unlocked                 bool   `json:"unlocked"`
}
type FlinkDealResult struct {
	JobRunID int `json:"jobRunID"`
	Data     struct {
		Status string `json:"status"`
		Data   struct {
			Deal SourceFileUploadDeal `json:"deal"`
		} `json:"data"`
		Result struct {
		} `json:"result"`
	} `json:"data"`
	Result struct {
	} `json:"result"`
	StatusCode int `json:"statusCode"`
}

type flinkParams struct {
	ID   int `json:"id"`
	Data struct {
		Deal    int    `json:"deal"`
		Network string `json:"network"`
	} `json:"data"`
}

func GetSourceFileUploadDeal(sourceFileUploadId int64, dealId int) (*SourceFileUploadDeal, error) {
	flinkDealResult := FlinkDealResult{}
	if dealId > 0 {
		url := config.GetConfig().FLinkUrl
		parameter := new(flinkParams)
		parameter.Data.Deal = dealId
		parameter.Data.Network = config.GetConfig().FilecoinNetwork

		response, err := web.HttpGetNoToken(url, parameter)
		if err != nil {
			logs.GetLogger().Error(err)
		} else {
			err = json.Unmarshal(response, &flinkDealResult)
			if err != nil {
				logs.GetLogger().Error(err)
			}
		}
	}

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

	transactionPay, err := models.GetTransactionBySourceFileUploadId(sourceFileUploadId)
	if err != nil {
		logs.GetLogger().Error(err)
		return nil, err
	}

	transactionUnlock, err := models.GetTransactionBySourceFileUploadId(sourceFileUploadId)
	if err != nil {
		logs.GetLogger().Error(err)
		return nil, err
	}

	sourceFileUploadDeal := &flinkDealResult.Data.Data.Deal
	sourceFileUploadDeal.IpfsUrl = sourceFile.IpfsUrl
	sourceFileUploadDeal.FileName = sourceFileUpload.FileName
	sourceFileUploadDeal.WCid = sourceFileUpload.Uuid + sourceFile.PayloadCid
	sourceFileUploadDeal.Unlocked = false

	if transactionPay != nil {
		sourceFileUploadDeal.LockedAt = transactionPay.CreateAt
		sourceFileUploadDeal.LockedFee = transactionPay.AmountLock
	}

	if transactionUnlock != nil {
		sourceFileUploadDeal.Unlocked = true
	}

	return sourceFileUploadDeal, nil
}

type DaoInfoResult struct {
	DaoName         string `json:"dao_name"`
	DaoAddress      string `json:"dao_address"`
	OrderIndex      string `json:"order_index"`
	DealId          int64  `json:"deal_id"`
	DaoPassTime     string `json:"dao_pass_time"`
	PayloadCid      string `json:"payload_cid"`
	DaoAddressEvent string `json:"dao_address_event"`
	TxHash          string `json:"tx_hash"`
	Status          string `json:"status"`
}

func GetDaoSignEventByDealId(dealId int64) ([]*DaoInfoResult, error) {
	if dealId == 0 {
		dealId = constants.FILE_BLOCK_NUMBER_MAX
	}
	finalSql := " select * from( " +
		" (select dao_name, dao_address,order_index from dao_info order by order_index asc)) as d  left  join " +
		" (select deal_id,tx_hash,dao_pass_time,if(deal_id > 0,1,2) as status,dao_address as dao_address_event,payload_cid  from event_dao_signature where deal_id = " + strconv.FormatInt(dealId, 10) + " ) as a " +
		" on d.dao_address=a.dao_address_event"

	var daoInfoResult []*DaoInfoResult
	err := database.GetDB().Raw(finalSql).Scan(&daoInfoResult).Limit(0).Offset(constants.DEFAULT_SELECT_LIMIT).Error
	if err != nil {
		logs.GetLogger().Error(err)
		return nil, err
	}
	return daoInfoResult, nil
}

func RecordMintInfo(sourceFileIploadId int64, txHash string, tokenId string, mintAddress string) (*models.SourceFileMint, error) {
	currentUtcSecond := libutils.GetCurrentUtcSecond()
	sourceFileMint := &models.SourceFileMint{
		SourceFileUploadId: sourceFileIploadId,
		NftTxHash:          txHash,
		TokenId:            tokenId,
		MintAddress:        mintAddress,
		CreateAt:           currentUtcSecond,
		UpdateAt:           currentUtcSecond,
	}

	sourceFileMint, err := models.CreateSourceFileMint(sourceFileMint)
	if err != nil {
		logs.GetLogger().Error(err)
		return nil, err
	}

	return sourceFileMint, nil
}
