package db

import "strconv"

type MultimediaTyper interface {
	Exists(id uint8) (bool, error)
}

type MultimediaType struct {
	*Context
}

func (db *DB) MultimediaType() MultimediaTyper {
	return &MultimediaType{
		newContext(db, "multimedia_type", nil),
	}
}

func (ctx *MultimediaType) Exists(id uint8) (bool, error) {
	return ctx.exists("id", strconv.FormatInt(int64(id), 10))
}
