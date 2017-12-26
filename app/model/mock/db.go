package mock

import "github.com/ivzb/achievers_server/app/model"

type Exists struct {
	B bool
	E error
}

type AchievementSingle struct {
	A *model.Achievement
	E error
}

type AchievementsAll struct {
	A []*model.Achievement
	E error
}

type AchievementCreate struct {
	A string
	E error
}

type UserCreate struct {
	S string
	E error
}

type UserAuth struct {
	S string
	E error
}

type DB struct {
	ExistsMock            Exists
	AchievementSingleMock AchievementSingle
	AchievementsAllMock   AchievementsAll
	AchievementCreateMock AchievementCreate
	UserCreateMock        UserCreate
	UserAuthMock          UserAuth
}

func (mock *DB) Exists(string, string, string) (bool, error) {
	return mock.ExistsMock.B, mock.ExistsMock.E
}

func (mock *DB) AchievementSingle(id string) (*model.Achievement, error) {
	return mock.AchievementSingleMock.A, mock.AchievementSingleMock.E
}

func (mock *DB) AchievementsAll(page int) ([]*model.Achievement, error) {
	return mock.AchievementsAllMock.A, mock.AchievementsAllMock.E
}

func (mock *DB) AchievementCreate(achievement *model.Achievement) (string, error) {
	return mock.AchievementCreateMock.A, mock.AchievementCreateMock.E
}

func (mock *DB) UserCreate(string, string, string, string) (string, error) {
	return mock.UserCreateMock.S, mock.UserCreateMock.E
}

func (mock *DB) UserAuth(string, string) (string, error) {
	return mock.UserAuthMock.S, mock.UserAuthMock.E
}
