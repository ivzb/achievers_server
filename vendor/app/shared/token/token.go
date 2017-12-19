package token

import (
	"app/shared/crypto"
	"io/ioutil"
	"os"
	"path"
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
		output := crypto.ExportPrivatePem(priv)
		err = ioutil.WriteFile(info.File, output, 0600)

		if err != nil {
			return err
		}
	}

	return nil
}
