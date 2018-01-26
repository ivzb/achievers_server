package crypto

import (
	"crypto/rand"
	"crypto/rsa"
	"crypto/x509"
	"encoding/pem"
	"errors"

	"github.com/minio/sha256-simd"
)

// Generate rsa private key with bits as param
func Generate(bits int) (*rsa.PrivateKey, error) {
	reader := rand.Reader

	return rsa.GenerateKey(reader, bits)
}

// Encrypt message with rsa private key
func Encrypt(message []byte, pub *rsa.PublicKey) ([]byte, error) {
	return rsa.EncryptOAEP(sha256.New(), rand.Reader, pub, message, nil)
}

// Decrypt message with rsa private key
func Decrypt(ciphertext []byte, priv *rsa.PrivateKey) ([]byte, error) {
	return rsa.DecryptOAEP(sha256.New(), rand.Reader, priv, ciphertext, nil)
}

// Export rsa private key
func Export(privateKey *rsa.PrivateKey) []byte {
	privateKeyBytes := x509.MarshalPKCS1PrivateKey(privateKey)
	privateKeyPem := pem.EncodeToMemory(
		&pem.Block{
			Type:  "RSA PRIVATE KEY",
			Bytes: privateKeyBytes,
		},
	)

	return privateKeyPem
}

// Import rsa private key
func Import(privateKeyPEM []byte) (*rsa.PrivateKey, error) {
	block, _ := pem.Decode(privateKeyPEM)

	if block == nil {
		return nil, errors.New("failed to parse PEM block containing the key")
	}

	return x509.ParsePKCS1PrivateKey(block.Bytes)
}
