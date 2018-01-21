package db

import "strconv"

type MultimediaTyper interface {
	Exists(id uint8) (bool, error)
}

type MultimediaType struct {
	db *DB
}

func (db *DB) MultimediaType() MultimediaTyper {
	return &MultimediaType{db}
}

func (ctx *MultimediaType) Exists(id uint8) (bool, error) {
	return exists(ctx.db, "multimedia_type", "id", strconv.FormatInt(int64(id), 10))
}
