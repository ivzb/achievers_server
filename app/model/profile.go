package model

import (
	"time"
)

type Profile struct {
	ID string `json:"id" select:"id"`

	Name   string `json:"name" select:"name" insert:"name"`
	UserID string `insert:"user_id"`

	CreatedAt time.Time `json:"created_at" select:"created_at"`
	UpdatedAt time.Time `json:"updated_at" select:"updated_at"`
	DeletedAt time.Time `json:"deleted_at" select:"deleted_at"`
}
