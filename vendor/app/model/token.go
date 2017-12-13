package model

import (
	"crypto/rsa"
    "app/shared/token"
    "app/shared/crypto"
	"io"
	"io/ioutil"
	"os"
	"encoding/base64"
)

type TokenSource interface {
	Encrypt(string) (string, error)
	Decrypt(string) (string, error)
}

type Token struct {
    t *rsa.PrivateKey
}

func NewToken(t token.Info /*value string*/) (*Token, error) {
    var err error
	var input = io.ReadCloser(os.Stdin)

	if input, err = os.Open(t.File); err != nil {
		return nil, err
	}

	// Read the config file
	pem, err := ioutil.ReadAll(input)
	input.Close()

	if err != nil {
		return nil, err
	}

	priv, err := crypto.ImportPrivatePem(pem)

	if err != nil {
		return nil, err
	}

	return &Token{priv}, nil
}

func (tk *Token) Encrypt(token string) (string, error) {
	encrypted, err := crypto.Encrypt([]byte(token), &tk.t.PublicKey)

	encoded := base64.StdEncoding.EncodeToString(encrypted)

	if (err != nil) {
		return "", err
	}

	return encoded, nil
}

func (tk *Token) Decrypt(encoded string) (string, error) {
	decoded, err := base64.StdEncoding.DecodeString(encoded)

	decrypted, err := crypto.Decrypt([]byte(decoded), tk.t)

	if err != nil {
		return "", err
	}

	return string(decrypted), err
}