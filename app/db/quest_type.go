package db

import "strconv"

type QuestTyper interface {
	Exists(id uint8) (bool, error)
}

type QuestType struct {
	db *DB
}

func (db *DB) QuestType() QuestTyper {
	return &QuestType{db}
}

func (ctx *QuestType) Exists(id uint8) (bool, error) {
	return exists(ctx.db, "quest_type", "id", strconv.FormatInt(int64(id), 10))
}
