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
	*Context
}

func (db *DB) User() Userer {
	return &User{
		&Context{
			db:    db,
			table: "user",
		},
	}
}

func (ctx *User) Exists(id string) (bool, error) {
	return exists(ctx.Context, "id", id)
}

func (ctx *User) EmailExists(email string) (bool, error) {
	return exists(ctx.Context, "email", email)
}

func (ctx *User) Auth(auth *model.Auth) (string, error) {
	stmt, err := ctx.db.Prepare("SELECT id, password FROM \"user\" WHERE email = $1 LIMIT 1")

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
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)

	if err != nil {
		return "", err
	}

	id := ""
	err = ctx.db.QueryRow(`INSERT INTO "user" (email, password)
        VALUES($1, $2)
		RETURNING id`,
		user.Email,
		hashedPassword).Scan(&id)

	if err != nil {
		return "", err
	}

	return id, nil
}
