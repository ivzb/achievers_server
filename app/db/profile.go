package db

import (
	"github.com/ivzb/achievers_server/app/model"
)

type Profiler interface {
	Exists(id string) (bool, error)
	Single(id string) (*model.Profile, error)
	SingleByUserID(userID string) (*model.Profile, error)
	Create(profile *model.Profile, userID string) (string, error)
}

type Profile struct {
	db *DB
}

func (db *DB) Profile() Profiler {
	return &Profile{db}
}

func (ctx *Profile) Exists(id string) (bool, error) {
	return exists(ctx.db, "profile", "id", id)
}

func (ctx *Profile) Single(id string) (*model.Profile, error) {
	prfl := new(model.Profile)

	prfl.ID = id

	row := ctx.db.QueryRow("SELECT name, created_at, updated_at, deleted_at "+
		"FROM profile "+
		"WHERE id = $1 "+
		"LIMIT 1", id)

	err := row.Scan(
		&prfl.Name,
		&prfl.CreatedAt,
		&prfl.UpdatedAt,
		&prfl.DeletedAt)

	if err != nil {
		return nil, err
	}

	return prfl, nil
}

func (ctx *Profile) SingleByUserID(userID string) (*model.Profile, error) {
	prfl := new(model.Profile)

	row := ctx.db.QueryRow("SELECT id, name, created_at, updated_at, deleted_at "+
		"FROM profile "+
		"WHERE user_id = $1 "+
		"LIMIT 1", userID)

	err := row.Scan(
		&prfl.ID,
		&prfl.Name,
		&prfl.CreatedAt,
		&prfl.UpdatedAt,
		&prfl.DeletedAt)

	if err != nil {
		return nil, err
	}

	return prfl, nil
}

func (ctx *Profile) Create(profile *model.Profile, userID string) (string, error) {
	return create(ctx.db, `INSERT INTO profile (name, user_id)
		VALUES($1, $2)`,
		profile.Name,
		userID)
}
