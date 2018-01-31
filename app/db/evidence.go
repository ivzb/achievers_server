package db

import (
	"github.com/ivzb/achievers_server/app/model"
	"github.com/ivzb/achievers_server/app/shared/consts"
)

type Evidencer interface {
	Exists(id string) (bool, error)
	Single(id string) (interface{}, error)
	Create(evidence *model.Evidence) (string, error)

	LastID() (string, error)
	After(id string) ([]interface{}, error)
}

type Evidence struct {
	*Context
}

func (db *DB) Evidence() Evidencer {
	return &Evidence{
		newContext(db, consts.Evidence, new(model.Evidence)),
	}
}

func (*Evidence) scan(row sqlScanner) (interface{}, error) {
	evd := new(model.Evidence)

	err := row.Scan(
		&evd.ID,
		&evd.Title,
		&evd.PictureURL,
		&evd.URL,
		&evd.MultimediaTypeID,
		&evd.AchievementID,
		&evd.UserID,
		&evd.CreatedAt,
		&evd.UpdatedAt,
		&evd.DeletedAt)

	return evd, err
}

func (ctx *Evidence) Exists(id string) (bool, error) {
	return ctx.exists(consts.ID, id)
}

func (ctx *Evidence) Single(id string) (interface{}, error) {
	return ctx.single(id, ctx.scan)
}

// Create saves evidence object to db
func (ctx *Evidence) Create(evidence *model.Evidence) (string, error) {
	return ctx.create(
		evidence.Title,
		evidence.PictureURL,
		evidence.URL,
		evidence.MultimediaTypeID,
		evidence.AchievementID,
		evidence.UserID)
}

func (ctx *Evidence) LastID() (string, error) {
	return ctx.lastID()
}

func (ctx *Evidence) After(id string) ([]interface{}, error) {
	return ctx.after(id, ctx.scan)
}
