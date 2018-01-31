package db

import "testing"

var mt = (&DB{}).MultimediaType().(*MultimediaType)

func TestMultimediaTypeTable(t *testing.T) {
	expected := "multimedia_type"
	actual := mt.table

	testAssert(t, "table", expected, actual)
}

func TestMultimediaTypeSelect(t *testing.T) {
	expected := ""
	actual := mt.selectArgs

	testAssert(t, "query", expected, actual)
}

func TestMultimediaTypeInsert(t *testing.T) {
	expected := ""
	actual := mt.insertArgs

	testAssert(t, "query", expected, actual)
}
