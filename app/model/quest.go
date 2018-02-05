package model

import (
	"time"
)

type Quest struct {
	ID string `json:"id" select:"id" exists:"id"`

	Title      string `json:"title" select:"title" insert:"title"`
	PictureURL string `json:"picture_url" select:"picture_url" insert:"picture_url"`

	InvolvementID uint8  `json:"involvement_id" select:"involvement_id" insert:"involvement_id"`
	QuestTypeID   uint8  `json:"quest_type_id" select:"quest_type_id" insert:"quest_type_id"`
	UserID        string `json:"user_id" select:"user_id" insert:"user_id"`

	CreatedAt time.Time `json:"created_at" select:"created_at"`
	UpdatedAt time.Time `json:"updated_at" select:"updated_at"`
	DeletedAt time.Time `json:"deleted_at" select:"deleted_at"`
}
