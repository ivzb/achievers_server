package mock

import (
	"github.com/ivzb/achievers_server/app/model"
)

type Evidence struct {
	ExistsMock EvidenceExists
	SingleMock EvidenceSingle
	CreateMock EvidenceCreate
	AllMock    EvidencesAll
}

type EvidenceExists struct {
	Bool bool
	Err  error
}

type EvidenceSingle struct {
	Evd *model.Evidence
	Err error
}

type EvidencesAll struct {
	Evds []*model.Evidence
	Err  error
}

type EvidenceCreate struct {
	ID  string
	Err error
}

func (ctx *Evidence) Exists(id string) (bool, error) {
	return ctx.ExistsMock.Bool, ctx.ExistsMock.Err
}

func (ctx *Evidence) Single(id string) (*model.Evidence, error) {
	return ctx.SingleMock.Evd, ctx.SingleMock.Err
}

func (ctx *Evidence) Create(evidence *model.Evidence) (string, error) {
	return ctx.CreateMock.ID, ctx.CreateMock.Err
}

func (ctx *Evidence) All(page int) ([]*model.Evidence, error) {
	return ctx.AllMock.Evds, ctx.AllMock.Err
}
