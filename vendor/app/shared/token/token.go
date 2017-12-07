package token

import (
	"crypto/rsa"
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

// ReadConfig returns the token information
func ReadConfig() TokenInfo {
	return ti
}

func Extract(t []byte) ([]byte, error) {
	r, err := crypto.Decrypt(t, ti.Priv)

	if err != nil {
		return nil, err
	}

	return r, nil
}
