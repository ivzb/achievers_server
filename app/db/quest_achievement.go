package db

import (
	"github.com/ivzb/achievers_server/app/model"
	"github.com/ivzb/achievers_server/app/shared/consts"
)

type QuestAchievementer interface {
	Exists(questID string, achievementID string) (bool, error)
	Create(qstAch *model.QuestAchievement) (string, error)
}

type QuestAchievement struct {
	*Context
}

func (db *DB) QuestAchievement() QuestAchievementer {
	return &QuestAchievement{
		newContext(db, consts.QuestAchievement, new(model.QuestAchievement)),
	}
}

func (ctx *QuestAchievement) Exists(questID string, achievementID string) (bool, error) {
	keys := []string{consts.QuestID, consts.AchievementID}
	values := []string{questID, achievementID}

	return ctx.existsMultiple(keys, values)
}

func (ctx *QuestAchievement) Create(qstAch *model.QuestAchievement) (string, error) {
	return ctx.create(
		qstAch.QuestID,
		qstAch.AchievementID,
		qstAch.UserID)
}
