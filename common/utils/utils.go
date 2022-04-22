package utils

import (
	"time"
)

func GetCurrentUtcSecond() int64 {
	return time.Now().UnixNano() / 1e9
}
