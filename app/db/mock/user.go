package mock

import (
	"github.com/ivzb/achievers_server/app/db"
	"github.com/ivzb/achievers_server/app/model"
)

type User struct {
	db *DB

	ExistsMock      UserExists
	EmailExistsMock UserEmailExists
	AuthMock        UserAuth
	CreateMock      UserCreate
}

func (db *DB) User() db.Userer {
	return &User{db: db}
}

func (ctx *User) Exists(id string) (bool, error) {
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

//func (mock *DB) UserExists(string) (bool, error) {
//return mock.UserExistsMock.Bool, mock.UserExistsMock.Err
//}

//func (mock *DB) UserEmailExists(string) (bool, error) {
//return mock.UserEmailExistsMock.Bool, mock.UserEmailExistsMock.Err
//}

//func (mock *DB) UserCreate(user *model.User) (string, error) {
//return mock.UserCreateMock.ID, mock.UserCreateMock.Err
//}

//func (mock *DB) UserAuth(auth *model.Auth) (string, error) {
//return mock.UserAuthMock.ID, mock.UserAuthMock.Err
//}
