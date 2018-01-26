package mock

import (
	"github.com/ivzb/achievers_server/app/model"
)

type Reward struct {
	ExistsMock RewardExists
	SingleMock RewardSingle
	CreateMock RewardCreate
	LastIDMock RewardsLastID
	AfterMock  RewardsAfter
}

type RewardExists struct {
	Bool bool
	Err  error
}

type RewardSingle struct {
	Rwd *model.Reward
	Err error
}

type RewardsLastID struct {
	ID  string
	Err error
}

type RewardsAfter struct {
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

func (ctx *Reward) Create(reward *model.Reward) (string, error) {
	return ctx.CreateMock.ID, ctx.CreateMock.Err
}

func (ctx *Reward) LastID() (string, error) {
	return ctx.LastIDMock.ID, ctx.LastIDMock.Err
}

func (ctx *Reward) After(afterID string) ([]*model.Reward, error) {
	return ctx.AfterMock.Rwds, ctx.AfterMock.Err
}
