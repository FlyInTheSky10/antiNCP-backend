package util

import (
	"crypto/rand"
	"crypto/rsa"
	"encoding/base64"
)

func RSADecryptFromString(toDecode string, key *rsa.PrivateKey) ([]byte, error) {
	bytes, err := base64.StdEncoding.DecodeString(toDecode)
	if err != nil {
		return nil, err
	}
	result, err := rsa.DecryptPKCS1v15(rand.Reader, key, bytes)

	if err != nil {
		return nil, err
	}
	return result, nil
}
