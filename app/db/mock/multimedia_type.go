package mock

type MultimediaType struct {
	ExistsMock MultimediaTypeExists
}

type MultimediaTypeExists struct {
	Bool bool
	Err  error
}

func (ctx *MultimediaType) Exists(id interface{}) (bool, error) {
	return ctx.ExistsMock.Bool, ctx.ExistsMock.Err
}
