package mock

import "github.com/ivzb/achievers_server/app/db"

type MultimediaTypeExists struct {
	Bool bool
	Err  error
}

type MultimediaType struct {
	db *DB

	ExistsMock MultimediaTypeExists
}

func (db *DB) MultimediaType() db.MultimediaTyper {
	return &MultimediaType{db: db}
}

func (ctx *MultimediaType) Exists(id uint8) (bool, error) {
	return ctx.ExistsMock.Bool, ctx.ExistsMock.Err
}

//func (mock *DB) MultimediaTypeExists(uint8) (bool, error) {
//return mock.MultimediaTypeExistsMock.Bool, mock.MultimediaTypeExistsMock.Err
//}
