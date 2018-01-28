package db

import "strconv"

type Involvementer interface {
	Exists(id uint8) (bool, error)
}

type Involvement struct {
	*Context
}

func (db *DB) Involvement() Involvementer {
	return &Involvement{
		&Context{
			db:    db,
			table: "involvement",
		},
	}
}

func (ctx *Involvement) Exists(id uint8) (bool, error) {
	return exists(ctx.Context, "id", strconv.FormatInt(int64(id), 10))
}
