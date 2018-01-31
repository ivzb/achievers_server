package db

import "testing"

var qstAch = (&DB{}).QuestAchievement().(*QuestAchievement)

func TestQuestAchievementTable(t *testing.T) {
	expected := "quest_achievement"
	actual := qstAch.table

	testAssert(t, "table", expected, actual)
}

func TestQuestAchievementSelect(t *testing.T) {
	expected := ""
	actual := qstAch.selectArgs

	testAssert(t, "query", expected, actual)
}

func TestQuestAchievementInsert(t *testing.T) {
	expected := "quest_id, achievement_id, user_id"
	actual := qstAch.insertArgs

	testAssert(t, "query", expected, actual)
}
