package model

import (
	"crypto/md5"
	"encoding/hex"
)

// GeneratePasswordHash
func GeneratePasswordHash(pwd string) string {
	haser := md5.New()
	haser.Write([]byte(pwd))
	pwdHash := hex.EncodeToString(haser.Sum(nil))
	return pwdHash
}
