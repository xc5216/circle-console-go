package util

import (
	"crypto/rand"
	"crypto/rsa"
	"crypto/sha256"
	"crypto/x509"
	"encoding/pem"
	"errors"
)

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
