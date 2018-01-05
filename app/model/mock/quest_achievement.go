package mock

import "github.com/ivzb/achievers_server/app/model"

type QuestAchievementExists struct {
	Bool bool
	Err  error
}

type QuestAchievementSingle struct {
	QstAch *model.QuestAchievement
	Err    error
}

func (mock *DB) QuestAchievementExists(string, string) (bool, error) {
	return mock.QuestAchievementExistsMock.Bool, mock.QuestAchievementExistsMock.Err
}

func (mock *DB) QuestAchievementSingle(string, string) (*model.QuestAchievement, error) {
	return mock.QuestAchievementSingleMock.QstAch, mock.QuestAchievementSingleMock.Err
}
