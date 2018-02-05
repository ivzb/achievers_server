package mock

type QuestAchievement struct {
	ExistsMock QuestAchievementExists
	CreateMock QuestAchievementCreate
}

type QuestAchievementExists struct {
	Bool bool
	Err  error
}

type QuestAchievementCreate struct {
	ID  string
	Err error
}

func (ctx *QuestAchievement) Exists(...interface{}) (bool, error) {
	return ctx.ExistsMock.Bool, ctx.ExistsMock.Err
}

func (ctx *QuestAchievement) Create(qstAch interface{}) (string, error) {
	return ctx.CreateMock.ID, ctx.CreateMock.Err
}
