package db

import (
	"strconv"

	"github.com/ivzb/achievers_server/app/shared/consts"
)

type RewardTyper interface {
	Exists(id uint8) (bool, error)
}

type RewardType struct {
	*Context
}

func (db *DB) RewardType() RewardTyper {
	return &RewardType{
		newContext(db, consts.RewardType, nil),
	}
}

func (ctx *RewardType) Exists(id uint8) (bool, error) {
	return ctx.exists("id", strconv.FormatInt(int64(id), 10))
}
