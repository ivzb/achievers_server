package db

import "testing"

var inv = (&DB{}).Involvement().(*Involvement)

func TestInvolvementTable(t *testing.T) {
	expected := "involvement"
	actual := inv.table

	testAssert(t, "table", expected, actual)
}

func TestInvolvementSelect(t *testing.T) {
	expected := ""
	actual := inv.selectArgs

	testAssert(t, "query", expected, actual)
}

func TestInvolvementInsert(t *testing.T) {
	expected := ""
	actual := inv.insertArgs

	testAssert(t, "query", expected, actual)
}
