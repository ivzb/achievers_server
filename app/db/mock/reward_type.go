package mock

import "github.com/ivzb/achievers_server/app/db"

type RewardType struct {
	db *DB

	ExistsMock RewardTypeExists
}

type RewardTypeExists struct {
	Bool bool
	Err  error
}

func (db *DB) RewardType() db.RewardTyper {
	return &RewardType{db: db}
}

func (ctx *RewardType) Exists(id uint8) (bool, error) {
	return ctx.ExistsMock.Bool, ctx.ExistsMock.Err
}
