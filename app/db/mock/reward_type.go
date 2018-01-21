package mock

import "github.com/ivzb/achievers_server/app/db"

type RewardTypeExists struct {
	Bool bool
	Err  error
}

type RewardType struct {
	db *DB

	ExistsMock RewardTypeExists
}

func (db *DB) RewardType() db.RewardTyper {
	return &RewardType{db: db}
}

func (ctx *RewardType) Exists(id uint8) (bool, error) {
	return ctx.ExistsMock.Bool, ctx.ExistsMock.Err
}

//func (mock *DB) RewardTypeExists(uint8) (bool, error) {
//return mock.RewardTypeExistsMock.Bool, mock.RewardTypeExistsMock.Err
//}
