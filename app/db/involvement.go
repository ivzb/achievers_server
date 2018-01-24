package db

import "strconv"

type Involvementer interface {
	Exists(id uint8) (bool, error)
}

type Involvement struct {
	db *DB
}

func (db *DB) Involvement() Involvementer {
	return &Involvement{db}
}

func (ctx *Involvement) Exists(id uint8) (bool, error) {
	return exists(ctx.db, "involvement", "id", strconv.FormatInt(int64(id), 10))
}
