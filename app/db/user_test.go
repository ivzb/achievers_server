package db

import (
	"testing"

	"github.com/ivzb/achievers_server/app/model"
	"github.com/ivzb/achievers_server/app/shared/consts"
)

var usr = (&DB{}).User().(*User)

func TestUserTable(t *testing.T) {
	expected := "user"
	actual := usr.table

	testAssert(t, "table", expected, actual)
}

func TestUserSelect(t *testing.T) {
	expected := ""
	actual := usr.selectArgs

	testAssert(t, "query", expected, actual)
}

func TestUserInsert(t *testing.T) {
	expected := ""
	actual := usr.insertArgs

	testAssert(t, "query", expected, actual)
}

func TestUserExists(t *testing.T) {
	id := "mock_id"
	expected := true

	rwd := &User{
		newContext(nil, consts.User, new(model.User)),
	}

	testExists(t, rwd, id, expected)
}
