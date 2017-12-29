package model

import (
	"crypto/rand"
	"database/sql"
	"errors"
	"fmt"

	"github.com/ivzb/achievers_server/app/shared/database"

	// MySQL DB driver
	_ "github.com/go-sql-driver/mysql"
)

const (
	limit = 9
)

// DBSource contains all available DAO functions
type DBSource interface {
	UserExists(id string) (bool, error)
	UserEmailExists(email string) (bool, error)
	UserCreate(user *User) (string, error)
	UserAuth(string, string) (string, error)

	AchievementExists(id string) (bool, error)
	AchievementSingle(id string) (*Achievement, error)
	AchievementsAll(page int) ([]*Achievement, error)
	AchievementCreate(achievement *Achievement) (string, error)

	EvidenceSingle(id string) (*Evidence, error)
	EvidenceExists(id string) (bool, error)
	EvidenceCreate(evidence *Evidence) (string, error)

	InvolvementExists(id string) (bool, error)

	MultimediaTypeExists(id uint8) (bool, error)
}

// DB struct holds the connection to DB
type DB struct {
	*sql.DB
}

// NewDB creates connection to the database
func NewDB(d database.Info) (*DB, error) {
	switch d.Type {
	case database.TypeMySQL:
		db, err := sql.Open("mysql", database.DSN(d.MySQL))
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

// UUID generates UUID for use as an ID
func (db *DB) UUID() (string, error) {
	b := make([]byte, 16)
	_, err := rand.Read(b)
	if err != nil {
		return "", err
	}

	return fmt.Sprintf("%x-%x-%x-%x-%x", b[0:4], b[4:6], b[6:8], b[8:10], b[10:]), nil
}

// exists checks whether row in specified table exists by column and value
func exists(db *DB, table string, column string, value string) (bool, error) {
	stmt, err := db.Prepare("SELECT COUNT(id) FROM " + table + " WHERE " + column + " = ? LIMIT 1")

	if err != nil {
		return false, err
	}

	var count int
	err = stmt.QueryRow(value).Scan(&count)

	if err != nil {
		return false, err
	}

	return count != 0, nil
}

// create executes passed query and args
func create(db *DB, query string, args ...interface{}) (string, error) {
	id, err := db.UUID()

	if err != nil {
		return "", err
	}

	args = append([]interface{}{id}, args...)

	result, err := db.Exec(query, args...)

	if err != nil {
		return "", err
	}

	if _, err = result.RowsAffected(); err != nil {
		return "", err
	}

	return args[0].(string), nil
}
