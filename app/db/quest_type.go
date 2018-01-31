package db

import (
	"strconv"

	"github.com/ivzb/achievers_server/app/shared/consts"
)

type QuestTyper interface {
	Exists(id uint8) (bool, error)
}

type QuestType struct {
	*Context
}

func (db *DB) QuestType() QuestTyper {
	return &QuestType{
		newContext(db, consts.QuestType, nil),
	}
}

func (ctx *QuestType) Exists(id uint8) (bool, error) {
	return ctx.exists(consts.ID, strconv.FormatInt(int64(id), 10))
}
