package mock

import (
	"github.com/ivzb/achievers_server/app/model"
)

type Profile struct {
	ExistsMock         ProfileExists
	SingleMock         ProfileSingle
	SingleByUserIDMock ProfileSingleByUserID
	CreateMock         ProfileCreate
}

type ProfileExists struct {
	Bool bool
	Err  error
}

type ProfileSingle struct {
	Prfl *model.Profile
	Err  error
}

type ProfileSingleByUserID struct {
	Prfl *model.Profile
	Err  error
}

type ProfileCreate struct {
	ID  string
	Err error
}

func (ctx *Profile) Exists(id string) (bool, error) {
	return ctx.ExistsMock.Bool, ctx.ExistsMock.Err
}

func (ctx *Profile) Single(id string) (*model.Profile, error) {
	return ctx.SingleMock.Prfl, ctx.SingleMock.Err
}

func (ctx *Profile) SingleByUserID(userID string) (*model.Profile, error) {
	return ctx.SingleByUserIDMock.Prfl, ctx.SingleByUserIDMock.Err
}

func (ctx *Profile) Create(profile *model.Profile, userID string) (string, error) {
	return ctx.CreateMock.ID, ctx.CreateMock.Err
}
