package model

import (
	"time"
)

type RewardType struct {
	ID uint8 `json:"id"`

	Title string `json:"title"`

	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	DeletedAt time.Time `json:"deleted_at"`
}
