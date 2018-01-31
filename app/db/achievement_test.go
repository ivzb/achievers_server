package db

import "testing"

var ach = (&DB{}).Achievement().(*Achievement)

func TestAchievementTable(t *testing.T) {
	expected := "achievement"
	actual := ach.table

	testAssert(t, "table", expected, actual)
}

func TestAchievementSelect(t *testing.T) {
	expected := "id, title, description, picture_url, involvement_id, user_id, created_at, updated_at, deleted_at"
	actual := ach.selectArgs

	testAssert(t, "query", expected, actual)
}

func TestAchievementInsert(t *testing.T) {
	expected := "title, description, picture_url, involvement_id, user_id"
	actual := ach.insertArgs

	testAssert(t, "query", expected, actual)
}
