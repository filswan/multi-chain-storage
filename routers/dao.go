package routers

import (
	"multi-chain-storage/common"
	"multi-chain-storage/common/errorinfo"
	"multi-chain-storage/common/utils"
	"multi-chain-storage/database"
	"multi-chain-storage/models"
	"multi-chain-storage/service"
	"net/http"
	"strconv"

	"github.com/filswan/go-swan-lib/logs"
	"github.com/gin-gonic/gin"
)

func Dao(router *gin.RouterGroup) {
	router.GET("/dao/signature/deals", GetDealListForDaoToSign)
	router.PUT("/dao/signature/deals", RecordDealListThatHaveBeenSignedByDao)
}

func GetDealListForDaoToSign(c *gin.Context) {
	dealList, err := service.GetShoulBeSignDealListFromDB()
	if err != nil {
		logs.GetLogger().Error(err)
		c.JSON(http.StatusInternalServerError, common.CreateErrorResponse(errorinfo.GET_RECORD_lIST_ERROR_CODE))
		return
	}
	c.JSON(http.StatusOK, common.CreateSuccessResponse(dealList))
}

type DealIdList struct {
	DealId     string `json:"deal_id"`
	PayloadCid string `json:"payload_cid"`
	Recipent   string `json:"recipent"`
	TxHash1    string `json:"tx_hash_1"`
	TxHash2    string `json:"tx_hash_2"`
	TxHash3    string `json:"tx_hash_3"`
}

type daoBackendResponse struct {
	PayloadCid      string `json:"payload_cid"`
	DealId          string `json:"deal_id"`
	SuccessDaoCount int    `json:"success_dao_count"`
}

func RecordDealListThatHaveBeenSignedByDao(c *gin.Context) {
	var dealIdList []DealIdList
	err := c.BindJSON(&dealIdList)
	if err != nil {
		logs.GetLogger().Error(err)
		c.AbortWithStatusJSON(http.StatusBadRequest, common.CreateErrorResponse(errorinfo.HTTP_REQUEST_PARSER_RESPONSE_TO_STRUCT_ERROR_CODE))
		return
	}

	daoSignRes := []daoBackendResponse{}

	for _, v := range dealIdList {
		daosignCount := 0
		deal_id, err := strconv.ParseInt(v.DealId, 10, 64)
		if err != nil {
			logs.GetLogger().Error(err)
		}

		if v.TxHash1 != "" {
			verification1, err := service.VerifyDaoSigOnContract(v.TxHash1)

			if err != nil {
				logs.GetLogger().Error(err)
			} else {
				err = service.SaveDaoEventFromTxHash(v.TxHash1, v.PayloadCid, v.Recipent, deal_id, verification1)
				if err != nil {
					logs.GetLogger().Error(err)
				}
				if verification1 {
					daosignCount++
				}
			}
		}

		if v.TxHash2 != "" {
			verification2, err := service.VerifyDaoSigOnContract(v.TxHash2)
			if err != nil {
				logs.GetLogger().Error(err)
			} else {
				err = service.SaveDaoEventFromTxHash(v.TxHash2, v.PayloadCid, v.Recipent, deal_id, verification2)
				if err != nil {
					logs.GetLogger().Error(err)
				}
				if verification2 {
					daosignCount++
				}
			}
		}

		if v.TxHash3 != "" {
			verification3, err := service.VerifyDaoSigOnContract(v.TxHash3)
			if err != nil {
				logs.GetLogger().Error(err)
			} else {
				err = service.SaveDaoEventFromTxHash(v.TxHash3, v.PayloadCid, v.Recipent, deal_id, verification3)
				if err != nil {
					logs.GetLogger().Error(err)
				}
				if verification3 {
					daosignCount++
				}
			}
		}

		events, err := models.GetEventDaoSignaturesByDealId(deal_id)

		if err != nil {
			logs.GetLogger().Error(err)
		} else {
			if daosignCount >= 2 || len(events) >= 2 {
				daoFetchedDeal := new(models.DaoFetchedDeal)
				dealIdIntValue, err := strconv.ParseInt(v.DealId, 10, 64)
				if err != nil {
					logs.GetLogger().Error(err)
					continue
				}
				daoFetchedDeal.DealId = dealIdIntValue
				daoFetchedDeal.CreateAt = utils.GetCurrentUtcSecond()
				err = database.SaveOne(daoFetchedDeal)
				if err != nil {
					logs.GetLogger().Error(err)
					continue
				}
				var response daoBackendResponse
				response.DealId = v.DealId
				response.PayloadCid = v.PayloadCid
				response.SuccessDaoCount = len(events)
				daoSignRes = append(daoSignRes, response)
			}
		}
	}

	c.JSON(http.StatusOK, common.CreateSuccessResponse(&daoSignRes))
}
