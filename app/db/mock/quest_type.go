package mock

type QuestType struct {
	ExistsMock QuestTypeExists
}

type QuestTypeExists struct {
	Bool bool
	Err  error
}

func (ctx *QuestType) Exists(id interface{}) (bool, error) {
	return ctx.ExistsMock.Bool, ctx.ExistsMock.Err
}
