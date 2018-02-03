package db

import (
	"testing"

	"github.com/ivzb/achievers_server/app/db/mock/generate"
	"github.com/ivzb/achievers_server/app/model"
	"github.com/ivzb/achievers_server/app/shared/consts"
)

var ach = (&DB{}).Achievement().(*Achievement)

func TestAchievementTable(t *testing.T) {
	expected := "achievement"
	actual := ach.table

	testAssert(t, "table", expected, actual)
}

func TestAchievementSelectQuery(t *testing.T) {
	expected := "id, title, description, picture_url, involvement_id, user_id, created_at, updated_at, deleted_at"
	actual := ach.selectArgs

	testAssert(t, "query", expected, actual)
}

func TestAchievementInsertQuery(t *testing.T) {
	expected := "title, description, picture_url, involvement_id, user_id"
	actual := ach.insertArgs

	testAssert(t, "query", expected, actual)
}

func TestAchievementSelect(t *testing.T) {
	expected := *generate.Achievement().(*model.Achievement)

	ach := &Achievement{
		newContext(nil, consts.Achievement, new(model.Achievement)),
	}

	testSingle(t, ach, expected)
}

func TestAchievementInsert(t *testing.T) {
	mdl := generate.Achievement().(*model.Achievement)
	expected := mdl.ID

	ach := &Achievement{
		newContext(nil, consts.Achievement, new(model.Achievement)),
	}

	testCreate(t, ach, mdl, expected)
}

func TestAchievementExists(t *testing.T) {
	mdl := generate.Achievement().(*model.Achievement)
	id := mdl.ID
	expected := true

	ach := &Achievement{
		newContext(nil, consts.Achievement, new(model.Achievement)),
	}

	testExists(t, ach, id, expected)
}
