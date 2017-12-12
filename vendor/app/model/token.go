package model

import (
	"crypto/rsa"
    "app/shared/token"
    "app/shared/crypto"
	"io"
	"io/ioutil"
	"os"
)

type TokenSource interface {
    GetPrivateKey() (*rsa.PrivateKey)
    GetPublicKey() (*rsa.PublicKey)
	// Encrypt(string) ([]byte, error)
	// Decrypt(string) ([]byte, error)
}


// type DBSource interface {
// 	Exist(table string, column string, value string) (bool, error)

// 	AchievementsAll() ([]*Achievement, error)
// 	UserCreate(string, string, string, string) (string, error)
// 	UserAuth(string, string) (string, error)
// }

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

func (tk *Token) GetPrivateKey() (*rsa.PrivateKey) {
	return tk.t
}

func (tk *Token) GetPublicKey() (*rsa.PublicKey) {
	return &tk.t.PublicKey
}