package model

import (
	"time"
)

type QuestAchievement struct {
	ID string `json:"id"`

	QuestID       string `json:"quest_id" insert:"quest_id"`
	AchievementID string `json:"achievement_id" insert:"achievement_id"`
	UserID        string `json:"user_id" insert:"user_id"`

	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	DeletedAt time.Time `json:"deleted_at"`
}
