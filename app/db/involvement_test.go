package db

import (
	"testing"

	"github.com/ivzb/achievers_server/app/db/mock/generate"
	"github.com/ivzb/achievers_server/app/model"
	"github.com/ivzb/achievers_server/app/shared/consts"
)

var inv = (&DB{}).Involvement().(*Involvement)

func TestInvolvementTable(t *testing.T) {
	expected := "involvement"
	actual := inv.table

	testAssert(t, "table", expected, actual)
}

func TestInvolvementSelectQuery(t *testing.T) {
	expected := ""
	actual := inv.selectArgs

	testAssert(t, "query", expected, actual)
}

func TestInvolvementInsertQuery(t *testing.T) {
	expected := ""
	actual := inv.insertArgs

	testAssert(t, "query", expected, actual)
}

func TestInvovlementExistsQuery(t *testing.T) {
	expected := "id"
	actual := inv.existsArgs

	testAssert(t, "query", expected, actual)
}

func TestInvolvementExists(t *testing.T) {
	mdl := generate.Involvement().(*model.Involvement)
	id := mdl.ID
	expected := true

	inv := &Involvement{
		newContext(nil, consts.Involvement, new(model.Involvement)),
	}

	testExists(t, inv, id, expected)
}
