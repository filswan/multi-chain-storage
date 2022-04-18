package utils

import (
	"strconv"
	"time"

	"github.com/filswan/go-swan-lib/logs"
)

func GetCurrentUtcSecond() int64 {
	return time.Now().UnixNano() / 1e9
}

func GetOffsetByPagenumber(pageNumber, pageSize string) (int64, error) {
	pageNumberInt, err := strconv.ParseInt(pageNumber, 10, 64)
	if err != nil {
		logs.GetLogger().Error(err)
		return 0, err
	}
	pageSizeInt, err := strconv.ParseInt(pageSize, 10, 64)
	if err != nil {
		logs.GetLogger().Error(err)
		return 0, err
	}
	offset := (pageNumberInt - 1) * pageSizeInt
	return offset, nil
}
