package db

import (
	"github.com/ivzb/achievers_server/app/model"
	"github.com/ivzb/achievers_server/app/shared/consts"
)

type RewardTyper interface {
	Exists(id interface{}) (bool, error)
}

type RewardType struct {
	*Context
}

func (db *DB) RewardType() RewardTyper {
	return &RewardType{
		newContext(db, consts.RewardType, new(model.RewardType)),
	}
}

func (ctx *RewardType) Exists(id interface{}) (bool, error) {
	return ctx.exists(&model.RewardType{ID: id.(uint8)})
}
