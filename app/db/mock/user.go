package mock

import (
	"github.com/ivzb/achievers_server/app/model"
)

type User struct {
	ExistsMock      UserExists
	EmailExistsMock UserEmailExists
	AuthMock        UserAuth
	CreateMock      UserCreate
}

type UserExists struct {
	Bool bool
	Err  error
}

type UserEmailExists struct {
	Bool bool
	Err  error
}

type UserCreate struct {
	ID  string
	Err error
}

type UserAuth struct {
	ID  string
	Err error
}

func (ctx *User) Exists(id interface{}) (bool, error) {
	return ctx.ExistsMock.Bool, ctx.ExistsMock.Err
}

func (ctx *User) EmailExists(email string) (bool, error) {
	return ctx.EmailExistsMock.Bool, ctx.EmailExistsMock.Err
}

func (ctx *User) Auth(auth *model.Auth) (string, error) {
	return ctx.AuthMock.ID, ctx.AuthMock.Err
}

func (ctx *User) Create(user *model.User) (string, error) {
	return ctx.CreateMock.ID, ctx.CreateMock.Err
}
