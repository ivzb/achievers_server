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
		newContext(db, "reward_type", nil),
	}
}

func (ctx *RewardType) Exists(id uint8) (bool, error) {
	return ctx.exists("id", strconv.FormatInt(int64(id), 10))
}
