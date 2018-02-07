package db

import (
	"database/sql/driver"
	"reflect"
	"regexp"
	"strings"
	"testing"

	"github.com/ivzb/achievers_server/app/db/mock/generate"
	"github.com/ivzb/achievers_server/app/model"
	"github.com/ivzb/achievers_server/app/shared/consts"
	sqlmock "gopkg.in/DATA-DOG/go-sqlmock.v1"
)

var ach = (&DB{}).Achievement().(*Achievement)

func TestAchievementTable(t *testing.T) {
	expected := "achievement"
	actual := ach.table

	testAssert(t, "table", expected, actual)
}

func TestAchievementSelectQuery(t *testing.T) {
	expected := "id, title, description, picture_url, involvement_id, user_id, created_at, updated_at, deleted_at"
	actual := ach.selectArgs

	testAssert(t, "query", expected, actual)
}

func TestAchievementInsertQuery(t *testing.T) {
	expected := "title, description, picture_url, involvement_id, user_id"
	actual := ach.insertArgs

	testAssert(t, "query", expected, actual)
}

func TestAchievementExistsQuery(t *testing.T) {
	expected := "id"
	actual := ach.existsArgs

	testAssert(t, "query", expected, actual)
}

func TestAchievementSelect(t *testing.T) {
	expected := *generate.Achievement().(*model.Achievement)

	ach := &Achievement{
		newContext(nil, consts.Achievement, new(model.Achievement)),
	}

	testSingle(t, ach, expected)
}

func TestAchievementInsert(t *testing.T) {
	mdl := generate.Achievement().(*model.Achievement)
	expected := mdl.ID

	ach := &Achievement{
		newContext(nil, consts.Achievement, new(model.Achievement)),
	}

	testCreate(t, ach, mdl, expected)
}

func TestAchievementExists(t *testing.T) {
	mdl := generate.Achievement().(*model.Achievement)
	id := mdl.ID
	expected := true

	ach := &Achievement{
		newContext(nil, consts.Achievement, new(model.Achievement)),
	}

	testExists(t, ach, id, expected)
}

func TestAchievementLastID(t *testing.T) {
	mdl := generate.Achievement().(*model.Achievement)
	expected := mdl.ID

	ach := &Achievement{
		newContext(nil, consts.Achievement, new(model.Achievement)),
	}

	testLastID(t, ach, expected)
}

func TestAchievementAfter(t *testing.T) {
	expected := generate.Achievements(9)

	ach := &Achievement{
		newContext(nil, consts.Achievement, new(model.Achievement)),
	}

	testAfter(t, ach, expected)
}

func TestAchievementLastIDByQuestID(t *testing.T) {
	mdl := generate.Achievement().(*model.Achievement)
	expected := mdl.ID

	ach := &Achievement{
		newContext(nil, consts.Achievement, new(model.Achievement)),
	}

	db, mock, err := sqlmock.New()

	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}

	ctx := reflect.ValueOf(ach).Elem().FieldByName("Context").Interface().(*Context)
	ctx.db = &DB{db, 9}

	defer db.Close()

	mockID := "mock_id"
	rows := sqlmock.NewRows([]string{"id"}).AddRow(mockID)

	mock.ExpectQuery("^SELECT a.id " +
		"FROM achievement as a " +
		"INNER JOIN quest_achievement as qa " +
		"ON a.id = qa.achievement_id " +
		"WHERE qa.quest_id = \\$1 " +
		"ORDER BY a.created_at DESC " +
		"LIMIT 1$").WithArgs(expected).WillReturnRows(rows)

	actual, err := ach.LastIDByQuestID(expected)

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

func TestAchievementAfterByQuestID(t *testing.T) {
	expected := generate.Achievements(9)

	ach := &Achievement{
		newContext(nil, consts.Achievement, new(model.Achievement)),
	}

	db, mock, err := sqlmock.New()

	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}

	ctx := reflect.ValueOf(ach).Elem().FieldByName("Context").Interface().(*Context)
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
	selectArgs := prefixArgsWith(ctx.selectArgs, "a.")

	mock.ExpectQuery("^"+regexp.QuoteMeta("SELECT "+selectArgs+
		" FROM "+ctx.table+" as a"+
		" INNER JOIN quest_achievement as qa"+
		" ON a.id = qa.achievement_id"+
		" WHERE qa.quest_id = $1 AND a.created_at <="+
		"  (SELECT created_at"+
		"   FROM "+ctx.table+
		"   WHERE id = $2)"+
		" ORDER BY a.created_at DESC"+
		" LIMIT $3")+"$").WithArgs(mockID, mockID, ctx.db.pageLimit).WillReturnRows(rows)

	actual, err := ach.AfterByQuestID(mockID, mockID)

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
}
