package rucweb

import (
	"crypto/hmac"
	"crypto/md5"
	"encoding/hex"
)

func GetMD5(password, token string) (string, error) {
	hash := hmac.New(md5.New, []byte(token))
	_, err := hash.Write([]byte(password))
	if err != nil {
		return "", err
	}
	return hex.EncodeToString(hash.Sum(nil)), nil
}
