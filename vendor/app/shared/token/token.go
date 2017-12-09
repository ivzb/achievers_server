package token

import (
	"crypto/rsa"
	"encoding/json"
	"io"
	"io/ioutil"
	"log"
	"os"

	"app/shared/crypto"
)

var (
	ti TokenInfo
)

type TokenInfo struct {
	File string `json:"File"`
	Priv *rsa.PrivateKey
}

type Token struct {
	AuthToken []byte `json:"auth_token"`
}

// Configure adds the settings for the SMTP server
func Configure(tokenInfo TokenInfo) {
	ti = tokenInfo

	var err error
	var input = io.ReadCloser(os.Stdin)

	if input, err = os.Open(ti.File); err != nil {
		log.Fatalln(err)
	}

	// Read the config file
	pem, err := ioutil.ReadAll(input)
	input.Close()

	if err != nil {
		log.Fatalln(err)
	}

	ti.Priv, err = crypto.ImportPrivatePem(pem)

	if err != nil {
		log.Fatalln(err)
	}
}

func Encrypt(s string) (*Token, error) {
	at, err := crypto.Encrypt([]byte(s), ti.Priv.PublicKey)

	if err != nil {
		return nil, err
	}

	return &Token{AuthToken: at}, nil
}

func Decrypt(at string) ([]byte, error) {
	j := `{"auth_token":"` + at + `"}`
	var t Token

	if err := json.Unmarshal([]byte(j), &t); err != nil {
		return nil, err
	}

	return crypto.Decrypt([]byte(t.AuthToken), ti.Priv)
}
