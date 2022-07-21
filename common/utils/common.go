package utils

import (
	"crypto/sha256"
	"encoding/hex"
	"fmt"

	"github.com/filswan/go-swan-lib/logs"
	"golang.org/x/crypto/sha3"
)

func GenerateHash(str string) string {
	hash := sha256.Sum256([]byte(str))
	hashStr := fmt.Sprintf("%x", hash)
	logs.GetLogger().Info(hashStr)
	return hashStr
}

func GenerateHashKeccak256(str string) (error, *string) {
	hash := sha3.NewLegacyKeccak256()

	var buf []byte
	_, err := hash.Write([]byte(str))
	if err != nil {
		logs.GetLogger().Error(err)
		return err, nil
	}
	buf = hash.Sum(nil)

	hashStr := hex.EncodeToString(buf)
	logs.GetLogger().Info("Keccak256 hash for string ", str, " is: ", hashStr)
	return nil, &hashStr
}
