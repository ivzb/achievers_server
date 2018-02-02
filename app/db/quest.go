package db

import (
	"github.com/ivzb/achievers_server/app/model"
	"github.com/ivzb/achievers_server/app/shared/consts"
)

type Quester interface {
	Exists(id string) (bool, error)
	Single(id string) (interface{}, error)
	Create(quest interface{}) (string, error)

	LastID() (string, error)
	After(id string) ([]interface{}, error)
}

type Quest struct {
	*Context
}

func (db *DB) Quest() Quester {
	return &Quest{
		newContext(db, consts.Quest, new(model.Quest)),
	}
}

func (ctx *Quest) Exists(id string) (bool, error) {
	return ctx.exists(consts.ID, id)
}

func (ctx *Quest) Single(id string) (interface{}, error) {
	return ctx.single(id)
}

func (ctx *Quest) Create(quest interface{}) (string, error) {
	return ctx.create(quest)
}

func (ctx *Quest) LastID() (string, error) {
	return ctx.lastID()
}

func (ctx *Quest) After(id string) ([]interface{}, error) {
	return ctx.after(id)
}
