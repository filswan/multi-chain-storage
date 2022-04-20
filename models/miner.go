package models

import (
	"multi-chain-storage/common/utils"
	"multi-chain-storage/database"

	"github.com/filswan/go-swan-lib/logs"
)

type Miner struct {
	ID       int64  `json:"id"`
	Fid      string `json:"fid"`
	CreateAt int64  `json:"create_at"`
}

func GeMinerByFid(fid string) (*Miner, error) {
	var miners []*Miner
	err := database.GetDB().Where("fid=?", fid).Find(&miners).Error
	if err != nil {
		logs.GetLogger().Error(err)
		return nil, err
	}

	if len(miners) > 0 {
		return miners[0], nil
	}

	miner, err := SaveMiner(fid)
	if err != nil {
		logs.GetLogger().Error(err)
		return nil, err
	}

	return miner, nil
}

func SaveMiner(fid string) (*Miner, error) {
	miner := Miner{
		Fid:      fid,
		CreateAt: utils.GetCurrentUtcSecond(),
	}

	minerResult := database.GetDB().Create(&miner)
	err := minerResult.Error
	if err != nil {
		logs.GetLogger().Error(err)
		return nil, err
	}

	minerCreated := minerResult.Value.(*Miner)

	return minerCreated, nil
}
