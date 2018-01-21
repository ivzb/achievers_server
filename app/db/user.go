package db

import (
	"github.com/ivzb/achievers_server/app/model"
	"golang.org/x/crypto/bcrypt"
)

type Userer interface {
	Exists(id string) (bool, error)
	EmailExists(email string) (bool, error)
	Create(user *model.User) (string, error)
	Auth(auth *model.Auth) (string, error)
}

type User struct {
	db *DB
}

func (db *DB) User() Userer {
	return &User{db}
}

func (ctx *User) Exists(id string) (bool, error) {
	return exists(ctx.db, "user", "id", id)
}

func (ctx *User) EmailExists(email string) (bool, error) {
	return exists(ctx.db, "user", "email", email)
}

func (ctx *User) Auth(auth *model.Auth) (string, error) {
	stmt, err := ctx.db.Prepare("SELECT id, password FROM user WHERE email = ? LIMIT 1")

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

func (ctx *User) Create(user *model.User) (string, error) {
	id, err := ctx.db.UUID()

	if err != nil {
		return "", err
	}

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)

	if err != nil {
		return "", err
	}

	result, err := ctx.db.Exec(`INSERT INTO user (id, email, password)
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
