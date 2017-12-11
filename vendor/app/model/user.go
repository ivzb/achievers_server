package model

import (
	"crypto/rand"
	"errors"
	"fmt"
	"time"
)

var (
	ErrNoResult = errors.New("no result")
	ErrExists   = errors.New("already exists")
	ErrNotExist = errors.New("does not exist")
)

type User struct {
	Id        string    `json:"id"`
	FirstName string    `json:"first_name"`
	LastName  string    `json:"last_name"`
	Email     string    `json:"email"`
	Password  string    `json:"password"`
	StatusID  uint8     `json:"status_id"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	DeletedAt time.Time `json:"deleted_at"`
}

func (db *DB) UserAuth(email string, password string) (string, error) {
	stmt, err := db.Prepare("SELECT id FROM user WHERE email = ? AND password = ? LIMIT 1")
	if err != nil {
		return "", err
	}

	var uID string
	err = stmt.QueryRow(email, password).Scan(&uID)

	return uID, err
}

func (db *DB) UserExist(column string, value string) (bool, error) {
	stmt, err := db.Prepare("SELECT COUNT(id) FROM user WHERE " + column + " = ? LIMIT 1")
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

func (db *DB) UserCreate(first_name string, last_name string, email string, password string) (string, error) {
	id, err := uuid()

	if err != nil {
		return "", err
	}

	result, err := db.Exec(`INSERT INTO user (id, first_name, last_name, email, password)
        VALUES(?, ?, ?, ?, ?)`,
		id, first_name, last_name, email, password)

	if err != nil {
		return "", err
	}

	if _, err = result.RowsAffected(); err != nil {
		return "", err
	}

	return id, nil
}

func uuid() (string, error) {
	b := make([]byte, 16)
	_, err := rand.Read(b)
	if err != nil {
		return "", err
	}

	return fmt.Sprintf("%x-%x-%x-%x-%x", b[0:4], b[4:6], b[6:8], b[8:10], b[10:]), nil
}
