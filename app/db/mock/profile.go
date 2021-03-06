package mock

type Profile struct {
	ExistsMock         ProfileExists
	SingleMock         ProfileSingle
	SingleByUserIDMock ProfileSingleByUserID
	CreateMock         ProfileCreate
}

type ProfileExists struct {
	Bool bool
	Err  error
}

type ProfileSingle struct {
	Prfl interface{}
	Err  error
}

type ProfileSingleByUserID struct {
	Prfl interface{}
	Err  error
}

type ProfileCreate struct {
	ID  string
	Err error
}

func (ctx *Profile) Exists(id interface{}) (bool, error) {
	return ctx.ExistsMock.Bool, ctx.ExistsMock.Err
}

func (ctx *Profile) Single(id string) (interface{}, error) {
	return ctx.SingleMock.Prfl, ctx.SingleMock.Err
}

func (ctx *Profile) SingleByUserID(userID string) (interface{}, error) {
	return ctx.SingleByUserIDMock.Prfl, ctx.SingleByUserIDMock.Err
}

func (ctx *Profile) Create(profile interface{}) (string, error) {
	return ctx.CreateMock.ID, ctx.CreateMock.Err
}
