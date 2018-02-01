package db

import (
	"strings"
	"testing"

	"github.com/ivzb/achievers_server/app/db/mock/generate"
	"github.com/ivzb/achievers_server/app/model"
	"github.com/ivzb/achievers_server/app/shared/consts"
	"gopkg.in/DATA-DOG/go-sqlmock.v1"
)

const (
	rewardSelectArgs = "id, title, description, picture_url, reward_type_id, user_id, created_at, updated_at, deleted_at"
	rewardInsertArgs = "title, description, picture_url, reward_type_id, user_id"
)

var rwd = (&DB{}).Reward().(*Reward)

func TestRewardTable(t *testing.T) {
	expected := "reward"
	actual := rwd.table

	testAssert(t, "table", expected, actual)
}

func TestRewardSelectQuery(t *testing.T) {
	actual := rwd.selectArgs

	testAssert(t, "query", rewardSelectArgs, actual)
}

func TestRewardInsertQuery(t *testing.T) {
	actual := rwd.insertArgs

	testAssert(t, "query", rewardInsertArgs, actual)
}

func TestRewardInsertDynamic(t *testing.T) {
	db, mock, err := sqlmock.New()

	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}

	defer db.Close()

	expected := generate.Reward().(*model.Reward)

	rows := sqlmock.NewRows(strings.Split(rewardSelectArgs, ", ")).
		AddRow(
			expected.ID,
			expected.Title,
			expected.Description,
			expected.PictureURL,
			expected.RewardTypeID,
			expected.UserID,
			expected.CreatedAt,
			expected.UpdatedAt,
			expected.DeletedAt)

	id := "mock_id"

	mock.ExpectQuery("^SELECT (.+) FROM reward WHERE id = \\$1 LIMIT 1$").WithArgs(id).WillReturnRows(rows)

	rwd := &Reward{
		newContext(&DB{db, 9}, consts.Reward, new(model.Reward)),
	}

	actual, err := rwd.Single(id)

	if err != nil {
		t.Errorf("error was not expected while updating stats: %s", err)
	}

	// we make sure that all expectations were met
	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("there were unfulfilled expectations: %s", err)
	}

	if *expected != actual {
		t.Errorf("unexpected result:\ngot %v\nwant %v", actual, expected)
	}
}

func TestRewardInsert(t *testing.T) {
	db, mock, err := sqlmock.New()

	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}

	defer db.Close()

	expected := generate.Reward().(*model.Reward)

	rows := sqlmock.NewRows(strings.Split(rewardSelectArgs, ", ")).
		AddRow(
			expected.ID,
			expected.Title,
			expected.Description,
			expected.PictureURL,
			expected.RewardTypeID,
			expected.UserID,
			expected.CreatedAt,
			expected.UpdatedAt,
			expected.DeletedAt)

	id := "mock_id"

	mock.ExpectQuery("^SELECT (.+) FROM reward WHERE id = \\$1 LIMIT 1$").WithArgs(id).WillReturnRows(rows)

	rwd := &Reward{
		newContext(&DB{db, 9}, consts.Reward, new(model.Reward)),
	}

	actual, err := rwd.Single(id)

	if err != nil {
		t.Errorf("error was not expected while updating stats: %s", err)
	}

	// we make sure that all expectations were met
	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("there were unfulfilled expectations: %s", err)
	}

	if *expected != actual {
		t.Errorf("unexpected result:\ngot %v\nwant %v", actual, expected)
	}
}
