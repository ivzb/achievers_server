package mock

import (
	"time"

	"github.com/ivzb/achievers_server/app/model"
)

type ProfileExists struct {
	Bool bool
	Err  error
}

type ProfileSingle struct {
	Prfl *model.Profile
	Err  error
}

type ProfileByUserID struct {
	Prfl *model.Profile
	Err  error
}

type ProfileCreate struct {
	ID  string
	Err error
}

func Profile() *model.Profile {
	prfl := &model.Profile{
		"fb7691eb-ea1d-b20f-edee-9cadcf23181f",
		"Ivan Zahariev",
		time.Date(2017, 12, 9, 15, 4, 23, 0, time.UTC),
		time.Date(2017, 12, 9, 15, 4, 23, 0, time.UTC),
		time.Date(0, 1, 1, 0, 0, 0, 0, time.UTC),
	}

	return prfl
}

func (mock *DB) ProfileExists(string) (bool, error) {
	return mock.ProfileExistsMock.Bool, mock.ProfileExistsMock.Err
}

func (mock *DB) ProfileSingle(id string) (*model.Profile, error) {
	return mock.ProfileSingleMock.Prfl, mock.ProfileSingleMock.Err
}

func (mock *DB) ProfileByUserID(userID string) (*model.Profile, error) {
	return mock.ProfileByUserIDMock.Prfl, mock.ProfileByUserIDMock.Err
}

func (mock *DB) ProfileCreate(profile *model.Profile, userID string) (string, error) {
	return mock.ProfileCreateMock.ID, mock.ProfileCreateMock.Err
}
