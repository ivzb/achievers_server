package mock

import "github.com/ivzb/achievers_server/app/db"

type QuestType struct {
	db *DB

	ExistsMock QuestTypeExists
}

type QuestTypeExists struct {
	Bool bool
	Err  error
}

func (db *DB) QuestType() db.QuestTyper {
	return &QuestType{db: db}
}

func (ctx *QuestType) Exists(id uint8) (bool, error) {
	return ctx.ExistsMock.Bool, ctx.ExistsMock.Err
}
