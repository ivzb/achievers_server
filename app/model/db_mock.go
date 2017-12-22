package model

type ExistsMock struct {
	B bool
	E error
}

type AchievementsAllMock struct {
	A []*Achievement
	E error
}

type UserCreateMock struct {
	S string
	E error
}

type UserAuthMock struct {
	S string
	E error
}

type DBMock struct {
	ExistsMock          ExistsMock
	AchievementsAllMock AchievementsAllMock
	UserCreateMock      UserCreateMock
	UserAuthMock        UserAuthMock
}

func (mock *DBMock) Exists(string, string, string) (bool, error) {
	return mock.ExistsMock.B, mock.ExistsMock.E
}

func (mock *DBMock) AchievementsAll() ([]*Achievement, error) {
	return mock.AchievementsAllMock.A, mock.AchievementsAllMock.E
}

func (mock *DBMock) UserCreate(string, string, string, string) (string, error) {
	return mock.UserCreateMock.S, mock.UserCreateMock.E
}

func (mock *DBMock) UserAuth(string, string) (string, error) {
	return mock.UserAuthMock.S, mock.UserAuthMock.E
}
