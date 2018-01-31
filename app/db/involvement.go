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
		newContext(db, "involvement", nil),
	}
}

func (ctx *Involvement) Exists(id uint8) (bool, error) {
	return ctx.exists("id", strconv.FormatInt(int64(id), 10))
}
