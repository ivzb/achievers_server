package model

import (
	"app/shared/database"
	"database/sql"
	"errors"
	_ "github.com/go-sql-driver/mysql"
)

type Datasource interface {
	AchievementsAll() ([]*Achievement, error)
	UserCreate(string, string, string, string) (string, error)
	UserExist(string, string) (bool, error)
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
