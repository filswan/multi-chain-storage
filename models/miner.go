package models

import (
	"fmt"
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

	err = fmt.Errorf("no miner for fid:%s", fid)
	logs.GetLogger().Error(err)

	return nil, err
}
