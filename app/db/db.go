package db

import (
	"database/sql"
	"errors"
	"fmt"
	"strconv"
	"strings"

	"github.com/ivzb/achievers_server/app/shared/database"

	// Postgre DB driver
	_ "github.com/lib/pq"
)

const (
	limit = 9
)

var (
	ErrNoRows = sql.ErrNoRows
)

// DBSourcer contains all available DAO functions
type DBSourcer interface {
	User() Userer
	Profile() Profiler

	Achievement() Achievementer
	Evidence() Evidencer
	Involvement() Involvementer

	Quest() Quester
	QuestType() QuestTyper
	QuestAchievement() QuestAchievementer

	Reward() Rewarder
	RewardType() RewardTyper

	MultimediaType() MultimediaTyper
}

type Exister interface {
	Exists(id string) (bool, error)
}

// DB struct holds the connection to DB
type DB struct {
	*sql.DB
}

// NewDB creates connection to the database
func NewDB(d database.Info) (*DB, error) {
	switch d.Type {
	case database.TypePostgre:
		db, err := sql.Open("postgres", database.DSN(d.Postgre))
		if err != nil {
			return nil, err
		}
		if err = db.Ping(); err != nil {
			return nil, err
		}
		return &DB{db}, nil
	default:
		return nil, errors.New("No registered database in config")
	}
}

// exists checks whether row in specified table exists by column and value
func exists(db *DB, table string, column string, value string) (bool, error) {
	query := fmt.Sprintf("SELECT COUNT(id) FROM \"%s\" WHERE %s = $1  LIMIT 1", table, column)
	stmt, err := db.Prepare(query)

	if err != nil {
		return false, err
	}

	var count int
	err = stmt.QueryRow(value).Scan(&count)

	if err != nil {
		return false, err
	}

	return count > 0, nil
}

// existsMultiple checks whether row in specified table exists by []columns and []values
func existsMultiple(db *DB, table string, columns []string, values []string) (bool, error) {
	query := fmt.Sprintf("SELECT COUNT(id) FROM %s WHERE %s LIMIT 1", table, whereClause(columns))
	stmt, err := db.Prepare(query)

	if err != nil {
		return false, err
	}

	var count int
	err = stmt.QueryRow(scanArgs(values)...).Scan(&count)

	if err != nil {
		return false, err
	}

	return count != 0, nil
}

func scanArgs(values []string) []interface{} {
	scanArgs := make([]interface{}, len(values))

	for i := range values {
		scanArgs[i] = &values[i]
	}

	return scanArgs
}

func whereClause(columns []string) string {
	placeholders := make([]string, 0, len(columns))

	for i, column := range columns {

		placeholders = append(placeholders, column+" = $"+strconv.Itoa(i+1))
	}

	return strings.Join(placeholders, " AND ")
}

// create executes passed query and args
func create(db *DB, query string, args ...interface{}) (string, error) {
	id := ""
	query = query + " RETURNING id"
	err := db.QueryRow(query, args...).Scan(&id)

	if err != nil {
		return "", err
	}

	return id, nil
}
