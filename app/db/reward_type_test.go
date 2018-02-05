package db

import (
	"testing"

	"github.com/ivzb/achievers_server/app/db/mock/generate"
	"github.com/ivzb/achievers_server/app/model"
	"github.com/ivzb/achievers_server/app/shared/consts"
)

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

func TestRewardTypeExistsQuery(t *testing.T) {
	expected := "id"
	actual := rt.existsArgs

	testAssert(t, "query", expected, actual)
}

func TestRewardTypeExists(t *testing.T) {
	mdl := generate.RewardType().(*model.RewardType)
	id := mdl.ID
	expected := true

	rt := &RewardType{
		newContext(nil, consts.RewardType, new(model.RewardType)),
	}

	testExists(t, rt, id, expected)
}
