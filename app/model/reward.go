package model

import "time"

type Reward struct {
	ID string `json:"id" select:"id"`

	Title       string `json:"title" select:"title" insert:"title"`
	Description string `json:"description" select:"description" insert:"description"`
	PictureURL  string `json:"picture_url" select:"picture_url" insert:"picture_url"`

	RewardTypeID uint8  `json:"reward_type_id" select:"reward_type_id" insert:"reward_type_id"`
	UserID       string `json:"user_id" select:"user_id" insert:"user_id"`

	CreatedAt time.Time `json:"created_at" select:"created_at"`
	UpdatedAt time.Time `json:"updated_at" select:"updated_at"`
	DeletedAt time.Time `json:"deleted_at" select:"deleted_at"`
}

//type Reward struct {
//ID string `json:"id" select:"id"`

//Title       string `json:"title" select:"title" insert:"title"`
//Description string `json:"description" select:"description" insert:"description"`
//PictureURL  string `json:"picture_url" select:"picture_url" insert:"picture_url"`

//RewardTypeID uint8  `json:"reward_type_id"`
//UserID       string `json:"user_id"`

//CreatedAt time.Time `json:"created_at"`
//UpdatedAt time.Time `json:"updated_at"`
//DeletedAt time.Time `json:"deleted_at"`
//}
