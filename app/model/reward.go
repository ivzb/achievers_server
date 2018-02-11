package model

import "time"

type Reward struct {
	ID string `json:"id" select:"id" exists:"id"`

	Title       string `json:"title"       validate:"string.(min=1,max=255)" select:"title"       insert:"title"`
	Description string `json:"description" validate:"string.(min=1,max=255)" select:"description" insert:"description"`
	PictureURL  string `json:"picture_url" validate:"string.(min=1,max=255)" select:"picture_url" insert:"picture_url"`

	RewardTypeID int    `json:"reward_type_id" validate:"id" select:"reward_type_id" insert:"reward_type_id"`
	UserID       string `json:"user_id"        validate:"-"  select:"user_id"        insert:"user_id"`

	CreatedAt time.Time `json:"created_at" select:"created_at"`
	UpdatedAt time.Time `json:"updated_at" select:"updated_at"`
	DeletedAt time.Time `json:"deleted_at" select:"deleted_at"`
}
