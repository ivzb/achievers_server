package token

import (
	"io/ioutil"
	"os"
	"path"

	"github.com/ivzb/achievers_server/app/shared/crypto"
)

type Info struct {
	File string `json:"File"`
}

func (info *Info) EnsureExists() error {
	if _, err := os.Stat(info.File); os.IsNotExist(err) {
		dirs := path.Dir(info.File)
		err = os.MkdirAll(dirs, 0777)

		if err != nil {
			return err
		}

		// Generate and write key to file if doesn't already exist
		priv, err := crypto.Generate()
		pem := crypto.Export(priv)
		err = ioutil.WriteFile(info.File, pem, 0600)

		if err != nil {
			return err
		}
	}

	return nil
}
