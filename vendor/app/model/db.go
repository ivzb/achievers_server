package model

import (
	"app/shared/database"
	"crypto/rand"
	"database/sql"
	"errors"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
)

type DBSource interface {
	Exists(table string, column string, value string) (bool, error)

	AchievementsAll() ([]*Achievement, error)
	UserCreate(string, string, string, string) (string, error)
	UserAuth(string, string) (string, error)
}

type DB struct {
	*sql.DB
}

// Connect to the database
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
