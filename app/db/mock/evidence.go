package mock

type Evidence struct {
	ExistsMock EvidenceExists
	SingleMock EvidenceSingle
	CreateMock EvidenceCreate
	LastIDMock EvidencesLastID
	AfterMock  EvidencesAfter
}

type EvidenceExists struct {
	Bool bool
	Err  error
}

type EvidenceSingle struct {
	Evd interface{}
	Err error
}

type EvidencesLastID struct {
	ID  string
	Err error
}

type EvidencesAfter struct {
	Evds []interface{}
	Err  error
}

type EvidenceCreate struct {
	ID  string
	Err error
}

func (ctx *Evidence) Exists(id string) (bool, error) {
	return ctx.ExistsMock.Bool, ctx.ExistsMock.Err
}

func (ctx *Evidence) Single(id string) (interface{}, error) {
	return ctx.SingleMock.Evd, ctx.SingleMock.Err
}

func (ctx *Evidence) Create(evidence interface{}) (string, error) {
	return ctx.CreateMock.ID, ctx.CreateMock.Err
}

func (ctx *Evidence) LastID() (string, error) {
	return ctx.LastIDMock.ID, ctx.LastIDMock.Err
}

func (ctx *Evidence) After(afterID string) ([]interface{}, error) {
	return ctx.AfterMock.Evds, ctx.AfterMock.Err
}
