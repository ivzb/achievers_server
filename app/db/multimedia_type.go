package db

import (
	"github.com/ivzb/achievers_server/app/model"
	"github.com/ivzb/achievers_server/app/shared/consts"
)

type MultimediaTyper interface {
	Exists(id interface{}) (bool, error)
}

type MultimediaType struct {
	*Context
}

func (db *DB) MultimediaType() MultimediaTyper {
	return &MultimediaType{
		newContext(db, consts.MultimediaType, new(model.MultimediaType)),
	}
}

func (ctx *MultimediaType) Exists(id interface{}) (bool, error) {
	return ctx.exists(&model.MultimediaType{ID: id.(uint8)})
}
