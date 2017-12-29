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

type AchievementExists struct {
	Bool bool
	Err  error
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

type EvidenceExists struct {
	Bool bool
	Err  error
}

type EvidenceSingle struct {
	Evd *model.Evidence
	Err error
}

type EvidenceCreate struct {
	ID  string
	Err error
}

type InvolvementExists struct {
	Bool bool
	Err  error
}

type MultimediaTypeExists struct {
	Bool bool
	Err  error
}

type DB struct {
	UserExistsMock      UserExists
	UserEmailExistsMock UserEmailExists
	UserCreateMock      UserCreate
	UserAuthMock        UserAuth

	AchievementExistsMock AchievementExists
	AchievementSingleMock AchievementSingle
	AchievementsAllMock   AchievementsAll
	AchievementCreateMock AchievementCreate

	EvidenceExistsMock EvidenceExists
	EvidenceSingleMock EvidenceSingle
	EvidenceCreateMock EvidenceCreate

	InvolvementExistsMock InvolvementExists

	MultimediaTypeExistsMock MultimediaTypeExists
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

func (mock *DB) UserAuth(string, string) (string, error) {
	return mock.UserAuthMock.ID, mock.UserAuthMock.Err
}

func (mock *DB) AchievementExists(string) (bool, error) {
	return mock.AchievementExistsMock.Bool, mock.AchievementExistsMock.Err
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

func (mock *DB) EvidenceExists(string) (bool, error) {
	return mock.EvidenceExistsMock.Bool, mock.EvidenceExistsMock.Err
}

func (mock *DB) EvidenceSingle(id string) (*model.Evidence, error) {
	return mock.EvidenceSingleMock.Evd, mock.EvidenceSingleMock.Err
}

func (mock *DB) EvidenceCreate(evidence *model.Evidence) (string, error) {
	return mock.EvidenceCreateMock.ID, mock.EvidenceCreateMock.Err
}

func (mock *DB) InvolvementExists(string) (bool, error) {
	return mock.InvolvementExistsMock.Bool, mock.InvolvementExistsMock.Err
}

func (mock *DB) MultimediaTypeExists(uint8) (bool, error) {
	return mock.MultimediaTypeExistsMock.Bool, mock.MultimediaTypeExistsMock.Err
}
