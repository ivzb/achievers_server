package db

import (
	"github.com/ivzb/achievers_server/app/model"
	"github.com/ivzb/achievers_server/app/shared/consts"
)

type Involvementer interface {
	Exists(id interface{}) (bool, error)
}

type Involvement struct {
	*Context
}

func (db *DB) Involvement() Involvementer {
	return &Involvement{
		newContext(db, consts.Involvement, new(model.Involvement)),
	}
}

func (ctx *Involvement) Exists(id interface{}) (bool, error) {
	return ctx.exists(&model.Involvement{ID: id.(uint8)})
}
