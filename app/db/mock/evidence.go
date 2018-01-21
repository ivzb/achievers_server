package mock

import (
	"github.com/ivzb/achievers_server/app/db"
	"github.com/ivzb/achievers_server/app/model"
)

type Evidence struct {
	db *DB

	ExistsMock EvidenceExists
	SingleMock EvidenceSingle
	CreateMock EvidenceCreate
	AllMock    EvidencesAll
}

func (db *DB) Evidence() db.Evidencer {
	return &Evidence{db: db}
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

//func Evidences(size int) []*model.Evidence {
//evds := make([]*model.Evidence, size)

//for i := 0; i < size; i++ {
//evds[i] = Evidence()
//}

//return evds
//}

//func Evidence() *model.Evidence {
//evd := &model.Evidence{
//"fb7691eb-ea1d-b20f-edee-9cadcf23181f",
//"desc",
//"http://preview-url.jpg",
//"http://url.jpg",
//3,
//"65hjl4aa-719c-ca7c-fb66-80ab235c8e39",
//"4e69c9ba-719c-ca7c-fb66-80ab235c8e39",
//time.Date(2017, 12, 9, 15, 4, 23, 0, time.UTC),
//time.Date(2017, 12, 9, 15, 4, 23, 0, time.UTC),
//time.Date(0, 1, 1, 0, 0, 0, 0, time.UTC),
//}

//return evd
//}

//func (mock *DB) EvidenceExists(string) (bool, error) {
//return mock.EvidenceExistsMock.Bool, mock.EvidenceExistsMock.Err
//}

//func (mock *DB) EvidenceSingle(id string) (*model.Evidence, error) {
//return mock.EvidenceSingleMock.Evd, mock.EvidenceSingleMock.Err
//}

//func (mock *DB) EvidencesAll(page int) ([]*model.Evidence, error) {
//return mock.EvidencesAllMock.Evds, mock.EvidencesAllMock.Err
//}

//func (mock *DB) EvidenceCreate(evidence *model.Evidence) (string, error) {
//return mock.EvidenceCreateMock.ID, mock.EvidenceCreateMock.Err
//}
