package db

import "testing"

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
