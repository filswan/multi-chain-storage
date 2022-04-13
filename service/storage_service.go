package service

import (
	"mime/multipart"
	"multi-chain-storage/common/constants"
	"multi-chain-storage/common/utils"
	"multi-chain-storage/config"
	"multi-chain-storage/database"
	"multi-chain-storage/models"
	"multi-chain-storage/scheduler"
	"os"
	"path/filepath"
	"strconv"
	"sync"

	"github.com/filswan/go-swan-lib/client/ipfs"
	"github.com/filswan/go-swan-lib/logs"
	libutils "github.com/filswan/go-swan-lib/utils"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

var uploadMutext sync.Mutex

type UploadResult struct {
	SourceFileId int64  `json:"source_file_id"`
	PayloadCid   string `json:"payload_cid"`
	IpfsUrl      string `json:"ipfs_url"`
	FileSize     int64  `json:"file_size"`
	Uuid         string `json:"uuid"`
}

func SaveFile(c *gin.Context, srcFile *multipart.FileHeader, duration, fileType int, walletAddress string) (*UploadResult, error) {
	wallet, err := models.GetWalletByAddressType(walletAddress, constants.WALLET_TYPE_USER)
	if err != nil {
		logs.GetLogger().Error(err)
		return nil, err
	}

	if wallet == nil {
		wallet, err = models.SaveWallet(walletAddress, constants.WALLET_TYPE_USER)
		if err != nil {
			logs.GetLogger().Error(err)
			return nil, err
		}
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
	uploadMutext.Lock()

	uploadUrl := libutils.UrlJoin(config.GetConfig().IpfsServer.UploadUrlPrefix, "api/v0/add?stream-channels=true&pin=true")
	ipfsFileHash, err := ipfs.IpfsUploadFileByWebApi(uploadUrl, srcFilepath)
	if err != nil {
		logs.GetLogger().Error(err)
		return nil, err
	}

	ipfsUrl := libutils.UrlJoin(config.GetConfig().IpfsServer.DownloadUrlPrefix, constants.IPFS_URL_PREFIX_BEFORE_HASH, *ipfsFileHash)

	sourceFile, err := models.GetSourceFileByPayloadCid(*ipfsFileHash)
	if err != nil {
		logs.GetLogger().Error(err)
		return nil, err
	}

	currentUtcMilliSec := utils.GetCurrentUtcMilliSecond()
	// not uploaded by anyone yet
	if sourceFile == nil {
		sourceFile := &models.SourceFile{
			FileSize:    srcFile.Size,
			ResourceUri: srcFilepath,
			IpfsUrl:     ipfsUrl,
			PinStatus:   constants.IPFS_File_PINNED_STATUS,
			PayloadCid:  *ipfsFileHash,
			FileType:    fileType,
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
	sourceFileUpload := models.SourceFileUpload{
		SourceFileId: sourceFile.ID,
		FileName:     srcFile.Filename,
		WalletId:     wallet.ID,
		Status:       constants.SOURCE_FILE_UPLOAD_STATUS_CREATED,
		Uuid:         sourceFileUploadUuid,
		CreateAt:     currentUtcMilliSec,
		UpdateAt:     currentUtcMilliSec,
	}

	err = database.SaveOne(&sourceFileUpload)
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
		SourceFileId: sourceFile.ID,
		PayloadCid:   *ipfsFileHash,
		IpfsUrl:      ipfsUrl,
		FileSize:     sourceFile.FileSize,
		Uuid:         sourceFileUploadUuid,
	}

	return uploadResult, nil
}
