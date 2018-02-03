package db

import (
	"testing"

	"github.com/ivzb/achievers_server/app/db/mock/generate"
	"github.com/ivzb/achievers_server/app/model"
	"github.com/ivzb/achievers_server/app/shared/consts"
)

var rwd = (&DB{}).Reward().(*Reward)

func TestRewardTable(t *testing.T) {
	expected := "reward"
	actual := rwd.table

	testAssert(t, "table", expected, actual)
}

func TestRewardSelectQuery(t *testing.T) {
	expected := "id, title, description, picture_url, reward_type_id, user_id, created_at, updated_at, deleted_at"
	actual := rwd.selectArgs

	testAssert(t, "query", expected, actual)
}

func TestRewardInsertQuery(t *testing.T) {
	expected := "title, description, picture_url, reward_type_id, user_id"
	actual := rwd.insertArgs

	testAssert(t, "query", expected, actual)
}

func TestRewardSelect(t *testing.T) {
	expected := *generate.Reward().(*model.Reward)

	rwd := &Reward{
		newContext(nil, consts.Reward, new(model.Reward)),
	}

	testSingle(t, rwd, expected)
}

func TestRewardInsert(t *testing.T) {
	mdl := generate.Reward().(*model.Reward)
	expected := mdl.ID

	rwd := &Reward{
		newContext(nil, consts.Reward, new(model.Reward)),
	}

	testCreate(t, rwd, mdl, expected)
}

//func TestRewardExists(t *testing.T) {
//mdl := generate.Reward().(*model.Reward)
//id := mdl.ID
//expected := true

//rwd := &Reward{
//newContext(nil, consts.Reward, new(model.Reward)),
//}

//testExists(t, rwd, id, expected)
//}

//func testExists(t *testing.T, exister Exister, id string, expected bool) {
//db, mock, err := sqlmock.New()

//if err != nil {
//t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
//}

//ctx := reflect.ValueOf(exister).Elem().FieldByName("Context").Interface().(*Context)
//ctx.db = &DB{db, 9}

//defer db.Close()

////cols := strings.Split(ctx.insertArgs, ", ")
//rows := sqlmock.NewRows([]string{"id"}).AddRow(1)

//fields := make([]driver.Value, 0)
//fields = append(fields, "id")

////columns := strings.Count(ctx.insertArgs, ",")
////placeholders := concatPlaceholders(columns)
////query := fmt.Sprintf("SELECT COUNT(id) FROM %s WHERE %s LIMIT 1", ctx.table, whereClause(columns))
////mock.ExpectQuery("^" + regexp.QuoteMeta("SELECT COUNT(id) FROM "+ctx.table+" WHERE id = $1 LIMIT 1") + "$").
//mock.ExpectQuery(".+").
//WithArgs(id).
//WillReturnRows(rows)

//actual, err := exister.Exists(id)

//if err != nil {
//t.Errorf("error was not expected while updating stats: %s", err)
//}

//// we make sure that all expectations were met
//if err := mock.ExpectationsWereMet(); err != nil {
//t.Errorf("there were unfulfilled expectations: %s", err)
//}

//if expected != actual {
//t.Errorf("unexpected result:\ngot %v\nwant %v", actual, expected)
//}
//}
