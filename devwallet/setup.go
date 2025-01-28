package devwallet

import (
	"crypto/rand"
	"crypto/rsa"
	"crypto/sha256"
	"crypto/x509"
	"encoding/json"
	"encoding/pem"
	"errors"
	"fmt"
	"io"
	"net/http"

	"github.com/xc5216/circle-console-go/internal/setting"
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

// GetPublicKey will get public key from circle server
func GetPublicKey(apiKey string) (string, error) {
	url := fmt.Sprintf("%s/config/entity/publicKey", setting.GetServerURL())

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return "", err
	}

	req.Header.Add("Content-Type", "application/json")
	req.Header.Add("Authorization", fmt.Sprintf("Bearer %s", apiKey))

	res, err := http.DefaultClient.Do(req)
	if err != nil {
		return "", err
	}

	defer res.Body.Close()
	body, err := io.ReadAll(res.Body)
	if err != nil {
		return "", err
	}

	publicRes := publicKeyResponse{}
	err = json.Unmarshal(body, &publicRes)

	return publicRes.Data.PublicKey, err
}

func EncryptEntitySecret(entitySecret []byte, publicKeyString string) ([]byte, error) {
	if len(entitySecret) != 32 {
		panic("invalid entity secret")
	}

	pubKey, err := parseRsaPublicKeyFromPem([]byte(publicKeyString))
	if err != nil {
		return nil, err
	}

	cipher, err := encryptOAEP(pubKey, entitySecret)
	if err != nil {
		panic(err)
	}
	return cipher, nil
}

// parseRsaPublicKeyFromPem parse rsa public key from pem.
func parseRsaPublicKeyFromPem(pubPEM []byte) (*rsa.PublicKey, error) {
	block, _ := pem.Decode(pubPEM)
	if block == nil {
		return nil, errors.New("failed to parse PEM block containing the key")
	}
	pub, err := x509.ParsePKIXPublicKey(block.Bytes)
	if err != nil {
		return nil, err
	}
	switch pub := pub.(type) {
	case *rsa.PublicKey:
		return pub, nil
	default:
	}
	return nil, errors.New("key type is not rsa")
}

// encryptOAEP rsa encrypt oaep.
func encryptOAEP(pubKey *rsa.PublicKey, message []byte) (ciphertext []byte, err error) {
	random := rand.Reader
	ciphertext, err = rsa.EncryptOAEP(sha256.New(), random, pubKey, message, nil)
	if err != nil {
		return nil, err
	}
	return
}
