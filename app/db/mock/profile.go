package mock

import (
	"github.com/ivzb/achievers_server/app/db"
	"github.com/ivzb/achievers_server/app/model"
)

type Profile struct {
	db *DB

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

func (db *DB) Profile() db.Profiler {
	return &Profile{db: db}
}

func (ctx *Profile) Exists(id string) (bool, error) {
	return ctx.ExistsMock.Bool, ctx.ExistsMock.Err
}

func (ctx *Profile) Single(id string) (*model.Profile, error) {
	return ctx.SingleMock.Prfl, ctx.SingleMock.Err
}

func (ctx *Profile) SingleByUserID(userID string) (*model.Profile, error) {
	return ctx.SingleByUserIDMock.Prfl, ctx.SingleMock.Err
}

func (ctx *Profile) Create(profile *model.Profile, userID string) (string, error) {
	return ctx.CreateMock.ID, ctx.CreateMock.Err
}
