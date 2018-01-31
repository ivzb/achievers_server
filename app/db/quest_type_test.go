package db

import "testing"

var qt = (&DB{}).QuestType().(*QuestType)

func TestQuestTypeTable(t *testing.T) {
	expected := "quest_type"
	actual := qt.table

	testAssert(t, "table", expected, actual)
}

func TestQuestTypeSelect(t *testing.T) {
	expected := ""
	actual := qt.selectArgs

	testAssert(t, "query", expected, actual)
}

func TestQuestTypeInsert(t *testing.T) {
	expected := ""
	actual := qt.insertArgs

	testAssert(t, "query", expected, actual)
}
