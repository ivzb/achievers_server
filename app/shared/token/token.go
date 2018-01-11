package token

import (
	"io/ioutil"
	"os"
	"path"

	"github.com/ivzb/achievers_server/app/shared/crypto"
	"github.com/ivzb/achievers_server/app/shared/file"
)

type Info struct {
	Path string `json:"Path"`
}

func (info *Info) EnsureExists() error {
	if !file.Exists(info.Path) {
		dirs := path.Dir(info.Path)
		os.MkdirAll(dirs, 0777)

		// Generate and write key to file if doesn't already exist
		priv, err := crypto.Generate()
		pem := crypto.Export(priv)
		err = ioutil.WriteFile(info.Path, pem, 0600)

		if err != nil {
			return err
		}
	}

	return nil
}
