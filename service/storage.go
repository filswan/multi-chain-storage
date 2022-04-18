package service

import (
	"context"
	"math/big"
	"mime/multipart"
	"multi-chain-storage/common/constants"
	"multi-chain-storage/common/utils"
	"multi-chain-storage/config"
	"multi-chain-storage/database"
	"multi-chain-storage/models"
	"multi-chain-storage/on-chain/client"
	"multi-chain-storage/service/scheduler"
	"os"
	"path/filepath"
	"strconv"
	"strings"
	"sync"

	"github.com/ethereum/go-ethereum/common"
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

	currentUtcMilliSec := utils.GetCurrentUtcSecond()
	// not uploaded by anyone yet
	if sourceFile == nil {
		sourceFile = &models.SourceFile{
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

func GetSourceFiles(pageSize, offset string, walletAddress, payloadCid string, file_name string, orderByColumn int, ascdesc string) ([]*models.SourceFileExt, error) {
	srcFiles, err := models.GetSourceFiles(pageSize, offset, walletAddress, payloadCid, file_name, orderByColumn, ascdesc)
	if err != nil {
		logs.GetLogger().Error(err)
		return nil, err
	}

	dealFileIds := map[int64]bool{}

	for _, srcFile := range srcFiles {
		dealFileIds[srcFile.DealFileId] = true
		if srcFile.Duration == 0 {
			srcFile.Duration = constants.DURATION_DAYS_DEFAULT
		}

		srcFile.OfflineDeals = []*models.OfflineDeal{}
	}

	if len(dealFileIds) > 0 {
		dealFileIdList := make([]int64, 0, len(dealFileIds))
		for dealFileId := range dealFileIds {
			dealFileIdList = append(dealFileIdList, dealFileId)
		}

		offlineDeals, err := models.GetOfflineDealsByDealFileIds(dealFileIdList)
		if err != nil {
			logs.GetLogger().Error(err)
			return nil, err
		}

		for _, srcFile := range srcFiles {
			for _, offlineDeal := range offlineDeals {
				if offlineDeal.CarFileId == srcFile.DealFileId {
					srcFile.OfflineDeals = append(srcFile.OfflineDeals, offlineDeal)
				}
			}
		}
	}

	return srcFiles, nil
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

type LockFound struct {
	PayloadCid string `json:"payload_cid"`
	CreateAt   string `json:"create_at"`
	LockedFee  string `json:"locked_fee"`
}

func GetLockFoundInfoByPayloadCid(payloadCid string) (*LockFound, error) {
	lockEventList, err := models.GetEventLockPaymentBySrcPayloadCid(payloadCid)
	if err != nil {
		logs.GetLogger().Error(err)
		return nil, err
	}

	lockFound := &LockFound{
		PayloadCid: payloadCid,
	}

	if len(lockEventList) > 0 {
		lockFound.LockedFee = lockEventList[0].LockedFee.String()
		lockFound.CreateAt = strconv.FormatInt(lockEventList[0].CreateAt, 10)
	}
	return lockFound, nil
}
func GetOfflineDealsBySourceFileId(sourceFileId int64) ([]*models.OfflineDeal, *models.SourceFile, error) {
	offlineDeals, err := models.GetOfflineDealsBySourceFileId(sourceFileId)
	if err != nil {
		logs.GetLogger().Error(err)
		return nil, nil, err
	}

	sourceFile, err := models.GetSourceFileById(sourceFileId)
	if err != nil {
		logs.GetLogger().Error(err)
		return nil, nil, err
	}

	return offlineDeals, sourceFile, nil
}

func SaveExpirePaymentEvent(txHash string) (*models.EventExpirePayment, error) {

	ethClient, rpcClient, err := client.GetEthClient()

	if err != nil {
		logs.GetLogger().Error(err)
		return nil, err
	}

	if txHash != "" && strings.HasPrefix(txHash, "0x") {
		var rpcTransaction *models.RpcTransaction
		err = rpcClient.CallContext(context.Background(), &rpcTransaction, "eth_getTransactionByHash", common.HexToHash(txHash))
		if err != nil {
			logs.GetLogger().Error(err)
			return nil, err
		}

		transactionReceipt, err := ethClient.TransactionReceipt(context.Background(), common.HexToHash(txHash))
		if err != nil {
			logs.GetLogger().Error(err)
			return nil, err
		}

		event := new(models.EventExpirePayment)
		event.TxHash = txHash

		block, err := ethClient.BlockByHash(context.Background(), *rpcTransaction.BlockHash)
		if err != nil {
			logs.GetLogger().Error(err)
		} else {
			event.BlockTime = strconv.FormatUint(block.Time(), 10)
		}
		blockNumberStr := strings.Replace(*rpcTransaction.BlockNumber, "0x", "", -1)
		blockNumberInt64, err := strconv.ParseUint(blockNumberStr, 16, 64)
		if err != nil {
			logs.GetLogger().Error(err)
			return nil, err
		}
		event.BlockNo = strconv.FormatUint(blockNumberInt64, 10)
		wfilCoinId, err := models.GetCoinByName(constants.COIN_NAME_USDC)
		if err != nil {
			logs.GetLogger().Error(err)
		} else {
			event.CoinId = wfilCoinId.ID
			event.NetworkId = wfilCoinId.NetworkId
		}

		contrackABI, err := client.GetContractAbi()

		if err != nil {
			logs.GetLogger().Error(err)
		}

		for _, v := range transactionReceipt.Logs {
			if v.Topics[0].Hex() == "0xe704d5e6168e602e91f017f25d889b182d9e11a90fd939a489cc2f04734c1f8a" {
				dataList, err := contrackABI.Unpack("ExpirePayment", v.Data)
				if err != nil {
					logs.GetLogger().Error(err)
				}
				event.PayloadCid = dataList[0].(string)
				event.TokenAddress = dataList[1].(common.Address).Hex()
				event.ExpireUserAmount = dataList[2].(*big.Int).String()
				event.UserAddress = dataList[3].(common.Address).Hex()
			}
		}
		event.CreateAt = strconv.FormatInt(utils.GetCurrentUtcSecond(), 10)
		event.ContractAddress = transactionReceipt.ContractAddress.Hex()

		eventList, err := models.FindEventExpirePayments(&models.EventExpirePayment{TxHash: txHash, BlockNo: strconv.
			FormatUint(blockNumberInt64, 10)}, "id desc", "10", "0")
		if err != nil {
			logs.GetLogger().Error(err)
		}
		if len(eventList) <= 0 {
			err = database.SaveOneWithTransaction(event)
			if err != nil {
				logs.GetLogger().Error(err)
			}
		}
		return event, nil
	}
	return nil, nil
}
