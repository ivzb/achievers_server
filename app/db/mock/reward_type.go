package mock

type RewardType struct {
	ExistsMock RewardTypeExists
}

type RewardTypeExists struct {
	Bool bool
	Err  error
}

func (ctx *RewardType) Exists(id uint8) (bool, error) {
	return ctx.ExistsMock.Bool, ctx.ExistsMock.Err
}
