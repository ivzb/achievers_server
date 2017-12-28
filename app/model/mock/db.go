package mock

import "github.com/ivzb/achievers_server/app/model"

type Exists struct {
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

type AchievementSingle struct {
	Ach *model.Achievement
	Err error
}

type AchievementsAll struct {
	Achs []*model.Achievement
	Err  error
}

type AchievementCreate struct {
	ID  string
	Err error
}

type EvidenceCreate struct {
	ID  string
	Err error
}

type DB struct {
	ExistsMock            Exists
	UserCreateMock        UserCreate
	UserAuthMock          UserAuth
	AchievementSingleMock AchievementSingle
	AchievementsAllMock   AchievementsAll
	AchievementCreateMock AchievementCreate
	EvidenceCreateMock    EvidenceCreate
}

func (mock *DB) Exists(string, string, string) (bool, error) {
	return mock.ExistsMock.Bool, mock.ExistsMock.Err
}

func (mock *DB) UserCreate(user *model.User) (string, error) {
	return mock.UserCreateMock.ID, mock.UserCreateMock.Err
}

func (mock *DB) UserAuth(string, string) (string, error) {
	return mock.UserAuthMock.ID, mock.UserAuthMock.Err
}

func (mock *DB) AchievementSingle(id string) (*model.Achievement, error) {
	return mock.AchievementSingleMock.Ach, mock.AchievementSingleMock.Err
}

func (mock *DB) AchievementsAll(page int) ([]*model.Achievement, error) {
	return mock.AchievementsAllMock.Achs, mock.AchievementsAllMock.Err
}

func (mock *DB) AchievementCreate(achievement *model.Achievement) (string, error) {
	return mock.AchievementCreateMock.ID, mock.AchievementCreateMock.Err
}

func (mock *DB) EvidenceCreate(evidence *model.Evidence) (string, error) {
	return mock.EvidenceCreateMock.ID, mock.EvidenceCreateMock.Err
}
