package mock

import (
	"github.com/ivzb/achievers_server/app/db"
	"github.com/ivzb/achievers_server/app/model"
)

type QuestAchievementExists struct {
	Bool bool
	Err  error
}

type QuestAchievementCreate struct {
	ID  string
	Err error
}

type QuestAchievement struct {
	db *DB

	ExistsMock QuestAchievementExists
	CreateMock QuestAchievementCreate
}

func (db *DB) QuestAchievement() db.QuestAchievementer {
	return &QuestAchievement{db: db}
}

func (ctx *QuestAchievement) Exists(questID string, achievementID string) (bool, error) {
	return ctx.ExistsMock.Bool, ctx.ExistsMock.Err
}

func (ctx *QuestAchievement) Create(qstAch *model.QuestAchievement) (string, error) {
	return ctx.CreateMock.ID, ctx.CreateMock.Err
}

//func QuestAchievements(size int) []*model.QuestAchievement {
//qstAchs := make([]*model.QuestAchievement, size)

//for i := 0; i < size; i++ {
//qstAchs[i] = QuestAchievement()
//}

//return qstAchs
//}

//func QuestAchievement() *model.QuestAchievement {
//qstAch := &model.QuestAchievement{
//"742127e4-6689-0c27-c319-d9952b713b8d",
//"fd9d70c9-f3aa-8f53-0946-b889a57ef22d",
//"a247c0d9-53bb-e6bf-4828-5ad5804570c4",
//"4e69c9ba-719c-ca7c-fb66-80ab235c8e39",
//time.Date(2017, 12, 9, 15, 4, 23, 0, time.UTC),
//time.Date(2017, 12, 9, 15, 4, 23, 0, time.UTC),
//time.Date(0, 1, 1, 0, 0, 0, 0, time.UTC),
//}

//return qstAch
//}

//func (mock *DB) QuestAchievementExists(string, string) (bool, error) {
//return mock.QuestAchievementExistsMock.Bool, mock.QuestAchievementExistsMock.Err
//}

//func (mock *DB) QuestAchievementCreate(questAchievement *model.QuestAchievement) (string, error) {
//return mock.QuestAchievementCreateMock.ID, mock.QuestAchievementCreateMock.Err
//}
