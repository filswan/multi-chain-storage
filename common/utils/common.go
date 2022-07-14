package utils

import (
	"crypto/md5"
	"fmt"

	"github.com/filswan/go-swan-lib/logs"
)

func GenerateHash(str string) string {
	var hash = md5.Sum([]byte(str))
	hashStr := fmt.Sprintf("%x", hash)
	logs.GetLogger().Info(hashStr)
	return hashStr
}
