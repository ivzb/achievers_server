package db

import (
	"strconv"

	"github.com/ivzb/achievers_server/app/shared/consts"
)

type MultimediaTyper interface {
	Exists(id uint8) (bool, error)
}

type MultimediaType struct {
	*Context
}

func (db *DB) MultimediaType() MultimediaTyper {
	return &MultimediaType{
		newContext(db, consts.MultimediaType, nil),
	}
}

func (ctx *MultimediaType) Exists(id uint8) (bool, error) {
	return ctx.exists(consts.ID, strconv.FormatInt(int64(id), 10))
}
