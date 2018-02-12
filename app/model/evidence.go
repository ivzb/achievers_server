package model

import "time"

type Evidence struct {
	ID string `json:"id" select:"id" exists:"id"`

	Title      string `json:"title"       validate:"string.(min=1,max=255)" select:"title"       insert:"title"`
	PictureURL string `json:"picture_url" validate:"string.(min=1,max=255)" select:"picture_url" insert:"picture_url"`
	URL        string `json:"url"         validate:"string.(min=1,max=255)" select:"url"         insert:"url"`

	MultimediaTypeID int    `json:"multimedia_type_id" validate:"id"   select:"multimedia_type_id" insert:"multimedia_type_id"`
	AchievementID    string `json:"achievement_id"     validate:"uuid" select:"achievement_id"     insert:"achievement_id"`
	UserID           string `json:"user_id"            validate:"-"    select:"user_id"            insert:"user_id"`

	CreatedAt time.Time `json:"created_at" select:"created_at"`
	UpdatedAt time.Time `json:"updated_at" select:"updated_at"`
	DeletedAt time.Time `json:"deleted_at" select:"deleted_at"`
}
