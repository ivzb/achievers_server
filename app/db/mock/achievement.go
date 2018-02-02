package mock

type Achievement struct {
	ExistsMock          AchievementExists
	SingleMock          AchievementSingle
	CreateMock          AchievementCreate
	LastIDMock          AchievementsLastID
	LastIDByQuestIDMock AchievementsLastIDByQuestID
	AfterMock           AchievementsAfter
	AfterByQuestIDMock  AchievementsAfterByQuestID
}

type AchievementExists struct {
	Bool bool
	Err  error
}

type AchievementSingle struct {
	Ach interface{}
	Err error
}

type AchievementsLastID struct {
	ID  string
	Err error
}

type AchievementsLastIDByQuestID struct {
	ID  string
	Err error
}

type AchievementsAfter struct {
	Achs []interface{}
	Err  error
}

type AchievementsAfterByQuestID struct {
	Achs []interface{}
	Err  error
}

type AchievementCreate struct {
	ID  string
	Err error
}

func (ctx *Achievement) Exists(id string) (bool, error) {
	return ctx.ExistsMock.Bool, ctx.ExistsMock.Err
}

func (ctx *Achievement) Single(id string) (interface{}, error) {
	return ctx.SingleMock.Ach, ctx.SingleMock.Err
}

func (ctx *Achievement) Create(achievement interface{}) (string, error) {
	return ctx.CreateMock.ID, ctx.CreateMock.Err
}

func (ctx *Achievement) LastID() (string, error) {
	return ctx.LastIDMock.ID, ctx.LastIDMock.Err
}

func (ctx *Achievement) LastIDByQuestID(questID string) (string, error) {
	return ctx.LastIDByQuestIDMock.ID, ctx.LastIDByQuestIDMock.Err
}

func (ctx *Achievement) After(afterID string) ([]interface{}, error) {
	return ctx.AfterMock.Achs, ctx.AfterMock.Err
}

func (ctx *Achievement) AfterByQuestID(questID string, afterID string) ([]interface{}, error) {
	return ctx.AfterByQuestIDMock.Achs, ctx.AfterByQuestIDMock.Err
}
