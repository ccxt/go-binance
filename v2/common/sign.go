package common

import (
	"crypto"
	"crypto/ed25519"
	"crypto/hmac"
	"crypto/rand"
	"crypto/rsa"
	"crypto/sha256"
	"crypto/x509"
	"encoding/base64"
	"encoding/pem"
	"errors"
	"fmt"
)

const (
	KeyTypeHmac    = "HMAC"
	KeyTypeRsa     = "RSA"
	KeyTypeEd25519 = "ED25519"
)

func SignFunc(keyType string) (func(string, string) (*string, error), error) {
	switch {
	case keyType == KeyTypeHmac:
		return Hmac, nil
	case keyType == KeyTypeRsa:
		return Rsa, nil
	case keyType == KeyTypeEd25519:
		return Ed25519, nil
	default:
		return nil, fmt.Errorf("unsupported keyType=%s", keyType)
	}
}

func Hmac(secretKey string, data string) (*string, error) {
	mac := hmac.New(sha256.New, []byte(secretKey))
	_, err := mac.Write([]byte(data))
	if err != nil {
		return nil, err
	}
	encodeData := fmt.Sprintf("%x", (mac.Sum(nil)))
	return &encodeData, nil
}

func Rsa(secretKey string, data string) (*string, error) {
	block, _ := pem.Decode([]byte(secretKey))
	if block == nil {
		return nil, errors.New("Rsa pem.Decode failed, invalid pem format secretKey")
	}
	privateKey, err := x509.ParsePKCS8PrivateKey(block.Bytes)
	if err != nil {
		return nil, fmt.Errorf("Rsa ParsePKCS8PrivateKey failed, error=%v", err.Error())
	}
	rsaPrivateKey, ok := privateKey.(*rsa.PrivateKey)
	if !ok {
		return nil, fmt.Errorf("Rsa convert PrivateKey failed")
	}
	hashed := sha256.Sum256([]byte(data))
	signature, err := rsa.SignPKCS1v15(rand.Reader, rsaPrivateKey, crypto.SHA256, hashed[:])
	if err != nil {
		return nil, err
	}
	encodedSignature := base64.StdEncoding.EncodeToString(signature)
	return &encodedSignature, nil
}

func Ed25519(secretKey string, data string) (*string, error) {
	block, _ := pem.Decode([]byte(secretKey))
	if block == nil {
		return nil, fmt.Errorf("Ed25519 pem.Decode failed, invalid pem format secretKey")
	}
	privateKey, err := x509.ParsePKCS8PrivateKey(block.Bytes)
	if err != nil {
		return nil, fmt.Errorf("Ed25519 call ParsePKCS8PrivateKey failed, error=%v", err.Error())
	}
	ed25519PrivateKey, ok := privateKey.(ed25519.PrivateKey)
	if !ok {
		return nil, fmt.Errorf("Ed25519 convert PrivateKey failed")
	}
	pk := ed25519.PrivateKey(ed25519PrivateKey)
	signature := ed25519.Sign(pk, []byte(data))
	encodedSignature := base64.StdEncoding.EncodeToString(signature)
	return &encodedSignature, nil
}
