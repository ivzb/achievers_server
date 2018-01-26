package crypto

import (
	"crypto/rand"
	"crypto/rsa"
	"crypto/sha256"
	"crypto/x509"
	"encoding/pem"
	"testing"
)

const bits = 1024

func TestGenerate(t *testing.T) {
	priv, err := Generate(bits)

	if err != nil {
		t.Fatalf("Generate returned error, expected rsa.PrivateKey")
	}

	if priv == nil {
		t.Fatalf("Generate returned nil, expected rsa.PrivateKey")
	}
}

func TestEncrypt_ValidMessage(t *testing.T) {
	priv := generatePrivateKey(t)

	expected := "message to be encrypted"

	encryptedMessage, err := Encrypt([]byte(expected), &priv.PublicKey)

	if err != nil {
		t.Fatalf("Encrypt returned error: %v", err)
	}

	decryptedMessage, err := rsa.DecryptOAEP(sha256.New(), rand.Reader, priv, encryptedMessage, nil)

	if err != nil {
		t.Fatalf("DecryptOAEP returned error: %v", err)
	}

	actual := string(decryptedMessage)

	if expected != actual {
		t.Fatalf("Decrypt returned wrong value: \nexpected %v, \nactual %v",
			expected, actual)
	}
}

func TestEncrypt_InvalidMessage(t *testing.T) {
	priv := generatePrivateKey(t)

	var expected []byte

	_, err := Encrypt(expected, &priv.PublicKey)

	if err != nil {
		t.Fatalf("Encrypt should return error")
	}
}

func TestDecrypt_ValidMessage(t *testing.T) {
	priv := generatePrivateKey(t)

	expected := "message to be encrypted"

	encryptedMessage, err := rsa.EncryptOAEP(sha256.New(), rand.Reader, &priv.PublicKey, []byte(expected), nil)

	if err != nil {
		t.Fatalf("Encrypt returned error: %v", err)
	}

	decryptedMessage, err := Decrypt(encryptedMessage, priv)

	if err != nil {
		t.Fatalf("Decrypt returned error: %v", err)
	}

	actual := string(decryptedMessage)

	if expected != actual {
		t.Fatalf("Decrypt returned wrong value: \nexpected %v, \nactual %v",
			expected, actual)
	}
}

func TestDecrypt_InvalidMessage(t *testing.T) {
	priv := generatePrivateKey(t)

	var expected []byte

	_, err := Decrypt(expected, priv)

	if err == nil {
		t.Fatalf("Decrypt should return error")
	}
}

func TestExport(t *testing.T) {
	priv := generatePrivateKey(t)

	pem := Export(priv)

	if len(pem) == 0 {
		t.Fatalf("Export returned empty pem")
	}
}

func TestImport_ValidPem(t *testing.T) {
	expected := generatePrivateKey(t)

	privateKeyBytes := x509.MarshalPKCS1PrivateKey(expected)
	privateKeyPem := pem.EncodeToMemory(
		&pem.Block{
			Type:  "RSA PRIVATE KEY",
			Bytes: privateKeyBytes,
		},
	)

	actual, err := Import(privateKeyPem)

	if err != nil {
		t.Fatalf("Import returned error: %v", err)
	}

	if actual == nil {
		t.Fatalf("Import returned nil PrivateKey")
	}
}

func TestImport_InvalidPemBytes(t *testing.T) {
	invalidPem := []byte{1, 2, 3}

	_, err := Import(invalidPem)

	if err == nil {
		t.Fatalf("Import should return error")
	}
}

func TestGenerateEncryptDecrypt(t *testing.T) {
	priv, err := Generate(bits)

	if err != nil {
		t.Fatalf("Generate returned error, expected rsa.PrivateKey")
	}

	expected := "message to be encrypted"

	encrypted, err := Encrypt([]byte(expected), &priv.PublicKey)

	if err != nil {
		t.Fatalf("Encrypt returned error, expected encrypted message")
	}

	decrypted, err := Decrypt(encrypted, priv)

	if err != nil {
		t.Fatalf("Decrypt returned error, expected decrypted message")
	}

	actual := string(decrypted)

	if expected != actual {
		t.Fatalf("Decrypt returned wrong value: \nexpected %v, \nactual %v",
			expected, actual)
	}
}

func TestGenerateExportImport(t *testing.T) {
	priv, err := Generate(bits)

	if err != nil {
		t.Fatalf("Generate returned error, expected PrivateKey")
	}

	expected := "message to be encrypted"
	encrypted, err := Encrypt([]byte(expected), &priv.PublicKey)

	if err != nil {
		t.Fatalf("Encrypt returned error, expected encrypted message")
	}

	pem := Export(priv)

	importedPriv, err := Import(pem)

	if err != nil {
		t.Fatalf("Import returned error, expected PrivateKey")
	}

	decrypted, err := Decrypt(encrypted, importedPriv)

	if err != nil {
		t.Fatalf("Decrypt returned error, expected decrypted message")
	}

	actual := string(decrypted)

	if expected != actual {
		t.Fatalf("Decrypt returned wrong value: \nexpected %v, \nactual %v",
			expected, actual)
	}
}

func generatePrivateKey(t *testing.T) *rsa.PrivateKey {
	reader := rand.Reader
	bits := 1024

	priv, err := rsa.GenerateKey(reader, bits)

	if err != nil {
		t.Fatalf("Generate returned error, expected rsa.PrivateKey")
	}

	return priv
}
