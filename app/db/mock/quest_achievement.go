package mock

import (
	"github.com/ivzb/achievers_server/app/model"
)

type QuestAchievement struct {
	ExistsMock QuestAchievementExists
	CreateMock QuestAchievementCreate
}

type QuestAchievementExists struct {
	Bool bool
	Err  error
}

type QuestAchievementCreate struct {
	ID  string
	Err error
}

func (ctx *QuestAchievement) Exists(questID string, achievementID string) (bool, error) {
	return ctx.ExistsMock.Bool, ctx.ExistsMock.Err
}

func (ctx *QuestAchievement) Create(qstAch *model.QuestAchievement) (string, error) {
	return ctx.CreateMock.ID, ctx.CreateMock.Err
}
