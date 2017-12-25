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
	pageSize = 9
)

// DBSource contains all available DAO functions
type DBSource interface {
	Exists(table string, column string, value string) (bool, error)

	AchievementSingle(id string) (*Achievement, error)
	AchievementsAll(page int) ([]*Achievement, error)
	UserCreate(string, string, string, string) (string, error)
	UserAuth(string, string) (string, error)
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

// Exists checks whether row in specified table exists by column and value
func (db *DB) Exists(table string, column string, value string) (bool, error) {
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

// UUID generates UUID for use as an ID
func (db *DB) UUID() (string, error) {
	b := make([]byte, 16)
	_, err := rand.Read(b)
	if err != nil {
		return "", err
	}

	return fmt.Sprintf("%x-%x-%x-%x-%x", b[0:4], b[4:6], b[6:8], b[8:10], b[10:]), nil
}
