package mock

import (
	"github.com/ivzb/achievers_server/app/db"
	"github.com/ivzb/achievers_server/app/model"
)

type Achievement struct {
	db *DB

	ExistsMock          AchievementExists
	SingleMock          AchievementSingle
	CreateMock          AchievementCreate
	LastIDMock          AchievementsLastID
	LastIDByQuestIDMock AchievementsLastIDByQuestID
	AfterMock           AchievementsAfter
	AfterByQuestIDMock  AchievementsAfterByQuestID
}

type AchievementExists struct {
	Bool bool
	Err  error
}

type AchievementSingle struct {
	Ach *model.Achievement
	Err error
}

type AchievementsLastID struct {
	ID  string
	Err error
}

type AchievementsLastIDByQuestID struct {
	ID  string
	Err error
}

type AchievementsAfter struct {
	Achs []*model.Achievement
	Err  error
}

type AchievementsAfterByQuestID struct {
	Achs []*model.Achievement
	Err  error
}

type AchievementCreate struct {
	ID  string
	Err error
}

func (db *DB) Achievement() db.Achievementer {
	return &Achievement{db: db}
}

func (ctx *Achievement) Exists(id string) (bool, error) {
	return ctx.ExistsMock.Bool, ctx.ExistsMock.Err
}

func (ctx *Achievement) Single(id string) (*model.Achievement, error) {
	return ctx.SingleMock.Ach, ctx.SingleMock.Err
}

func (ctx *Achievement) Create(achievement *model.Achievement) (string, error) {
	return ctx.CreateMock.ID, ctx.CreateMock.Err
}

func (ctx *Achievement) LastID() (string, error) {
	return ctx.LastIDMock.ID, ctx.LastIDMock.Err
}

func (ctx *Achievement) LastIDByQuestID(questID string) (string, error) {
	return ctx.LastIDByQuestIDMock.ID, ctx.LastIDByQuestIDMock.Err
}

func (ctx *Achievement) After(afterID string) ([]*model.Achievement, error) {
	return ctx.AfterMock.Achs, ctx.AfterMock.Err
}

func (ctx *Achievement) AfterByQuestID(questID string, afterID string) ([]*model.Achievement, error) {
	return ctx.AfterByQuestIDMock.Achs, ctx.AfterByQuestIDMock.Err
}
