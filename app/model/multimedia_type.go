package model

import (
	"time"
)

type MultimediaType struct {
	ID uint8 `json:"id"`

	Title string `json:"title"`

	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	DeletedAt time.Time `json:"deleted_at"`
}
