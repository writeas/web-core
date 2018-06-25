package activitypub

import (
	"crypto"
	"crypto/rsa"
	"crypto/x509"
	"encoding/pem"
	"fmt"
	"github.com/spacemonkeygo/openssl"
	"log"
)

const keyBitSize = 2048

// GenerateKey creates an RSA keypair.
func GenerateKey() (openssl.PrivateKey, error) {
	return openssl.GenerateRSAKey(keyBitSize)
}

// EncodeKeysToPEM encodes public and private key to PEM format, returning
// them in that order.
func EncodeKeysToPEM(privKey openssl.PrivateKey) (pubPEM []byte, privPEM []byte) {
	var err error
	privPEM, err = privKey.MarshalPKCS1PrivateKeyPEM()
	if err != nil {
		log.Printf("Unable to marshal private key: %v", err)
		return nil, nil
	}

	pubPEM, err = privKey.MarshalPKIXPublicKeyPEM()
	if err != nil {
		log.Printf("Unable to marshal public key: %v", err)
		return nil, nil
	}
	return
}

func parsePrivateKey(der []byte) (crypto.PrivateKey, error) {
	if key, err := x509.ParsePKCS1PrivateKey(der); err == nil {
		return key, nil
	}
	if key, err := x509.ParsePKCS8PrivateKey(der); err == nil {
		switch key := key.(type) {
		case *rsa.PrivateKey:
			return key, nil
		default:
			return nil, fmt.Errorf("found unknown private key type in PKCS#8 wrapping")
		}
	}
	if key, err := x509.ParseECPrivateKey(der); err == nil {
		return key, nil
	}

	return nil, fmt.Errorf("failed to parse private key")
}

// DecodePrivateKey encodes public and private key to PEM format, returning
// them in that order.
func DecodePrivateKey(k []byte) (crypto.PrivateKey, error) {
	block, _ := pem.Decode(k)
	if block == nil || block.Type != "RSA PRIVATE KEY" {
		return nil, fmt.Errorf("failed to decode PEM block containing private key")
	}

	return parsePrivateKey(block.Bytes)
}
