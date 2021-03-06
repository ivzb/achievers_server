package mock

type Quest struct {
	ExistsMock QuestExists
	SingleMock QuestSingle
	CreateMock QuestCreate
	LastIDMock QuestsLastID
	AfterMock  QuestsAfter
}

type QuestExists struct {
	Bool bool
	Err  error
}

type QuestSingle struct {
	Qst interface{}
	Err error
}

type QuestCreate struct {
	ID  string
	Err error
}

type QuestsLastID struct {
	ID  string
	Err error
}

type QuestsAfter struct {
	Qsts []interface{}
	Err  error
}

func (ctx *Quest) Exists(id interface{}) (bool, error) {
	return ctx.ExistsMock.Bool, ctx.ExistsMock.Err
}

func (ctx *Quest) Single(id string) (interface{}, error) {
	return ctx.SingleMock.Qst, ctx.SingleMock.Err
}

func (ctx *Quest) Create(quest interface{}) (string, error) {
	return ctx.CreateMock.ID, ctx.CreateMock.Err
}

func (ctx *Quest) LastID() (string, error) {
	return ctx.LastIDMock.ID, ctx.LastIDMock.Err
}

func (ctx *Quest) After(afterID string) ([]interface{}, error) {
	return ctx.AfterMock.Qsts, ctx.AfterMock.Err
}
