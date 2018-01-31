package db

import "testing"

var prfl = (&DB{}).Profile().(*Profile)

func TestProfileTable(t *testing.T) {
	expected := "profile"
	actual := prfl.table

	testAssert(t, "table", expected, actual)
}

func TestProfileSelect(t *testing.T) {
	expected := "id, name, created_at, updated_at, deleted_at"
	actual := prfl.selectArgs

	testAssert(t, "query", expected, actual)
}

func TestProfileInsert(t *testing.T) {
	expected := "name, user_id"
	actual := prfl.insertArgs

	testAssert(t, "query", expected, actual)
}
