package db

import (
	"testing"

	"github.com/ivzb/achievers_server/app/db/mock/generate"
	"github.com/ivzb/achievers_server/app/model"
	"github.com/ivzb/achievers_server/app/shared/consts"
)

var rwd = (&DB{}).Reward().(*Reward)

func TestRewardTable(t *testing.T) {
	expected := "reward"
	actual := rwd.table

	testAssert(t, "table", expected, actual)
}

func TestRewardSelectQuery(t *testing.T) {
	expected := "id, title, description, picture_url, reward_type_id, user_id, created_at, updated_at, deleted_at"
	actual := rwd.selectArgs

	testAssert(t, "query", expected, actual)
}

func TestRewardInsertQuery(t *testing.T) {
	expected := "title, description, picture_url, reward_type_id, user_id"
	actual := rwd.insertArgs

	testAssert(t, "query", expected, actual)
}

func TestRewardExistsQuery(t *testing.T) {
	expected := "id"
	actual := rwd.existsArgs

	testAssert(t, "query", expected, actual)
}

func TestRewardSelect(t *testing.T) {
	expected := *generate.Reward().(*model.Reward)

	rwd := &Reward{
		newContext(nil, consts.Reward, new(model.Reward)),
	}

	testSingle(t, rwd, expected)
}

func TestRewardInsert(t *testing.T) {
	mdl := generate.Reward().(*model.Reward)
	expected := mdl.ID

	rwd := &Reward{
		newContext(nil, consts.Reward, new(model.Reward)),
	}

	testCreate(t, rwd, mdl, expected)
}

func TestRewardExists(t *testing.T) {
	mdl := generate.Reward().(*model.Reward)
	id := mdl.ID
	expected := true

	rwd := &Reward{
		newContext(nil, consts.Reward, new(model.Reward)),
	}

	testExists(t, rwd, id, expected)
}

func TestRewardLastID(t *testing.T) {
	mdl := generate.Reward().(*model.Reward)
	expected := mdl.ID

	rwd := &Reward{
		newContext(nil, consts.Reward, new(model.Reward)),
	}

	testLastID(t, rwd, expected)
}

func TestRewardAfter(t *testing.T) {
	expected := generate.Rewards(9)

	rwd := &Reward{
		newContext(nil, consts.Reward, new(model.Reward)),
	}

	testAfter(t, rwd, expected)
}
