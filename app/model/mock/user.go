package mock

import "github.com/ivzb/achievers_server/app/model"

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

func (mock *DB) UserExists(string) (bool, error) {
	return mock.UserExistsMock.Bool, mock.UserExistsMock.Err
}

func (mock *DB) UserEmailExists(string) (bool, error) {
	return mock.UserEmailExistsMock.Bool, mock.UserEmailExistsMock.Err
}

func (mock *DB) UserCreate(user *model.User) (string, error) {
	return mock.UserCreateMock.ID, mock.UserCreateMock.Err
}

func (mock *DB) UserAuth(auth *model.Auth) (string, error) {
	return mock.UserAuthMock.ID, mock.UserAuthMock.Err
}
