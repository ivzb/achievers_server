package db

import (
	"database/sql/driver"
	"reflect"
	"strings"
	"testing"

	"github.com/ivzb/achievers_server/app/db/mock/generate"
	"github.com/ivzb/achievers_server/app/model"
	"github.com/ivzb/achievers_server/app/shared/consts"
	sqlmock "gopkg.in/DATA-DOG/go-sqlmock.v1"
)

var prfl = (&DB{}).Profile().(*Profile)

func TestProfileTable(t *testing.T) {
	expected := "profile"
	actual := prfl.table

	testAssert(t, "table", expected, actual)
}

func TestProfileSelectQuery(t *testing.T) {
	expected := "id, name, created_at, updated_at, deleted_at"
	actual := prfl.selectArgs

	testAssert(t, "query", expected, actual)
}

func TestProfileInsertQuery(t *testing.T) {
	expected := "name, user_id"
	actual := prfl.insertArgs

	testAssert(t, "query", expected, actual)
}

func TestProfileExistsQuery(t *testing.T) {
	expected := "id"
	actual := prfl.existsArgs

	testAssert(t, "query", expected, actual)
}

func TestProfileSelect(t *testing.T) {
	expected := *generate.Profile().(*model.Profile)

	prfl := &Profile{
		newContext(nil, consts.Profile, new(model.Profile)),
	}

	testSingle(t, prfl, expected)
}

func TestProfileInsert(t *testing.T) {
	mdl := generate.Profile().(*model.Profile)
	expected := mdl.ID

	prfl := &Profile{
		newContext(nil, consts.Profile, new(model.Profile)),
	}

	testCreate(t, prfl, mdl, expected)
}

func TestProfileExists(t *testing.T) {
	mdl := generate.Profile().(*model.Profile)
	id := mdl.ID
	expected := true

	prfl := &Profile{
		newContext(nil, consts.Profile, new(model.Profile)),
	}

	testExists(t, prfl, id, expected)
}

func TestProfileSingleByUserID(t *testing.T) {
	expected := *generate.Profile().(*model.Profile)

	prfl := &Profile{
		newContext(nil, consts.Profile, new(model.Profile)),
	}

	db, mock, err := sqlmock.New()

	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}

	ctx := reflect.ValueOf(prfl).Elem().FieldByName("Context").Interface().(*Context)
	ctx.db = &DB{db, 9}

	defer db.Close()

	structInstance := reflect.ValueOf(expected)
	structType := structInstance.Type()
	fields := make([]driver.Value, 0)
	tag := "select"

	// enumerate struct fields
	for i := 0; i < structType.NumField(); i++ {
		hasTag := structType.Field(i).Tag.Get(tag)

		if len(hasTag) > 0 {
			// add field for scan since it has tag we are looking for
			field := structInstance.Field(i).Interface()
			fields = append(fields, field)
		}
	}

	cols := strings.Split(ctx.selectArgs, ", ")
	rows := sqlmock.NewRows(cols).AddRow(fields...)
	mockID := "mock_id"

	mock.ExpectQuery("^SELECT (.+) FROM " + ctx.table + " WHERE user_id = \\$1 LIMIT 1$").WithArgs(mockID).WillReturnRows(rows)

	actual, err := prfl.SingleByUserID(mockID)

	if err != nil {
		t.Errorf("error was not expected while updating stats: %s", err)
	}

	// we make sure that all expectations were met
	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("there were unfulfilled expectations: %s", err)
	}

	if expected != actual {
		t.Errorf("unexpected result:\ngot %v\nwant %v", actual, expected)
	}
}
