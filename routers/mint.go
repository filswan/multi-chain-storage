package routers

import (
	"errors"
	"multi-chain-storage/common"
	"multi-chain-storage/common/errorinfo"
	"multi-chain-storage/common/utils"
	"multi-chain-storage/database"
	"multi-chain-storage/models"
	"net/http"

	"github.com/filswan/go-swan-lib/logs"
	"github.com/gin-gonic/gin"
)

func Mint(router *gin.RouterGroup) {
	router.POST("/mint/info", RecordMintInfo)
}

type mintInfoUpload struct {
	PayloadCid  string `json:"payload_cid"`
	TxHash      string `json:"tx_hash"`
	TokenId     string `json:"token_id"`
	MintAddress string `json:"mint_address"`
}

func RecordMintInfo(c *gin.Context) {
	var model mintInfoUpload
	c.BindJSON(&model)

	payloadCid := model.PayloadCid
	nftTxHash := model.TxHash
	tokenId := model.TokenId
	mintAddress := model.MintAddress

	if payloadCid == "" || nftTxHash == "" || tokenId == "" || mintAddress == "" {
		errMsg := "payload_cid, tx_hash, token_id and mint_address cannot be nil"
		err := errors.New(errMsg)
		logs.GetLogger().Error(err)
		c.JSON(http.StatusBadRequest, common.CreateErrorResponse(errorinfo.HTTP_REQUEST_PARAMS_NULL_ERROR_CODE, errMsg))
		return
	}

	sourceFile, err := models.GetSourceFileByPayloadCid(payloadCid)
	if err != nil {
		logs.GetLogger().Error(err)
		c.JSON(http.StatusBadRequest, common.CreateErrorResponse(errorinfo.SAVE_DATA_TO_DB_ERROR_CODE))
		return
	} else {
		sourceFileMint := models.SourceFileMint{
			SourceFileUploadId: sourceFile.ID,
			NftTxHash:          nftTxHash,
			TokenId:            tokenId,
			MintAddress:        mintAddress,
			CreateAt:           utils.GetCurrentUtcSecond(),
		}

		database.SaveOne(sourceFileMint)
		c.JSON(http.StatusOK, common.CreateSuccessResponse(sourceFileMint))
	}
}
