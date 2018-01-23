package mock

import (
	"github.com/ivzb/achievers_server/app/model"
)

type Quest struct {
	ExistsMock QuestExists
	SingleMock QuestSingle
	AllMock    QuestsAll
	CreateMock QuestCreate
}

type QuestExists struct {
	Bool bool
	Err  error
}

type QuestSingle struct {
	Qst *model.Quest
	Err error
}

type QuestsAll struct {
	Qsts []*model.Quest
	Err  error
}

type QuestCreate struct {
	ID  string
	Err error
}

func (ctx *Quest) Exists(id string) (bool, error) {
	return ctx.ExistsMock.Bool, ctx.ExistsMock.Err
}

func (ctx *Quest) Single(id string) (*model.Quest, error) {
	return ctx.SingleMock.Qst, ctx.SingleMock.Err
}

func (ctx *Quest) All(page int) ([]*model.Quest, error) {
	return ctx.AllMock.Qsts, ctx.AllMock.Err
}

func (ctx *Quest) Create(quest *model.Quest) (string, error) {
	return ctx.CreateMock.ID, ctx.CreateMock.Err
}
