package db

import "strconv"

type RewardTyper interface {
	Exists(id uint8) (bool, error)
}

type RewardType struct {
	db *DB
}

func (db *DB) RewardType() RewardTyper {
	return &RewardType{db}
}

func (ctx *RewardType) Exists(id uint8) (bool, error) {
	return exists(ctx.db, "reward_type", "id", strconv.FormatInt(int64(id), 10))
}
