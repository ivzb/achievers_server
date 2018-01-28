package db

import "github.com/ivzb/achievers_server/app/model"

type Quester interface {
	Exists(id string) (bool, error)
	Single(id string) (*model.Quest, error)
	Create(quest *model.Quest) (string, error)

	LastID() (string, error)
	After(afterID string) ([]*model.Quest, error)
}

type Quest struct {
	*Context
}

func (db *DB) Quest() Quester {
	return &Quest{
		&Context{
			db:         db,
			table:      "quest",
			selectArgs: "id, title, picture_url, involvement_id, user_id, created_at, updated_at, deleted_at",
		},
	}
}

func (*Quest) scan(row sqlScanner) (*model.Quest, error) {
	rwd := new(model.Quest)

	err := row.Scan(
		&rwd.ID,
		&rwd.Title,
		&rwd.PictureURL,
		&rwd.InvolvementID,
		&rwd.UserID,
		&rwd.CreatedAt,
		&rwd.UpdatedAt,
		&rwd.DeletedAt)

	if err != nil {
		return nil, err
	}

	return rwd, nil
}

func (ctx *Quest) Exists(id string) (bool, error) {
	return exists(ctx.Context, "id", id)
}

func (ctx *Quest) Single(id string) (*model.Quest, error) {
	row := ctx.db.QueryRow("SELECT "+ctx.selectArgs+
		" FROM "+ctx.table+
		" WHERE id = $1 "+
		" LIMIT 1", id)

	return ctx.scan(row)
}

func (ctx *Quest) Create(quest *model.Quest) (string, error) {
	return create(ctx.db, `INSERT INTO quest (title, picture_url, involvement_id, quest_type_id, user_id)
		VALUES($1, $2, $3, $4, $5)`,
		quest.Title,
		quest.PictureURL,
		quest.InvolvementID,
		quest.QuestTypeID,
		quest.UserID)
}

func (ctx *Quest) LastID() (string, error) {
	return lastID(ctx.db, ctx.table)
}

func (ctx *Quest) After(afterID string) ([]*model.Quest, error) {
	selectArgs := "id, title, picture_url, involvement_id, quest_type_id, user_id, created_at, updated_at, deleted_at "
	rows, err := ctx.db.Query("SELECT "+selectArgs+
		"FROM quest "+
		"WHERE created_at <= "+
		"  (SELECT created_at "+
		"   FROM quest "+
		"   WHERE id = $1) "+
		"ORDER BY created_at DESC "+
		"LIMIT $2", afterID, limit)

	if err != nil {
		return nil, err
	}

	defer rows.Close()

	qsts := make([]*model.Quest, 0)

	for rows.Next() {
		qst := new(model.Quest)

		err := rows.Scan(
			&qst.ID,
			&qst.Title,
			&qst.PictureURL,
			&qst.InvolvementID,
			&qst.QuestTypeID,
			&qst.UserID,
			&qst.CreatedAt,
			&qst.UpdatedAt,
			&qst.DeletedAt)

		if err != nil {
			return nil, err
		}

		qsts = append(qsts, qst)
	}

	if err = rows.Err(); err != nil {
		return nil, err
	}

	return qsts, nil
}
