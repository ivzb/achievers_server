package model

import (
	"time"
)

type QuestAchievement struct {
	ID string `json:"id"`

	QuestID       string `json:"quest_id"`
	AchievementID string `json:"achievement_id"`
	AuthorID      string `json:"author_id"`

	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	DeletedAt time.Time `json:"deleted_at"`
}

func (db *DB) QuestAchievementExists(questID string, achievementID string) (bool, error) {
	return existsMultiple(db, "quest_achievement", []string{"quest_id", "achievement_id"}, []string{questID, achievementID})
}

func (db *DB) QuestAchievementCreate(qstAch *QuestAchievement) (string, error) {
	return create(db, `INSERT INTO quest_achievement (id, quest_id, achievement_id, author_id)
        VALUES(?, ?, ?, ?)`,
		qstAch.QuestID,
		qstAch.AchievementID,
		qstAch.AuthorID)
}
