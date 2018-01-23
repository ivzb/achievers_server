package db

import "github.com/ivzb/achievers_server/app/model"

type QuestAchievementer interface {
	Exists(questID string, achievementID string) (bool, error)
	Create(qstAch *model.QuestAchievement) (string, error)
}

type QuestAchievement struct {
	db *DB
}

func (db *DB) QuestAchievement() QuestAchievementer {
	return &QuestAchievement{db}
}

func (ctx *QuestAchievement) Exists(questID string, achievementID string) (bool, error) {
	return existsMultiple(ctx.db, "quest_achievement", []string{"quest_id", "achievement_id"}, []string{questID, achievementID})
}

func (ctx *QuestAchievement) Create(qstAch *model.QuestAchievement) (string, error) {
	return create(ctx.db, `INSERT INTO quest_achievement (id, quest_id, achievement_id, user_id)
		VALUES($1, $2, $3, $4)`,
		qstAch.QuestID,
		qstAch.AchievementID,
		qstAch.UserID)
}
