package utils

import (
	"crypto/sha1"
	"fmt"

	"github.com/filswan/go-swan-lib/logs"
)

func GenerateHash(str string) string {
	h := sha1.New()
	h.Write([]byte(str))
	bs := h.Sum(nil)

	hashStrHex := fmt.Sprintf("%x\n", bs)
	logs.GetLogger().Info(hashStrHex)
	return hashStrHex
}
