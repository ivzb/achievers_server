package db

import (
	"testing"

	"github.com/ivzb/achievers_server/app/db/mock/generate"
	"github.com/ivzb/achievers_server/app/model"
	"github.com/ivzb/achievers_server/app/shared/consts"
)

var qstAch = (&DB{}).QuestAchievement().(*QuestAchievement)

func TestQuestAchievementTable(t *testing.T) {
	expected := "quest_achievement"
	actual := qstAch.table

	testAssert(t, "table", expected, actual)
}

func TestQuestAchievementSelectQuery(t *testing.T) {
	expected := ""
	actual := qstAch.selectArgs

	testAssert(t, "query", expected, actual)
}

func TestQuestAchievementAchievementInsertQuery(t *testing.T) {
	expected := "quest_id, achievement_id, user_id"
	actual := qstAch.insertArgs

	testAssert(t, "query", expected, actual)
}

func TestQuestAchievementExistsQuery(t *testing.T) {
	expected := "quest_id, achievement_id"
	actual := qstAch.existsArgs

	testAssert(t, "query", expected, actual)
}

func TestQuestAchievementInsert(t *testing.T) {
	mdl := generate.QuestAchievement().(*model.QuestAchievement)
	expected := mdl.ID

	qstAch := &QuestAchievement{
		newContext(nil, consts.QuestAchievement, new(model.QuestAchievement)),
	}

	testCreate(t, qstAch, mdl, expected)
}

func TestQuestAchievementExists(t *testing.T) {
	qa := generate.QuestAchievement().(*model.QuestAchievement)
	achID := qa.AchievementID
	qstID := qa.QuestID
	expected := true

	qstAch := &QuestAchievement{
		newContext(nil, consts.QuestAchievement, new(model.QuestAchievement)),
	}

	testExistsMultiple(t, qstAch, expected, achID, qstID)
}

func TestQuestAchievementExists_NotEnoughArguments(t *testing.T) {
	testQuestAchievementExistsArguments(t, "mock_id")
}

func TestQuestAchievementExists_TooManyArguments(t *testing.T) {
	testQuestAchievementExistsArguments(t, "mock_id", "mock_id", "mock_id")
}

func testQuestAchievementExistsArguments(t *testing.T, args ...interface{}) {
	expected := false

	qstAch := &QuestAchievement{
		newContext(nil, consts.QuestAchievement, new(model.QuestAchievement)),
	}

	actual, err := qstAch.Exists(args)

	if err == nil {
		t.Error("scan should have returned error")
	}

	testAssert(t, "exists", expected, actual)
}
