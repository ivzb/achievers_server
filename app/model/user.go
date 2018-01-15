package model

import (
	"time"

	"golang.org/x/crypto/bcrypt"
)

type User struct {
	ID string `json:"id"`

	Email    string `json:"email"`
	Password string `json:"password"`

	StatusID uint8 `json:"status_id"`

	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	DeletedAt time.Time `json:"deleted_at"`
}

type Auth struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

func (db *DB) UserExists(id string) (bool, error) {
	return exists(db, "user", "id", id)
}

func (db *DB) UserEmailExists(email string) (bool, error) {
	return exists(db, "user", "email", email)
}

func (db *DB) UserAuth(auth *Auth) (string, error) {
	stmt, err := db.Prepare("SELECT id, password FROM user WHERE email = ? LIMIT 1")

	if err != nil {
		return "", err
	}

	var uID string
	var hashedPassword []byte
	err = stmt.QueryRow(auth.Email).Scan(&uID, &hashedPassword)

	err = bcrypt.CompareHashAndPassword(hashedPassword, []byte(auth.Password))

	if err != nil {
		return "", err
	}

	return uID, err
}

func (db *DB) UserCreate(user *User) (string, error) {
	id, err := db.UUID()

	if err != nil {
		return "", err
	}

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)

	if err != nil {
		return "", err
	}

	result, err := db.Exec(`INSERT INTO user (id, email, password)
        VALUES(?, ?, ?)`,
		id, user.Email, hashedPassword)

	if err != nil {
		return "", err
	}

	if _, err = result.RowsAffected(); err != nil {
		return "", err
	}

	return id, nil
}
