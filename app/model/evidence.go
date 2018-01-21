package model

import "time"

type Evidence struct {
	ID string `json:"id"`

	Title      string `json:"title"`
	PictureURL string `json:"picture_url"`
	URL        string `json:"url"`

	MultimediaTypeID uint8  `json:"multimedia_type_id"`
	AchievementID    string `json:"achievement_id"`
	UserID           string `json:"user_id"`

	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	DeletedAt time.Time `json:"deleted_at"`
}
