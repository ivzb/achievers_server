package db

import (
	"github.com/ivzb/achievers_server/app/model"
	"github.com/ivzb/achievers_server/app/shared/consts"
)

type Achievementer interface {
	Exists(id interface{}) (bool, error)
	Single(id string) (interface{}, error)
	Create(achievement interface{}) (string, error)

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
		newContext(db, consts.Achievement, new(model.Achievement)),
	}
}

func (ctx *Achievement) Exists(id interface{}) (bool, error) {
	return ctx.exists(&model.Achievement{ID: id.(string)})
}

func (ctx *Achievement) Single(id string) (interface{}, error) {
	return ctx.single(id)
}

func (ctx *Achievement) Create(achievement interface{}) (string, error) {
	return ctx.create(achievement)
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
	return ctx.after(id)
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
