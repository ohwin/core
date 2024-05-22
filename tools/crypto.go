package tools

import (
	"crypto/md5"
	"encoding/hex"
)

func MD5(str string) string {
	hash := md5.Sum([]byte(str))
	encryptedData := hex.EncodeToString(hash[:])
	return encryptedData
}
