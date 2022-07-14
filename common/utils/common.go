package utils

import (
	"crypto/sha256"
	"fmt"

	"github.com/filswan/go-swan-lib/logs"
)

func GenerateHash(str string) string {
	hash := sha256.Sum256([]byte(str))
	hashStr := fmt.Sprintf("%x", hash)
	logs.GetLogger().Info(hashStr)
	return hashStr
}
