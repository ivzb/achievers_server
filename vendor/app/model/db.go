package model

import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
)

type Datasource interface {
	AchievementsAll() ([]*Achievement, error)
	UserCreate(string, string, string, string) (string, error)
}

type DB struct {
	*sql.DB
}

func NewDB(dataSourceName string) (*DB, error) {
	db, err := sql.Open("mysql", dataSourceName)
	if err != nil {
		return nil, err
	}
	if err = db.Ping(); err != nil {
		return nil, err
	}
	return &DB{db}, nil
}
