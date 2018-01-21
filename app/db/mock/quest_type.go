package mock

import "github.com/ivzb/achievers_server/app/db"

type QuestTypeExists struct {
	Bool bool
	Err  error
}

type QuestType struct {
	db *DB

	ExistsMock QuestTypeExists
}

func (db *DB) QuestType() db.QuestTyper {
	return &QuestType{db: db}
}

func (ctx *QuestType) Exists(id uint8) (bool, error) {
	return ctx.ExistsMock.Bool, ctx.ExistsMock.Err
}

//func (mock *DB) QuestTypeExists(uint8) (bool, error) {
//return mock.QuestTypeExistsMock.Bool, mock.QuestTypeExistsMock.Err
//}
