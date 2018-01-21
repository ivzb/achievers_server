package model

import (
	"time"
)

type Achievement struct {
	ID string `json:"id"`

	Title       string `json:"title"`
	Description string `json:"description"`
	PictureURL  string `json:"picture_url"`

	InvolvementID string `json:"involvement_id"`
	UserID        string `json:"user_id"`

	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	DeletedAt time.Time `json:"deleted_at"`
}
