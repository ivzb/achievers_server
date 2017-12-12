package token

import (
	"log"
	"encoding/json"
	"crypto/rsa"
	"app/shared/crypto"
)

type Info struct {
	File string `json:"File"`
	// Priv *rsa.PrivateKey
}

// type Token struct {
// 	AuthToken []byte `json:"auth_token"`
// }

func Encrypt(pub *rsa.PublicKey, s string) ([]byte, error) {
	at, err := crypto.Encrypt([]byte(s), pub)

	log.Println(at)
	log.Println(string(at))

	r, err := json.Marshal(at)

	log.Println(r)
	log.Println(string(r))

	var js []byte
	err = json.Unmarshal([]byte(r), &js)

	log.Println(js)
	log.Println(string(js))

	if err != nil {
		return nil, err
	}

	return at, nil
}

func Decrypt(priv *rsa.PrivateKey, at string) ([]byte, error) {
	// j := `{"auth_token":"` + at + `"}`
	var r []byte

	if err := json.Unmarshal([]byte(at), &r); err != nil {
		return nil, err
	}

    d, err := crypto.Decrypt(r, priv)

	return d, err
}