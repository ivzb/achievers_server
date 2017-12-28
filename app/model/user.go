package model

import (
	"time"
)

type User struct {
	ID string `json:"id"`

	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	Email     string `json:"email"`
	Password  string `json:"password"`

	StatusID uint8 `json:"status_id"`

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

func (db *DB) UserCreate(user *User) (string, error) {
	id, err := db.UUID()

	if err != nil {
		return "", err
	}

	result, err := db.Exec(`INSERT INTO user (id, first_name, last_name, email, password)
        VALUES(?, ?, ?, ?, ?)`,
		id, user.FirstName, user.LastName, user.Email, user.Password)

	if err != nil {
		return "", err
	}

	if _, err = result.RowsAffected(); err != nil {
		return "", err
	}

	return id, nil
}
