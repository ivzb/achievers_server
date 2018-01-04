package mock

import (
	"time"

	"github.com/ivzb/achievers_server/app/model"
)

type QuestExists struct {
	Bool bool
	Err  error
}

type QuestSingle struct {
	Qst *model.Quest
	Err error
}

type QuestsAll struct {
	Qsts []*model.Quest
	Err  error
}

type QuestCreate struct {
	ID  string
	Err error
}

func Quests(size int) []*model.Quest {
	qsts := make([]*model.Quest, size)

	for i := 0; i < size; i++ {
		qsts[i] = Quest()
	}

	return qsts
}

func Quest() *model.Quest {
	qst := &model.Quest{
		"fb7691eb-ea1d-b20f-edee-9cadcf23181f",
		"name",
		"http://picture.jpg",
		"3",
		3,
		"4e69c9ba-719c-ca7c-fb66-80ab235c8e39",
		time.Date(2017, 12, 9, 15, 4, 23, 0, time.UTC),
		time.Date(2017, 12, 9, 15, 4, 23, 0, time.UTC),
		time.Date(0, 1, 1, 0, 0, 0, 0, time.UTC),
	}

	return qst
}

func (mock *DB) QuestExists(string) (bool, error) {
	return mock.QuestExistsMock.Bool, mock.QuestExistsMock.Err
}

func (mock *DB) QuestSingle(id string) (*model.Quest, error) {
	return mock.QuestSingleMock.Qst, mock.QuestSingleMock.Err
}

func (mock *DB) QuestsAll(page int) ([]*model.Quest, error) {
	return mock.QuestsAllMock.Qsts, mock.QuestsAllMock.Err
}

func (mock *DB) QuestCreate(reward *model.Quest) (string, error) {
	return mock.QuestCreateMock.ID, mock.QuestCreateMock.Err
}
