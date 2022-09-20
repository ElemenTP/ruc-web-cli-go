package rucweb

import (
	"crypto/sha1"
	"encoding/hex"
)

func GetSHA1(value string) (string, error) {
	hash := sha1.New()
	_, err := hash.Write([]byte(value))
	if err != nil {
		return "", err
	}
	return hex.EncodeToString(hash.Sum(nil)), nil
}
