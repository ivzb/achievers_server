package mock

import (
	"github.com/ivzb/achievers_server/app/model"
)

type Quest struct {
	ExistsMock QuestExists
	SingleMock QuestSingle
	CreateMock QuestCreate
	LastIDMock QuestsLastID
	AfterMock  QuestsAfter
}

type QuestExists struct {
	Bool bool
	Err  error
}

type QuestSingle struct {
	Qst *model.Quest
	Err error
}

type QuestCreate struct {
	ID  string
	Err error
}

type QuestsLastID struct {
	ID  string
	Err error
}

type QuestsAfter struct {
	Qsts []*model.Quest
	Err  error
}

func (ctx *Quest) Exists(id string) (bool, error) {
	return ctx.ExistsMock.Bool, ctx.ExistsMock.Err
}

func (ctx *Quest) Single(id string) (*model.Quest, error) {
	return ctx.SingleMock.Qst, ctx.SingleMock.Err
}

func (ctx *Quest) Create(quest *model.Quest) (string, error) {
	return ctx.CreateMock.ID, ctx.CreateMock.Err
}

func (ctx *Quest) LastID() (string, error) {
	return ctx.LastIDMock.ID, ctx.LastIDMock.Err
}

func (ctx *Quest) After(afterID string) ([]*model.Quest, error) {
	return ctx.AfterMock.Qsts, ctx.AfterMock.Err
}
