package activitypub

import (
	"crypto/rand"
	"crypto/rsa"
	"crypto/x509"
	"encoding/pem"
	"fmt"
	"log"
)

const keyBitSize = 2048

// GenerateKey creates an RSA keypair.
func GenerateKey() (*rsa.PrivateKey, error) {
	priv, err := rsa.GenerateKey(rand.Reader, keyBitSize)
	if err != nil {
		return nil, err
	}

	err = priv.Validate()
	if err != nil {
		return nil, err
	}

	return priv, nil
}

// EncodeKeysToPEM encodes public and private key to PEM format, returning
// them in that order.
func EncodeKeysToPEM(privKey *rsa.PrivateKey) ([]byte, []byte) {
	privDER := x509.MarshalPKCS1PrivateKey(privKey)

	// pem.Block
	privBlock := pem.Block{
		Type:  "RSA PRIVATE KEY",
		Bytes: privDER,
	}

	// Private key in PEM format
	privPEM := pem.EncodeToMemory(&privBlock)

	// Encode public key
	pubKey, ok := privKey.Public().(*rsa.PublicKey)
	if !ok {
		log.Printf("Public key isn't RSA!")
		return nil, nil
	}
	pubDER := x509.MarshalPKCS1PublicKey(pubKey)
	pubBlock := pem.Block{
		Type:  "PUBLIC KEY",
		Bytes: pubDER,
	}
	pubPEM := pem.EncodeToMemory(&pubBlock)

	return pubPEM, privPEM
}

// DecodePrivateKey encodes public and private key to PEM format, returning
// them in that order.
func DecodePrivateKey(k []byte) (*rsa.PrivateKey, error) {
	block, _ := pem.Decode(k)
	if block == nil || block.Type != "RSA PRIVATE KEY" {
		return nil, fmt.Errorf("failed to decode PEM block containing private key")
	}

	return x509.ParsePKCS1PrivateKey(block.Bytes)
}
