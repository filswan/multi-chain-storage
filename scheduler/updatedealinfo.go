package scheduler

import (
	"encoding/json"
	"github.com/robfig/cron"
	"net/http"
	"payment-bridge/common/httpClient"
	"payment-bridge/common/utils"
	"payment-bridge/config"
	"payment-bridge/logs"
	"payment-bridge/models"
	"strconv"
	"time"
)

func UpdateDealScheduler() {
	c := cron.New()
	err := c.AddFunc(config.GetConfig().ScheduleRule.SendDealRule, func() {
		logs.GetLogger().Println("^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^ send deal scheduler is running at " + time.Now().Format("2006-01-02 15:04:05"))
		err := UpdateDealListSigService()
		if err != nil {
			logs.GetLogger().Error(err)
			return
		}
	})
	if err != nil {
		logs.GetLogger().Error(err)
		return
	}
	c.Start()
}

func UpdateDealListSigService() error {
	pageSize := 10
	currentTime := strconv.FormatInt(utils.GetEpochInMillis(), 10)
	url := config.GetConfig().SwanApi.ApiUrl + "/paymentgateway/deals?source_id=4&is_public=1&offset=" + strconv.Itoa(0) + "&limit=" + strconv.Itoa(pageSize)
	response, err := httpClient.SendRequestAndGetBytes(http.MethodGet, url, nil, nil)
	if err != nil {
		logs.GetLogger().Error(err)
		return err
	}
	var results *models.OfflineDealResult
	err = json.Unmarshal(response, &results)
	if err != nil {
		logs.GetLogger().Error(err)
		return err
	}
	for _, v := range results.Data.Deals {
		err := models.UpdateDealFile(&models.DealFile{PayloadCid: v.PayloadCid, PieceCid: v.PieceCid}, map[string]interface{}{"deal_status": v.Status, "deal_cid": v.DealCid, "cost": v.Cost, "miner_fid": v.MinerFid, "update_at": currentTime})
		if err != nil {
			logs.GetLogger().Error(err)
			continue
		}
	}
	totalRecord := results.PagingInfo.TotalItems
	totalPageNum := (totalRecord + pageSize - 1) / pageSize
	for i := 1; i <= totalPageNum; i++ {
		if i == 1 {
			continue
		}
		offset, err := utils.GetOffsetByPagenumber(strconv.Itoa(i), strconv.Itoa(pageSize))
		if err != nil {
			logs.GetLogger().Error(err)
			time.Sleep(60 * time.Second)
			continue
		}
		url := config.GetConfig().SwanApi.ApiUrl + "/paymentgateway/deals?source_id=4&is_public=1&offset=" + strconv.FormatInt(offset, 10) + "&limit=" + strconv.Itoa(pageSize)
		response, err := httpClient.SendRequestAndGetBytes(http.MethodGet, url, nil, nil)
		if err != nil {
			logs.GetLogger().Error(err)
			time.Sleep(60 * time.Second)
			continue
		}
		var results *models.OfflineDealResult
		err = json.Unmarshal(response, &results)
		if err != nil {
			logs.GetLogger().Error(err)
			time.Sleep(60 * time.Second)
			continue
		}
		for _, v := range results.Data.Deals {
			err := models.UpdateDealFile(&models.DealFile{PayloadCid: v.PayloadCid, PieceCid: v.PieceCid}, map[string]interface{}{"deal_status": v.Status, "deal_cid": v.DealCid, "cost": v.Cost, "miner_fid": v.MinerFid, "update_at": currentTime})
			if err != nil {
				logs.GetLogger().Error(err)
				continue
			}
		}
	}
	return nil
}
