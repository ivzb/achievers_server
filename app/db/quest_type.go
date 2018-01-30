package db

import "strconv"

type QuestTyper interface {
	Exists(id uint8) (bool, error)
}

type QuestType struct {
	*Context
}

func (db *DB) QuestType() QuestTyper {
	return &QuestType{
		&Context{
			db:    db,
			table: "quest_type",
		},
	}
}

func (ctx *QuestType) Exists(id uint8) (bool, error) {
	return ctx.exists("id", strconv.FormatInt(int64(id), 10))
}
