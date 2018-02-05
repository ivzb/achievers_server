package db

import (
	"testing"

	"github.com/ivzb/achievers_server/app/db/mock/generate"
	"github.com/ivzb/achievers_server/app/model"
	"github.com/ivzb/achievers_server/app/shared/consts"
)

var mt = (&DB{}).MultimediaType().(*MultimediaType)

func TestMultimediaTypeTable(t *testing.T) {
	expected := "multimedia_type"
	actual := mt.table

	testAssert(t, "table", expected, actual)
}

func TestMultimediaTypeSelectQuery(t *testing.T) {
	expected := ""
	actual := mt.selectArgs

	testAssert(t, "query", expected, actual)
}

func TestMultimediaTypeInsertQuery(t *testing.T) {
	expected := ""
	actual := mt.insertArgs

	testAssert(t, "query", expected, actual)
}

func TestMultimediaTypeExistsQuery(t *testing.T) {
	expected := "id"
	actual := mt.existsArgs

	testAssert(t, "query", expected, actual)
}

func TestMultimediaTypeExists(t *testing.T) {
	mdl := generate.MultimediaType().(*model.MultimediaType)
	id := mdl.ID
	expected := true

	mt := &MultimediaType{
		newContext(nil, consts.MultimediaType, new(model.MultimediaType)),
	}

	testExists(t, mt, id, expected)
}
