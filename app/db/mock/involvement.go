package mock

import "github.com/ivzb/achievers_server/app/db"

type Involvement struct {
	db *DB

	ExistsMock InvolvementExists
}

type InvolvementExists struct {
	Bool bool
	Err  error
}

func (db *DB) Involvement() db.Involvementer {
	return &Involvement{db: db}
}

func (ctx *Involvement) Exists(id string) (bool, error) {
	return ctx.ExistsMock.Bool, ctx.ExistsMock.Err
}
