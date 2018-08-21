package activitypub

import (
	"crypto"
	"crypto/rsa"
	"crypto/x509"
	"encoding/pem"
	"fmt"
	"github.com/writeas/openssl-go"
	"log"
)

const keyBitSize = 2048

// GenerateKeys creates an RSA keypair and returns the public and private key,
// in that order.
func GenerateKeys() (pubPEM []byte, privPEM []byte) {
	var err error
	privPEM, err = openssl.Call(nil, "genrsa", fmt.Sprintf("%d", keyBitSize))
	if err != nil {
		log.Printf("Unable to generate private key: %v", err)
		return nil, nil
	}

	pubPEM, err = openssl.Call(privPEM, "rsa", "-in", "/dev/stdin", "-pubout")
	if err != nil {
		log.Printf("Unable to get public key: %v", err)
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

func parsePublicKey(der []byte) (crypto.PublicKey, error) {
	if key, err := x509.ParsePKCS1PublicKey(der); err == nil {
		return key, nil
	}
	if key, err := x509.ParsePKIXPublicKey(der); err == nil {
		switch key := key.(type) {
		case *rsa.PublicKey:
			return key, nil
		default:
			return nil, fmt.Errorf("found unknown public key type in PKIX wrapping")
		}
	}

	return nil, fmt.Errorf("failed to parse public key")
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

// DecodePublicKey decodes public keys
func DecodePublicKey(k []byte) (crypto.PublicKey, error) {
	block, _ := pem.Decode(k)
	if block == nil || block.Type != "PUBLIC KEY" {
		if block != nil {
			return nil, fmt.Errorf("failed to decode PEM block containing public key. type: %v", block.Type)
		} else {
			return nil, fmt.Errorf("failed to decode PEM block containing public key.")
		}
	}

	return parsePublicKey(block.Bytes)
}
