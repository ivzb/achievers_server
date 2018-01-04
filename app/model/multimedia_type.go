package model

import (
	"strconv"
	"time"
)

type MultimediaType struct {
	ID uint8 `json:"id"`

	Title string `json:"title"`

	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	DeletedAt time.Time `json:"deleted_at"`
}

func (db *DB) MultimediaTypeExists(id uint8) (bool, error) {
	return exists(db, "multimedia_type", "id", strconv.FormatInt(int64(id), 10))
}
