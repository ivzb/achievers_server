package mock

type Involvement struct {
	ExistsMock InvolvementExists
}

type InvolvementExists struct {
	Bool bool
	Err  error
}

func (ctx *Involvement) Exists(id string) (bool, error) {
	return ctx.ExistsMock.Bool, ctx.ExistsMock.Err
}
