package db

import "testing"

var qst = (&DB{}).Quest().(*Quest)

func TestQuestTable(t *testing.T) {
	expected := "quest"
	actual := qst.table

	testAssert(t, "table", expected, actual)
}

func TestQuestSelect(t *testing.T) {
	expected := "id, title, picture_url, involvement_id, quest_type_id, user_id, created_at, updated_at, deleted_at"
	actual := qst.selectArgs

	testAssert(t, "query", expected, actual)
}

func TestQuestInsert(t *testing.T) {
	expected := "title, picture_url, involvement_id, quest_type_id, user_id"
	actual := qst.insertArgs

	testAssert(t, "query", expected, actual)
}
