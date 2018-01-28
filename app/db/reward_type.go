package db

import "strconv"

type RewardTyper interface {
	Exists(id uint8) (bool, error)
}

type RewardType struct {
	*Context
}

func (db *DB) RewardType() RewardTyper {
	return &RewardType{
		&Context{
			db:    db,
			table: "reward_type",
		},
	}
}

func (ctx *RewardType) Exists(id uint8) (bool, error) {
	return exists(ctx.Context, "id", strconv.FormatInt(int64(id), 10))
}
