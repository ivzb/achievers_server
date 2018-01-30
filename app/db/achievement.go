package db

import (
	"github.com/ivzb/achievers_server/app/model"
)

type Achievementer interface {
	Exists(id string) (bool, error)
	Single(id string) (interface{}, error)
	Create(achievement *model.Achievement) (string, error)

	LastID() (string, error)
	LastIDByQuestID(questID string) (string, error)

	After(id string) ([]interface{}, error)
	AfterByQuestID(questID string, afterID string) ([]interface{}, error)
}

type Achievement struct {
	*Context
}

func (db *DB) Achievement() Achievementer {
	return &Achievement{
		&Context{
			db:         db,
			table:      "achievement",
			selectArgs: "id, title, description, picture_url, involvement_id, user_id, created_at, updated_at, deleted_at",
			insertArgs: "title, description, picture_url, involvement_id, user_id",
		},
	}
}

func (*Achievement) scan(row sqlScanner) (interface{}, error) {
	ach := new(model.Achievement)

	err := row.Scan(
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

	return ach, nil
}

func (ctx *Achievement) Exists(id string) (bool, error) {
	return ctx.exists("id", id)
}

func (ctx *Achievement) Single(id string) (interface{}, error) {
	return ctx.single(id, ctx.scan)
}

func (ctx *Achievement) Create(achievement *model.Achievement) (string, error) {
	return ctx.create(
		achievement.Title,
		achievement.Description,
		achievement.PictureURL,
		achievement.InvolvementID,
		achievement.UserID)
}

func (ctx *Achievement) LastID() (string, error) {
	return ctx.lastID()
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

func (ctx *Achievement) After(id string) ([]interface{}, error) {
	return ctx.after(id, ctx.scan)
}

func (ctx *Achievement) AfterByQuestID(questID string, afterID string) ([]interface{}, error) {
	rows, err := ctx.db.Query("SELECT a.id, a.title, a.description, a.picture_url, a.involvement_id, a.user_id, a.created_at, a.updated_at, a.deleted_at "+
		"FROM achievement as a "+
		"INNER JOIN quest_achievement as qa "+
		"ON a.id = qa.achievement_id "+
		"WHERE qa.quest_id = $1 AND a.created_at <= "+
		"  (SELECT created_at "+
		"   FROM achievement "+
		"   WHERE id = $2) "+
		"ORDER BY a.created_at DESC "+
		"LIMIT $3", questID, afterID, ctx.db.pageLimit)

	if err != nil {
		return nil, err
	}

	defer rows.Close()

	achs := make([]interface{}, 0)

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
