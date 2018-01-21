package model

import (
	"time"
)

type Reward struct {
	ID string `json:"id"`

	Title       string `json:"title"`
	Description string `json:"description"`
	PictureURL  string `json:"picture_url"`

	RewardTypeID uint8  `json:"reward_type_id"`
	UserID       string `json:"user_id"`

	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	DeletedAt time.Time `json:"deleted_at"`
}
