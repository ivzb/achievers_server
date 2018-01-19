package mock

import (
	"time"

	"github.com/ivzb/achievers_server/app/model"
)

type AchievementExists struct {
	Bool bool
	Err  error
}

type AchievementSingle struct {
	Ach *model.Achievement
	Err error
}

type AchievementsLastID struct {
	ID  string
	Err error
}

type AchievementsAfter struct {
	Achs []*model.Achievement
	Err  error
}

type AchievementsByQuestIDLastID struct {
	ID  string
	Err error
}

type AchievementsByQuestIDAfter struct {
	Achs []*model.Achievement
	Err  error
}

type AchievementCreate struct {
	ID  string
	Err error
}

func Achievements(size int) []*model.Achievement {
	achs := make([]*model.Achievement, size)

	for i := 0; i < size; i++ {
		achs[i] = Achievement()
	}

	return achs
}

func Achievement() *model.Achievement {
	ach := &model.Achievement{
		"fb7691eb-ea1d-b20f-edee-9cadcf23181f",
		"title",
		"desc",
		"http://picture.jpg",
		"3",
		"4e69c9ba-719c-ca7c-fb66-80ab235c8e39",
		time.Date(2017, 12, 9, 15, 4, 23, 0, time.UTC),
		time.Date(2017, 12, 9, 15, 4, 23, 0, time.UTC),
		time.Date(0, 1, 1, 0, 0, 0, 0, time.UTC),
	}

	return ach
}

func (mock *DB) AchievementExists(string) (bool, error) {
	return mock.AchievementExistsMock.Bool, mock.AchievementExistsMock.Err
}

func (mock *DB) AchievementSingle(id string) (*model.Achievement, error) {
	return mock.AchievementSingleMock.Ach, mock.AchievementSingleMock.Err
}

func (mock *DB) AchievementsLastID() (string, error) {
	return mock.AchievementsLastIDMock.ID, mock.AchievementsLastIDMock.Err
}

func (mock *DB) AchievementsAfter(id string) ([]*model.Achievement, error) {
	return mock.AchievementsAfterMock.Achs, mock.AchievementsAfterMock.Err
}

func (mock *DB) AchievementsByQuestIDLastID(questID string) (string, error) {
	return mock.AchievementsByQuestIDLastIDMock.ID, mock.AchievementsByQuestIDLastIDMock.Err
}

func (mock *DB) AchievementsByQuestIDAfter(questID string, afterID string) ([]*model.Achievement, error) {
	return mock.AchievementsByQuestIDAfterMock.Achs, mock.AchievementsByQuestIDAfterMock.Err
}

func (mock *DB) AchievementCreate(achievement *model.Achievement) (string, error) {
	return mock.AchievementCreateMock.ID, mock.AchievementCreateMock.Err
}
