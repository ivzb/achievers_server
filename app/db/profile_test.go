package db

import (
	"testing"

	"github.com/ivzb/achievers_server/app/db/mock/generate"
	"github.com/ivzb/achievers_server/app/model"
	"github.com/ivzb/achievers_server/app/shared/consts"
)

var prfl = (&DB{}).Profile().(*Profile)

func TestProfileTable(t *testing.T) {
	expected := "profile"
	actual := prfl.table

	testAssert(t, "table", expected, actual)
}

func TestProfileSelectQuery(t *testing.T) {
	expected := "id, name, created_at, updated_at, deleted_at"
	actual := prfl.selectArgs

	testAssert(t, "query", expected, actual)
}

func TestProfileInsertQuery(t *testing.T) {
	expected := "name, user_id"
	actual := prfl.insertArgs

	testAssert(t, "query", expected, actual)
}

func TestProfileSelect(t *testing.T) {
	expected := *generate.Profile().(*model.Profile)

	prfl := &Profile{
		newContext(nil, consts.Profile, new(model.Profile)),
	}

	testSingle(t, prfl, expected)
}
