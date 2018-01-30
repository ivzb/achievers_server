package db

import (
	"github.com/ivzb/achievers_server/app/model"
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
		&Context{
			db:         db,
			table:      "evidence",
			selectArgs: "id, title, picture_url, url, multimedia_type_id, achievement_id, user_id, created_at, updated_at, deleted_at",
		},
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

	if err != nil {
		return nil, err
	}

	return evd, nil
}

func (ctx *Evidence) Exists(id string) (bool, error) {
	return ctx.exists("id", id)
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
