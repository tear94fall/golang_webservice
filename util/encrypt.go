package util

import (
	"crypto/sha256"
	"encoding/hex"
	"errors"
)

func EncStr(data string) (string, error) {
	if data == "" || len(data) == 0 {
		return data, errors.New("input data is empty")
	}

	hash := sha256.New()

	hash.Write([]byte(data))

	md := hash.Sum(nil)
	mdStr := hex.EncodeToString(md)

	return mdStr, nil
}
