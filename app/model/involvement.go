package model

import (
	"time"
)

type Involvement struct {
	ID string `json:"id"`

	Name string `json:"name"`

	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	DeletedAt time.Time `json:"deleted_at"`
}

func (db *DB) InvolvementExists(id string) (bool, error) {
	return exists(db, "involvement", "id", id)
}
