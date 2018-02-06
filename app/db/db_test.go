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

func testExists(t *testing.T, exister Exister, id interface{}, expected bool) {
	db, mock, err := sqlmock.New()

	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}

	ctx := reflect.ValueOf(exister).Elem().FieldByName("Context").Interface().(*Context)
	ctx.db = &DB{db, 9}

	defer db.Close()

	rows := sqlmock.NewRows([]string{"count"}).AddRow(1)

	fields := make([]driver.Value, 0)
	fields = append(fields, "id")

	mock.ExpectQuery("^" + regexp.QuoteMeta("SELECT COUNT(id) FROM "+ctx.table+" WHERE id = $1 LIMIT 1") + "$").
		WithArgs(id).
		WillReturnRows(rows)

	actual, err := exister.Exists(id)

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

func testExistsMultiple(t *testing.T, exister ExisterMultiple, expected bool, args ...driver.Value) {
	db, mock, err := sqlmock.New()

	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}

	ctx := reflect.ValueOf(exister).Elem().FieldByName("Context").Interface().(*Context)
	ctx.db = &DB{db, 9}

	defer db.Close()

	rows := sqlmock.NewRows([]string{"count"}).AddRow(1)

	columns := strings.Split(ctx.existsArgs, ", ")
	placeholders := whereClause(columns)

	mock.ExpectQuery("^" + regexp.QuoteMeta("SELECT COUNT(id) FROM "+ctx.table+" WHERE "+placeholders+" LIMIT 1") + "$").
		WithArgs(args...).
		WillReturnRows(rows)

	exstArgs := make([]interface{}, len(args))

	for i := range args {
		exstArgs[i] = args[i]
	}

	actual, err := exister.Exists(exstArgs...)

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

func testLastID(t *testing.T, laster Laster, expected string) {
	db, mock, err := sqlmock.New()

	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}

	ctx := reflect.ValueOf(laster).Elem().FieldByName("Context").Interface().(*Context)
	ctx.db = &DB{db, 9}

	defer db.Close()

	mockID := "mock_id"
	rows := sqlmock.NewRows([]string{"id"}).AddRow(mockID)

	mock.ExpectQuery("^SELECT id FROM " + ctx.table + " ORDER BY created_at DESC LIMIT 1$").WillReturnRows(rows)

	actual, err := laster.LastID()

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

func testAfter(t *testing.T, afterer Afterer, expected []interface{}) {
	db, mock, err := sqlmock.New()

	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}

	ctx := reflect.ValueOf(afterer).Elem().FieldByName("Context").Interface().(*Context)
	ctx.db = &DB{db, 9}

	defer db.Close()

	cols := strings.Split(ctx.selectArgs, ", ")
	rows := sqlmock.NewRows(cols)

	for row := 0; row < len(expected); row++ {
		structInstance := reflect.ValueOf(expected[row]).Elem()
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

		rows = rows.AddRow(fields...)
	}

	mockID := "mock_id"

	mock.ExpectQuery("^"+regexp.QuoteMeta("SELECT "+ctx.selectArgs+
		" FROM "+ctx.table+
		" WHERE created_at <= "+
		"  (SELECT created_at"+
		"   FROM "+ctx.table+
		"   WHERE id = $1)"+
		" ORDER BY created_at DESC"+
		" LIMIT $2")+"$").WithArgs(mockID, ctx.db.pageLimit).WillReturnRows(rows)

	actual, err := afterer.After(mockID)

	if err != nil {
		t.Errorf("error was not expected while updating stats: %s", err)
	}

	// we make sure that all expectations were met
	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("there were unfulfilled expectations: %s", err)
	}

	if len(expected) != len(actual) {
		t.Errorf("unexpected length:\ngot %v\nwant %v", len(actual), len(expected))
	}

	if reflect.DeepEqual(expected, actual) {
		t.Errorf("unexpected result:\ngot %v\nwant %v", actual, expected)
	}
}

func testAssert(t *testing.T, param string, expected interface{}, actual interface{}) {
	if expected != actual {
		t.Errorf("model returned wrong %v: \ngot \"%v\" \nwant \"%v\"", param, actual, expected)
	}
}
