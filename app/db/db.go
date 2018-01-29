package db

import (
	"database/sql"
	"errors"
	"fmt"
	"log"
	"strconv"
	"strings"

	"github.com/ivzb/achievers_server/app/shared/database"
	"github.com/ivzb/achievers_server/app/shared/uuid"

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

type sqlScanner interface {
	Scan(dest ...interface{}) error
}

// DB struct holds the connection to DB
type DB struct {
	*sql.DB
}

type Context struct {
	db         *DB
	table      string
	selectArgs string
	insertArgs string
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
func exists(model *Context, column string, value string) (bool, error) {
	query := fmt.Sprintf("SELECT COUNT(id) FROM \"%s\" WHERE %s = $1  LIMIT 1", model.table, column)
	stmt, err := model.db.Prepare(query)

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

func concatPlaceholders(columns int) string {
	placeholders := make([]string, 0, columns)

	for i := 1; i <= columns+1; i++ {
		placeholders = append(placeholders, "$"+strconv.Itoa(i))
	}

	return strings.Join(placeholders, ", ")
}

func single(ctx *Context, id string) *sql.Row {
	row := ctx.db.QueryRow("SELECT "+ctx.selectArgs+
		" FROM "+ctx.table+
		" WHERE id = $1 "+
		" LIMIT 1", id)

	return row
}

func after(ctx *Context, afterID string) (*sql.Rows, error) {
	rows, err := ctx.db.Query("SELECT "+ctx.selectArgs+
		" FROM "+ctx.table+
		" WHERE created_at <= "+
		"  (SELECT created_at"+
		"   FROM "+ctx.table+
		"   WHERE id = $1)"+
		" ORDER BY created_at DESC"+
		" LIMIT $2", afterID, limit)

	return rows, err
}

// create executes passed query and args
func create(ctx *Context, args ...interface{}) (string, error) {
	columns := strings.Count(ctx.insertArgs, ",")
	placeholders := concatPlaceholders(columns)
	query := "INSERT INTO " + ctx.table + "(" + ctx.insertArgs + ") VALUES(" + placeholders + ") RETURNING id"
	log.Println(query)

	var id string
	err := ctx.db.QueryRow(query, args...).Scan(&id)

	if err != nil {
		return "", err
	}

	return id, nil
}

func lastID(ctx *Context) (string, error) {
	var id string

	row := ctx.db.QueryRow("SELECT id" +
		" FROM " + ctx.table +
		" ORDER BY created_at DESC" +
		" LIMIT 1")

	err := row.Scan(&id)

	if err == ErrNoRows {
		// return random uuid since there are no ids in db
		// it doesn't matter what the id is as long as it is a valid one
		id, err := uuid.NewUUID().Generate()

		if err != nil {
			return "", err
		}

		return id, nil
	} else if err != nil {
		return "", err
	}

	return id, nil
}
