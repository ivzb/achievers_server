package db

import (
	"github.com/ivzb/achievers_server/app/model"
	"github.com/ivzb/achievers_server/app/shared/consts"
)

type Profiler interface {
	Exists(id string) (bool, error)
	Single(id string) (interface{}, error)
	SingleByUserID(userID string) (interface{}, error)
	Create(profile *model.Profile, userID string) (string, error)
}

type Profile struct {
	*Context
}

func (db *DB) Profile() Profiler {
	return &Profile{
		newContext(db, consts.Profile, new(model.Profile)),
	}
}

func (*Profile) scan(row sqlScanner) (interface{}, error) {
	prfl := new(model.Profile)

	err := row.Scan(
		&prfl.ID,
		&prfl.Name,
		&prfl.CreatedAt,
		&prfl.UpdatedAt,
		&prfl.DeletedAt)

	return prfl, err
}

func (ctx *Profile) Exists(id string) (bool, error) {
	return ctx.exists(consts.ID, id)
}

func (ctx *Profile) Single(id string) (interface{}, error) {
	return ctx.single(id, ctx.scan)
}

func (ctx *Profile) SingleByUserID(userID string) (interface{}, error) {
	row := ctx.db.QueryRow("SELECT "+ctx.selectArgs+
		" FROM "+ctx.table+
		" WHERE user_id = $1 "+
		" LIMIT 1", userID)

	return ctx.scan(row)
}

func (ctx *Profile) Create(profile *model.Profile, userID string) (string, error) {
	return ctx.create(profile.Name, userID)
}
