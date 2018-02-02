package db

import (
	"testing"

	"github.com/ivzb/achievers_server/app/db/mock/generate"
	"github.com/ivzb/achievers_server/app/model"
	"github.com/ivzb/achievers_server/app/shared/consts"
)

var qst = (&DB{}).Quest().(*Quest)

func TestQuestTable(t *testing.T) {
	expected := "quest"
	actual := qst.table

	testAssert(t, "table", expected, actual)
}

func TestQuestSelectQuery(t *testing.T) {
	expected := "id, title, picture_url, involvement_id, quest_type_id, user_id, created_at, updated_at, deleted_at"
	actual := qst.selectArgs

	testAssert(t, "query", expected, actual)
}

func TestQuestInsertQuery(t *testing.T) {
	expected := "title, picture_url, involvement_id, quest_type_id, user_id"
	actual := qst.insertArgs

	testAssert(t, "query", expected, actual)
}

func TestQuestSelect(t *testing.T) {
	expected := *generate.Quest().(*model.Quest)

	qst := &Quest{
		newContext(nil, consts.Quest, new(model.Quest)),
	}

	testSingle(t, qst, expected)
}

func TestQuestInsert(t *testing.T) {
	mdl := generate.Quest().(*model.Quest)
	expected := mdl.ID

	qst := &Quest{
		newContext(nil, consts.Quest, new(model.Quest)),
	}

	testCreate(t, qst, mdl, expected)
}
