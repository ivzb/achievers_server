package model

import (
	"time"
)

type Quest struct {
	ID string `json:"id"`

	Title      string `json:"title"`
	PictureURL string `json:"picture_url"`

	InvolvementID string `json:"involvement_id"`
	QuestTypeID   uint8  `json:"quest_type_id"`
	UserID        string `json:"user_id"`

	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	DeletedAt time.Time `json:"deleted_at"`
}
