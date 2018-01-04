package model

import (
	"strconv"
	"time"
)

type RewardType struct {
	ID uint8 `json:"id"`

	Title string `json:"title"`

	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	DeletedAt time.Time `json:"deleted_at"`
}

func (db *DB) RewardTypeExists(id uint8) (bool, error) {
	return exists(db, "reward_type", "id", strconv.FormatInt(int64(id), 10))
}
