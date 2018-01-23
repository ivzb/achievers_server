package mock

import (
	"github.com/ivzb/achievers_server/app/model"
)

type Reward struct {
	ExistsMock RewardExists
	SingleMock RewardSingle
	AllMock    RewardsAll
	CreateMock RewardCreate
}

type RewardExists struct {
	Bool bool
	Err  error
}

type RewardSingle struct {
	Rwd *model.Reward
	Err error
}

type RewardsAll struct {
	Rwds []*model.Reward
	Err  error
}

type RewardCreate struct {
	ID  string
	Err error
}

func (ctx *Reward) Exists(id string) (bool, error) {
	return ctx.ExistsMock.Bool, ctx.ExistsMock.Err
}

func (ctx *Reward) Single(id string) (*model.Reward, error) {
	return ctx.SingleMock.Rwd, ctx.SingleMock.Err
}

func (ctx *Reward) All(page int) ([]*model.Reward, error) {
	return ctx.AllMock.Rwds, ctx.AllMock.Err
}

func (ctx *Reward) Create(reward *model.Reward) (string, error) {
	return ctx.CreateMock.ID, ctx.CreateMock.Err
}
