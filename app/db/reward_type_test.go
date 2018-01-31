package db

import "testing"

var rt = (&DB{}).RewardType().(*RewardType)

func TestRewardTypeTable(t *testing.T) {
	expected := "reward_type"
	actual := rt.table

	testAssert(t, "table", expected, actual)
}

func TestRewardTypeSelect(t *testing.T) {
	expected := ""
	actual := rt.selectArgs

	testAssert(t, "query", expected, actual)
}

func TestRewardTypeInsert(t *testing.T) {
	expected := ""
	actual := rt.insertArgs

	testAssert(t, "query", expected, actual)
}
