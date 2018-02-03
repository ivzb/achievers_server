package db

import (
	"testing"

	"github.com/ivzb/achievers_server/app/db/mock/generate"
	"github.com/ivzb/achievers_server/app/model"
	"github.com/ivzb/achievers_server/app/shared/consts"
)

var evd = (&DB{}).Evidence().(*Evidence)

func TestEvidenceTable(t *testing.T) {
	expected := "evidence"
	actual := evd.table

	testAssert(t, "table", expected, actual)
}

func TestEvidenceSelectQuery(t *testing.T) {
	expected := "id, title, picture_url, url, multimedia_type_id, achievement_id, user_id, created_at, updated_at, deleted_at"
	actual := evd.selectArgs

	testAssert(t, "query", expected, actual)
}

func TestEvidenceInsertQuery(t *testing.T) {
	expected := "title, picture_url, url, multimedia_type_id, achievement_id, user_id"
	actual := evd.insertArgs

	testAssert(t, "query", expected, actual)
}

func TestEvidenceSelect(t *testing.T) {
	expected := *generate.Evidence().(*model.Evidence)

	evd := &Evidence{
		newContext(nil, consts.Evidence, new(model.Evidence)),
	}

	testSingle(t, evd, expected)
}

func TestEvidenceInsert(t *testing.T) {
	mdl := generate.Evidence().(*model.Evidence)
	expected := mdl.ID

	evd := &Evidence{
		newContext(nil, consts.Evidence, new(model.Evidence)),
	}

	testCreate(t, evd, mdl, expected)
}

func TestEvidenceExists(t *testing.T) {
	mdl := generate.Evidence().(*model.Evidence)
	id := mdl.ID
	expected := true

	evd := &Evidence{
		newContext(nil, consts.Evidence, new(model.Evidence)),
	}

	testExists(t, evd, id, expected)
}
