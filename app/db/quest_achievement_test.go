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

func TestQuestAchievementAchievementInsert(t *testing.T) {
	mdl := generate.QuestAchievement().(*model.QuestAchievement)
	expected := mdl.ID

	qstAch := &QuestAchievement{
		newContext(nil, consts.QuestAchievement, new(model.QuestAchievement)),
	}

	testCreate(t, qstAch, mdl, expected)
}
