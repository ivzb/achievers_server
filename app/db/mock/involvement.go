package mock

import "github.com/ivzb/achievers_server/app/db"

type InvolvementExists struct {
	Bool bool
	Err  error
}

type Involvement struct {
	db *DB

	ExistsMock InvolvementExists
}

func (db *DB) Involvement() db.Involvementer {
	return &Involvement{db: db}
}

func (ctx *Involvement) Exists(id string) (bool, error) {
	return ctx.ExistsMock.Bool, ctx.ExistsMock.Err
}

//func (mock *DB) InvolvementExists(string) (bool, error) {
//return mock.InvolvementExistsMock.Bool, mock.InvolvementExistsMock.Err
//}
