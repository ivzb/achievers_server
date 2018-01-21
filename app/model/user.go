package model

import (
	"time"
)

type User struct {
	ID string `json:"id"`

	Email    string `json:"email"`
	Password string `json:"password"`

	StatusID uint8 `json:"status_id"`

	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	DeletedAt time.Time `json:"deleted_at"`
}

type Auth struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}
