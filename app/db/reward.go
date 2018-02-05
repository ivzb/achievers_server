package db

import (
	"github.com/ivzb/achievers_server/app/model"
	"github.com/ivzb/achievers_server/app/shared/consts"
)

type Rewarder interface {
	Exists(id interface{}) (bool, error)
	Single(id string) (interface{}, error)
	Create(reward interface{}) (string, error)

	LastID() (string, error)
	After(id string) ([]interface{}, error)
}

type Reward struct {
	*Context
}

func (db *DB) Reward() Rewarder {
	return &Reward{
		newContext(db, consts.Reward, new(model.Reward)),
	}
}

func (ctx *Reward) Exists(id interface{}) (bool, error) {
	return ctx.exists(&model.Reward{ID: id.(string)})
}

func (ctx *Reward) Single(id string) (interface{}, error) {
	return ctx.single(id)
}

func (ctx *Reward) Create(reward interface{}) (string, error) {
	return ctx.create(reward)
}

func (ctx *Reward) LastID() (string, error) {
	return ctx.lastID()
}

func (ctx *Reward) After(id string) ([]interface{}, error) {
	return ctx.after(id)
}
