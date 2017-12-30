package model

import (
	"time"
)

type Reward struct {
	ID string `json:"id"`

	Name        string `json:"name"`
	Description string `json:"description"`
	PictureUrl  string `json:"picture_url"`

	RewardType string `json:"reward_type"`

	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	DeletedAt time.Time `json:"deleted_at"`
}
