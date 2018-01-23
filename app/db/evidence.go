package db

import "github.com/ivzb/achievers_server/app/model"

type Evidencer interface {
	Exists(id string) (bool, error)
	Single(id string) (*model.Evidence, error)
	Create(evidence *model.Evidence) (string, error)

	All(page int) ([]*model.Evidence, error)
}

type Evidence struct {
	db *DB
}

func (db *DB) Evidence() Evidencer {
	return &Evidence{db}
}

func (ctx *Evidence) Exists(id string) (bool, error) {
	return exists(ctx.db, "evidence", "id", id)
}

func (ctx *Evidence) Single(id string) (*model.Evidence, error) {
	evd := new(model.Evidence)

	evd.ID = id

	row := ctx.db.QueryRow("SELECT title, picture_url, url, multimedia_type_id, achievement_id, user_id, created_at, updated_at, deleted_at "+
		"FROM evidence "+
		"WHERE id = $1 "+
		"LIMIT 1", id)

	err := row.Scan(
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

// Create saves evidence object to db
func (ctx *Evidence) Create(evidence *model.Evidence) (string, error) {
	return create(ctx.db, `INSERT INTO evidence (id, title, picture_url, url, multimedia_type_id, achievement_id, user_id)
		VALUES($1, $2, $3, $4, $5, $6, $7)`,
		evidence.Title,
		evidence.PictureURL,
		evidence.URL,
		evidence.MultimediaTypeID,
		evidence.AchievementID,
		evidence.UserID)
}

func (ctx *Evidence) All(page int) ([]*model.Evidence, error) {
	offset := limit * page

	rows, err := ctx.db.Query("SELECT id, title, picture_url, url, multimedia_type_id, achievement_id, user_id, created_at, updated_at, deleted_at "+
		"FROM evidence "+
		"ORDER BY created_at DESC "+
		"LIMIT $1 OFFSET $2", limit, offset)

	if err != nil {
		return nil, err
	}

	defer rows.Close()

	evds := make([]*model.Evidence, 0)

	for rows.Next() {
		evd := new(model.Evidence)
		err := rows.Scan(
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

		evds = append(evds, evd)
	}

	if err = rows.Err(); err != nil {
		return nil, err
	}

	return evds, nil
}
