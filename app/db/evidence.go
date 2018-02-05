package db

import (
	"github.com/ivzb/achievers_server/app/model"
	"github.com/ivzb/achievers_server/app/shared/consts"
)

type Evidencer interface {
	Exists(id interface{}) (bool, error)
	Single(id string) (interface{}, error)
	Create(evidence interface{}) (string, error)

	LastID() (string, error)
	After(id string) ([]interface{}, error)
}

type Evidence struct {
	*Context
}

func (db *DB) Evidence() Evidencer {
	return &Evidence{
		newContext(db, consts.Evidence, new(model.Evidence)),
	}
}

func (ctx *Evidence) Exists(id interface{}) (bool, error) {
	return ctx.exists(&model.Evidence{ID: id.(string)})
}

func (ctx *Evidence) Single(id string) (interface{}, error) {
	return ctx.single(id)
}

func (ctx *Evidence) Create(evidence interface{}) (string, error) {
	return ctx.create(evidence)
}

func (ctx *Evidence) LastID() (string, error) {
	return ctx.lastID()
}

func (ctx *Evidence) After(id string) ([]interface{}, error) {
	return ctx.after(id)
}
