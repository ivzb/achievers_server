package db

import (
	"testing"
)

var rwd = (&DB{}).Reward().(*Reward)

func TestRewardTable(t *testing.T) {
	expected := "reward"
	actual := rwd.table

	testAssert(t, "table", expected, actual)
}

func TestRewardSelect(t *testing.T) {
	expected := "id, title, description, picture_url, reward_type_id, user_id, created_at, updated_at, deleted_at"
	actual := rwd.selectArgs

	testAssert(t, "query", expected, actual)
}

func TestRewardInsert(t *testing.T) {
	expected := "title, description, picture_url, reward_type_id, user_id"
	actual := rwd.insertArgs

	testAssert(t, "query", expected, actual)
}
