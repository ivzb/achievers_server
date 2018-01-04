package model

import (
	"time"
)

type Involvement struct {
	ID string `json:"id"`

	Title string `json:"title"`

	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	DeletedAt time.Time `json:"deleted_at"`
}

func (db *DB) InvolvementExists(id string) (bool, error) {
	return exists(db, "involvement", "id", id)
}
