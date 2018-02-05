package db

import (
	"testing"

	"github.com/ivzb/achievers_server/app/db/mock/generate"
	"github.com/ivzb/achievers_server/app/model"
	"github.com/ivzb/achievers_server/app/shared/consts"
)

var qt = (&DB{}).QuestType().(*QuestType)

func TestQuestTypeTable(t *testing.T) {
	expected := "quest_type"
	actual := qt.table

	testAssert(t, "table", expected, actual)
}

func TestQuestTypeSelectQuery(t *testing.T) {
	expected := ""
	actual := qt.selectArgs

	testAssert(t, "query", expected, actual)
}

func TestQuestTypeInsertQuery(t *testing.T) {
	expected := ""
	actual := qt.insertArgs

	testAssert(t, "query", expected, actual)
}

func TestQuestTypeExistsQuery(t *testing.T) {
	expected := "id"
	actual := qt.existsArgs

	testAssert(t, "query", expected, actual)
}

func TestQuestTypeExists(t *testing.T) {
	mdl := generate.QuestType().(*model.QuestType)
	id := mdl.ID
	expected := true

	qt := &QuestType{
		newContext(nil, consts.QuestType, new(model.QuestType)),
	}

	testExists(t, qt, id, expected)
}
