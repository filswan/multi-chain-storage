package routers

import (
	"errors"
	"fmt"
	"multi-chain-storage/common"
	"multi-chain-storage/common/errorinfo"
	"multi-chain-storage/service"
	"net/http"

	"github.com/filswan/go-swan-lib/logs"
	"github.com/gin-gonic/gin"
)

func Mint(router *gin.RouterGroup) {
	router.POST("/mint/info", RecordMintInfo)
}

type mintInfoUpload struct {
	SourceFileIploadId int64  `json:"source_file_upload_id"`
	TxHash             string `json:"tx_hash"`
	TokenId            string `json:"token_id"`
	MintAddress        string `json:"mint_address"`
}

func RecordMintInfo(c *gin.Context) {
	var model mintInfoUpload
	c.BindJSON(&model)

	sourceFileIploadId := model.SourceFileIploadId
	nftTxHash := model.TxHash
	tokenId := model.TokenId
	mintAddress := model.MintAddress

	if sourceFileIploadId <= 0 {
		err := fmt.Errorf("source_file_upload_id must be greater than 0")
		logs.GetLogger().Error(err)
		c.JSON(http.StatusBadRequest, common.CreateErrorResponse(errorinfo.ERROR_PARAM_INVALID_VALUE, err.Error()))
		return
	}

	if nftTxHash == "" || tokenId == "" || mintAddress == "" {
		errMsg := "tx_hash, token_id and mint_address cannot be empty"
		err := errors.New(errMsg)
		logs.GetLogger().Error(err)
		c.JSON(http.StatusBadRequest, common.CreateErrorResponse(errorinfo.ERROR_PARAM_NULL, errMsg))
		return
	}

	sourceFileMint, err := service.RecordMintInfo(sourceFileIploadId, nftTxHash, tokenId, mintAddress)
	if err != nil {
		logs.GetLogger().Error(err)
		c.JSON(http.StatusBadRequest, common.CreateErrorResponse(errorinfo.ERROR_INTERNAL, err.Error()))
		return
	}

	c.JSON(http.StatusOK, common.CreateSuccessResponse(sourceFileMint))
}
