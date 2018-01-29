package db

import "github.com/ivzb/achievers_server/app/model"

type QuestAchievementer interface {
	Exists(questID string, achievementID string) (bool, error)
	Create(qstAch *model.QuestAchievement) (string, error)
}

type QuestAchievement struct {
	*Context
}

func (db *DB) QuestAchievement() QuestAchievementer {
	return &QuestAchievement{
		&Context{
			db:         db,
			table:      "quest_achievement",
			insertArgs: "quest_id, achievement_id, user_id",
		},
	}
}

func (ctx *QuestAchievement) Exists(questID string, achievementID string) (bool, error) {
	return existsMultiple(ctx.db, "quest_achievement", []string{"quest_id", "achievement_id"}, []string{questID, achievementID})
}

func (ctx *QuestAchievement) Create(qstAch *model.QuestAchievement) (string, error) {
	return create(ctx.Context,
		qstAch.QuestID,
		qstAch.AchievementID,
		qstAch.UserID)
}
