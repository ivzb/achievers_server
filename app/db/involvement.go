package db

import (
	"strconv"

	"github.com/ivzb/achievers_server/app/shared/consts"
)

type Involvementer interface {
	Exists(id uint8) (bool, error)
}

type Involvement struct {
	*Context
}

func (db *DB) Involvement() Involvementer {
	return &Involvement{
		newContext(db, consts.Involvement, nil),
	}
}

func (ctx *Involvement) Exists(id uint8) (bool, error) {
	return ctx.exists(consts.ID, strconv.FormatInt(int64(id), 10))
}
