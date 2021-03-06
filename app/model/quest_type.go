package model

import (
	"time"
)

type QuestType struct {
	ID uint8 `json:"id" exists:"id"`

	Title string `json:"title"`

	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	DeletedAt time.Time `json:"deleted_at"`
}
