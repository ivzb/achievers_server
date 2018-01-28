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
	*Context
}

func (db *DB) Profile() Profiler {
	return &Profile{
		&Context{
			db:         db,
			table:      "profile",
			selectArgs: "id, name, created_at, updated_at, deleted_at",
		},
	}
}

func (*Profile) scan(row sqlScanner) (*model.Profile, error) {
	prfl := new(model.Profile)

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

func (ctx *Profile) Exists(id string) (bool, error) {
	return exists(ctx.Context, "id", id)
}

func (ctx *Profile) Single(id string) (*model.Profile, error) {
	row := ctx.db.QueryRow("SELECT "+ctx.selectArgs+
		" FROM "+ctx.table+
		" WHERE id = $1 "+
		" LIMIT 1", id)

	return ctx.scan(row)
}

func (ctx *Profile) SingleByUserID(userID string) (*model.Profile, error) {
	row := ctx.db.QueryRow("SELECT "+ctx.selectArgs+
		"FROM "+ctx.table+
		"WHERE user_id = $1 "+
		"LIMIT 1", userID)

	return ctx.scan(row)
}

func (ctx *Profile) Create(profile *model.Profile, userID string) (string, error) {
	return create(ctx.db, `INSERT INTO profile (name, user_id)
		VALUES($1, $2)`,
		profile.Name,
		userID)
}
