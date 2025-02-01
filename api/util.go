package api

import (
	"crypto/rand"
	"io"

	"github.com/xc5216/circle-console-go/internal/util"
)

// GenerateRandomEntitySecret generates a random entity secret
func GenerateRandomEntitySecret() []byte {
	mainBuff := make([]byte, 32)
	_, err := io.ReadFull(rand.Reader, mainBuff)
	if err != nil {
		panic("reading from crypto/rand failed: " + err.Error())
	}
	return mainBuff
}

func EncryptEntitySecret(entitySecret []byte, publicKeyString string) ([]byte, error) {
	if len(entitySecret) != 32 {
		panic("invalid entity secret")
	}

	pubKey, err := util.ParseRsaPublicKeyFromPem([]byte(publicKeyString))
	if err != nil {
		return nil, err
	}

	cipher, err := util.EncryptOAEP(pubKey, entitySecret)
	if err != nil {
		panic(err)
	}
	return cipher, nil
}
