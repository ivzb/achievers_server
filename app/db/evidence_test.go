package db

import (
	"testing"
)

var evd = (&DB{}).Evidence().(*Evidence)

func TestEvidenceTable(t *testing.T) {
	expected := "evidence"
	actual := evd.table

	testAssert(t, "table", expected, actual)
}

func TestEvidenceSelect(t *testing.T) {
	expected := "id, title, picture_url, url, multimedia_type_id, achievement_id, user_id, created_at, updated_at, deleted_at"
	actual := evd.selectArgs

	testAssert(t, "query", expected, actual)
}

func TestEvidenceInsert(t *testing.T) {
	expected := "title, picture_url, url, multimedia_type_id, achievement_id, user_id"
	actual := evd.insertArgs

	testAssert(t, "query", expected, actual)
}
