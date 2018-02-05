package db

import (
	"database/sql"
	"errors"
	"fmt"
	"reflect"
	"strconv"
	"strings"

	"github.com/ivzb/achievers_server/app/shared/database"
	"github.com/ivzb/achievers_server/app/shared/uuid"

	// Postgre DB driver
	_ "github.com/lib/pq"
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

type Singler interface {
	Single(id string) (interface{}, error)
}

type Creator interface {
	Create(model interface{}) (string, error)
}

type Exister interface {
	Exists(field interface{}) (bool, error)
}

type ExisterMultiple interface {
	Exists(field ...interface{}) (bool, error)
}

type sqlScanner interface {
	Scan(...interface{}) error
}

// DB struct holds the connection to DB
type DB struct {
	*sql.DB
	pageLimit int
}

type Context struct {
	db    *DB
	table string
	model interface{}

	selectArgs string
	insertArgs string
	existsArgs string
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
		return &DB{db, d.Postgre.PageLimit}, nil
	default:
		return nil, errors.New("No registered database in config")
	}
}

func newContext(db *DB, table string, model interface{}) *Context {
	return &Context{
		db:         db,
		table:      table,
		model:      model,
		selectArgs: buildQuery("select", model),
		insertArgs: buildQuery("insert", model),
		existsArgs: buildQuery("exists", model),
	}
}

// scan crawls struct fields for given tag and passes these who match to sqlScanner
// in order to populate struct fields with data from db
func scan(row sqlScanner, tag string, model interface{}) (interface{}, error) {
	// instantiate struct via its type
	structType := reflect.TypeOf(model).Elem()
	structInstance := reflect.New(structType).Elem()
	fieldsForScan := make([]interface{}, 0)

	// enumerate struct fields
	for i := 0; i < structType.NumField(); i++ {
		hasTag := structType.Field(i).Tag.Get(tag)

		if len(hasTag) > 0 {
			// add field for scan since it has tag we are looking for
			field := structInstance.Field(i).Addr().Interface()
			fieldsForScan = append(fieldsForScan, field)
		}
	}

	err := row.Scan(fieldsForScan...)

	return structInstance.Interface(), err
}

// exists checks whether row in specified table exists by column and value
func (ctx *Context) existsBy(column string, value string) (bool, error) {
	query := fmt.Sprintf("SELECT COUNT(id) FROM %s WHERE %s = $1 LIMIT 1", ctx.table, column)

	var count int
	err := ctx.db.QueryRow(query, value).Scan(&count)

	if err != nil {
		return false, err
	}

	return count > 0, nil
}

func buildQuery(tag string, model interface{}) string {
	if model == nil {
		return ""
	}

	// get the struct type
	modelValue := reflect.ValueOf(model).Elem()
	modelType := modelValue.Type()

	query := make([]string, 0)

	// enumerate model fields
	for i := 0; i < modelType.NumField(); i++ {
		field := modelType.Field(i)

		key := field.Tag.Get(tag)

		if len(key) > 0 {
			query = append(query, key)
		}
	}

	return strings.Join(query, ", ")
}

func whereClause(columns []string) string {
	placeholders := make([]string, 0, len(columns))

	for i, column := range columns {
		placeholders = append(placeholders, column+" = $"+strconv.Itoa(i+1))
	}

	return strings.Join(placeholders, " AND ")
}

func (ctx *Context) single(id string) (interface{}, error) {
	row := ctx.db.QueryRow("SELECT "+ctx.selectArgs+
		" FROM "+ctx.table+
		" WHERE id = $1 "+
		" LIMIT 1", id)

	return scan(row, "select", ctx.model)
}

// create executes passed query and args
func (ctx *Context) create(model interface{}) (string, error) {
	columns := strings.Count(ctx.insertArgs, ",")
	placeholders := concatPlaceholders(columns)
	query := "INSERT INTO " + ctx.table + "(" + ctx.insertArgs + ") VALUES(" + placeholders + ") RETURNING id"

	// instantiate struct via its type
	structInstance := reflect.ValueOf(model).Elem()
	structType := structInstance.Type()
	fieldsToInsert := make([]interface{}, 0)
	tag := "insert"

	// enumerate struct fields
	for i := 0; i < structType.NumField(); i++ {
		hasTag := structType.Field(i).Tag.Get(tag)

		if len(hasTag) > 0 {
			field := structInstance.Field(i).Addr().Interface()
			fieldsToInsert = append(fieldsToInsert, field)
		}
	}

	var id string

	err := ctx.db.QueryRow(query, fieldsToInsert...).Scan(&id)

	if err != nil {
		return "", err
	}

	return id, nil
}

func concatPlaceholders(columns int) string {
	placeholders := make([]string, 0, columns)

	for i := 1; i <= columns+1; i++ {
		placeholders = append(placeholders, "$"+strconv.Itoa(i))
	}

	return strings.Join(placeholders, ", ")
}

func (ctx *Context) lastID() (string, error) {
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

func (ctx *Context) exists(model interface{}) (bool, error) {
	columns := strings.Split(ctx.existsArgs, ", ")
	where := whereClause(columns)
	query := "SELECT COUNT(id) FROM " + ctx.table + " WHERE " + where + " LIMIT 1"

	// instantiate struct via its type
	structInstance := reflect.ValueOf(model).Elem()
	structType := structInstance.Type()
	fieldsForExists := make([]interface{}, 0)
	tag := "exists"

	// enumerate struct fields
	for i := 0; i < structType.NumField(); i++ {
		hasTag := structType.Field(i).Tag.Get(tag)

		if len(hasTag) > 0 {
			field := structInstance.Field(i).Addr().Interface()
			fieldsForExists = append(fieldsForExists, field)
		}
	}

	var count int
	err := ctx.db.QueryRow(query, fieldsForExists...).Scan(&count)

	if err != nil {
		return false, err
	}

	return count > 0, nil
}

func (ctx *Context) after(id string) ([]interface{}, error) {
	rows, err := ctx.db.Query("SELECT "+ctx.selectArgs+
		" FROM "+ctx.table+
		" WHERE created_at <= "+
		"  (SELECT created_at"+
		"   FROM "+ctx.table+
		"   WHERE id = $1)"+
		" ORDER BY created_at DESC"+
		" LIMIT $2", id, ctx.db.pageLimit)

	defer rows.Close()

	mdls := make([]interface{}, 0)

	for rows.Next() {
		mdl, err := scan(rows, "select", ctx.model)

		if err != nil {
			return nil, err
		}

		mdls = append(mdls, mdl)
	}

	if err = rows.Err(); err != nil {
		return nil, err
	}

	return mdls, nil
}
