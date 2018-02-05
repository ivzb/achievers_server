package db

import (
	"github.com/ivzb/achievers_server/app/model"
	"github.com/ivzb/achievers_server/app/shared/consts"
)

type QuestTyper interface {
	Exists(id interface{}) (bool, error)
}

type QuestType struct {
	*Context
}

func (db *DB) QuestType() QuestTyper {
	return &QuestType{
		newContext(db, consts.QuestType, new(model.QuestType)),
	}
}

func (ctx *QuestType) Exists(id interface{}) (bool, error) {
	return ctx.exists(&model.QuestType{ID: id.(uint8)})
}
