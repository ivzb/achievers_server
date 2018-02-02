package db

import (
	"database/sql/driver"
	"reflect"
	"regexp"
	"strings"
	"testing"

	sqlmock "gopkg.in/DATA-DOG/go-sqlmock.v1"
)

func testSingle(t *testing.T, singler Singler, expected interface{}) {
	db, mock, err := sqlmock.New()

	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}

	ctx := reflect.ValueOf(singler).Elem().FieldByName("Context").Interface().(*Context)
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

	mock.ExpectQuery("^SELECT (.+) FROM " + ctx.table + " WHERE id = \\$1 LIMIT 1$").WithArgs(mockID).WillReturnRows(rows)

	actual, err := singler.Single(mockID)

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

func testCreate(t *testing.T, creator Creator, mdl interface{}, expected string) {
	db, mock, err := sqlmock.New()

	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}

	ctx := reflect.ValueOf(creator).Elem().FieldByName("Context").Interface().(*Context)
	ctx.db = &DB{db, 9}

	defer db.Close()

	structInstance := reflect.ValueOf(mdl).Elem()
	structType := structInstance.Type()
	fields := make([]driver.Value, 0)
	tag := "insert"

	// enumerate struct fields
	for i := 0; i < structType.NumField(); i++ {
		hasTag := structType.Field(i).Tag.Get(tag)

		if len(hasTag) > 0 {
			// add field for scan since it has tag we are looking for
			field := structInstance.Field(i).Interface()
			fields = append(fields, field)
		}
	}

	//cols := strings.Split(ctx.insertArgs, ", ")
	mockID := "mock_id"
	rows := sqlmock.NewRows([]string{"id"}).AddRow(mockID)

	columns := strings.Count(ctx.insertArgs, ",")
	placeholders := concatPlaceholders(columns)
	mock.ExpectQuery("^" + regexp.QuoteMeta("INSERT INTO "+ctx.table+"("+ctx.insertArgs+") VALUES("+placeholders+") RETURNING id") + "$").
		WithArgs(fields...).
		WillReturnRows(rows)

	actual, err := creator.Create(mdl)

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

func testAssert(t *testing.T, param string, expected string, actual string) {
	if expected != actual {
		t.Errorf("model returned wrong %v: \ngot \"%v\" \nwant \"%v\"", param, actual, expected)
	}
}
