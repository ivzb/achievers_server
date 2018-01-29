package db

import (
	"github.com/ivzb/achievers_server/app/model"
)

type Evidencer interface {
	Exists(id string) (bool, error)
	Single(id string) (*model.Evidence, error)
	Create(evidence *model.Evidence) (string, error)

	LastID() (string, error)
	After(afterID string) ([]*model.Evidence, error)
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

func (*Evidence) scan(row sqlScanner) (*model.Evidence, error) {
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
	return exists(ctx.Context, "id", id)
}

func (ctx *Evidence) Single(id string) (*model.Evidence, error) {
	row := single(ctx.Context, id)

	return ctx.scan(row)
}

// Create saves evidence object to db
func (ctx *Evidence) Create(evidence *model.Evidence) (string, error) {
	return create(ctx.Context,
		evidence.Title,
		evidence.PictureURL,
		evidence.URL,
		evidence.MultimediaTypeID,
		evidence.AchievementID,
		evidence.UserID)
}

func (ctx *Evidence) LastID() (string, error) {
	return lastID(ctx.db, ctx.table)
}

func (ctx *Evidence) After(afterID string) ([]*model.Evidence, error) {
	rows, err := after(ctx.Context, afterID)

	if err != nil {
		return nil, err
	}

	defer rows.Close()

	evds := make([]*model.Evidence, 0)

	for rows.Next() {
		evd, err := ctx.scan(rows)

		if err != nil {
			return nil, err
		}

		evds = append(evds, evd)
	}

	if err = rows.Err(); err != nil {
		return nil, err
	}

	return evds, nil
}
