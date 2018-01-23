package db

import (
	"github.com/ivzb/achievers_server/app/model"
)

type Achievementer interface {
	Exists(id string) (bool, error)
	Single(id string) (*model.Achievement, error)
	Create(achievement *model.Achievement) (string, error)

	LastID() (string, error)
	LastIDByQuestID(questID string) (string, error)

	After(afterID string) ([]*model.Achievement, error)
	AfterByQuestID(questID string, afterID string) ([]*model.Achievement, error)
}

type Achievement struct {
	db *DB
}

func (db *DB) Achievement() Achievementer {
	return &Achievement{db}
}

func (ctx *Achievement) Exists(id string) (bool, error) {
	return exists(ctx.db, "achievement", "id", id)
}

func (ctx *Achievement) Single(id string) (*model.Achievement, error) {
	ach := new(model.Achievement)

	ach.ID = id

	row := ctx.db.QueryRow("SELECT title, description, picture_url, involvement_id, user_id, created_at, updated_at, deleted_at "+
		"FROM achievement "+
		"WHERE id = $1 "+
		"LIMIT 1", id)

	err := row.Scan(
		&ach.Title,
		&ach.Description,
		&ach.PictureURL,
		&ach.InvolvementID,
		&ach.UserID,
		&ach.CreatedAt,
		&ach.UpdatedAt,
		&ach.DeletedAt)

	if err != nil {
		return nil, err
	}

	return ach, nil
}

func (ctx *Achievement) Create(achievement *model.Achievement) (string, error) {
	return create(ctx.db, `INSERT INTO achievement (title, description, picture_url, involvement_id, user_id)
        VALUES($1, $2, $3, $4, $5)`,
		achievement.Title,
		achievement.Description,
		achievement.PictureURL,
		achievement.InvolvementID,
		achievement.UserID)
}

func (ctx *Achievement) LastID() (string, error) {
	var id string

	row := ctx.db.QueryRow("SELECT id " +
		"FROM achievement " +
		"ORDER BY created_at DESC " +
		"LIMIT 1")

	err := row.Scan(&id)

	if err == ErrNoRows {
		return "", nil
	} else if err != nil {
		return "", err
	}

	return id, nil
}

func (ctx *Achievement) LastIDByQuestID(questID string) (string, error) {
	var id string

	row := ctx.db.QueryRow("SELECT a.id "+
		"FROM achievement as a "+
		"INNER JOIN quest_achievement as qa "+
		"ON a.id = qa.achievement_id "+
		"WHERE qa.quest_id = $1 "+
		"ORDER BY a.created_at DESC "+
		"LIMIT 1", questID)

	err := row.Scan(&id)

	if err == ErrNoRows {
		return "", nil
	} else if err != nil {
		return "", err
	}

	return id, nil
}

func (ctx *Achievement) After(afterID string) ([]*model.Achievement, error) {
	selectArgs := "id, title, description, picture_url, involvement_id, user_id, created_at, updated_at, deleted_at "
	rows, err := ctx.db.Query("SELECT "+selectArgs+
		"FROM achievement "+
		"WHERE created_at <= "+
		"  (SELECT created_at "+
		"   FROM achievement "+
		"   WHERE id = $1) "+
		"ORDER BY created_at DESC "+
		"LIMIT $2", afterID, limit)

	if err != nil {
		return nil, err
	}

	defer rows.Close()

	achs := make([]*model.Achievement, 0)

	for rows.Next() {
		ach := new(model.Achievement)
		err := rows.Scan(
			&ach.ID,
			&ach.Title,
			&ach.Description,
			&ach.PictureURL,
			&ach.InvolvementID,
			&ach.UserID,
			&ach.CreatedAt,
			&ach.UpdatedAt,
			&ach.DeletedAt)

		if err != nil {
			return nil, err
		}

		achs = append(achs, ach)
	}

	if err = rows.Err(); err != nil {
		return nil, err
	}

	return achs, nil
}

func (ctx *Achievement) AfterByQuestID(questID string, afterID string) ([]*model.Achievement, error) {
	rows, err := ctx.db.Query("SELECT a.id, a.title, a.description, a.picture_url, a.involvement_id, a.user_id, a.created_at, a.updated_at, a.deleted_at "+
		"FROM achievement as a "+
		"INNER JOIN quest_achievement as qa "+
		"ON a.id = qa.achievement_id "+
		"WHERE qa.quest_id = $1 AND a.created_at <= "+
		"  (SELECT created_at "+
		"   FROM achievement "+
		"   WHERE id = $2) "+
		"ORDER BY a.created_at DESC "+
		"LIMIT $3", questID, afterID, limit)

	if err != nil {
		return nil, err
	}

	defer rows.Close()

	achs := make([]*model.Achievement, 0)

	for rows.Next() {
		ach := new(model.Achievement)

		err := rows.Scan(
			&ach.ID,
			&ach.Title,
			&ach.Description,
			&ach.PictureURL,
			&ach.InvolvementID,
			&ach.UserID,
			&ach.CreatedAt,
			&ach.UpdatedAt,
			&ach.DeletedAt)

		if err != nil {
			return nil, err
		}

		achs = append(achs, ach)
	}

	if err = rows.Err(); err != nil {
		return nil, err
	}

	return achs, nil
}
