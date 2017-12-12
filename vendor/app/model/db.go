package model

import (
	"app/shared/database"
	"database/sql"
	"errors"
	_ "github.com/go-sql-driver/mysql"
)

type DBSource interface {
	Exist(table string, column string, value string) (bool, error)

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

func (db *DB) Exist(table string, column string, value string) (bool, error) {
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