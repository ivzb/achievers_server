package model

import (
	"time"
)

type Quest struct {
	ID string `json:"id"`

	Title      string `json:"title"`
	PictureURL string `json:"picture_url"`

	InvolvementID string `json:"involvement_id"`
	QuestTypeID   uint8  `json:"quest_type_id"`
	AuthorID      string `json:"author_id"`

	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	DeletedAt time.Time `json:"deleted_at"`
}

func (db *DB) QuestExists(id string) (bool, error) {
	return exists(db, "quest", "id", id)
}

func (db *DB) QuestSingle(id string) (*Quest, error) {
	qst := new(Quest)

	qst.ID = id

	row := db.QueryRow("SELECT `title`, `picture_url`, `involvement_id`, `quest_type_id`, `author_id`, `created_at`, `updated_at`, `deleted_at` "+
		"FROM quest "+
		"WHERE id = ? "+
		"LIMIT 1", id)

	err := row.Scan(
		&qst.Title,
		&qst.PictureURL,
		&qst.InvolvementID,
		&qst.QuestTypeID,
		&qst.AuthorID,
		&qst.CreatedAt,
		&qst.UpdatedAt,
		&qst.DeletedAt)

	if err != nil {
		return nil, err
	}

	return qst, nil
}

func (db *DB) QuestsAll(page int) ([]*Quest, error) {
	offset := limit * page

	rows, err := db.Query("SELECT `id`, `title`, `picture_url`, `involvement_id`, `quest_type_id`, `author_id`, `created_at`, `updated_at`, `deleted_at` "+
		"FROM quest "+
		"ORDER BY `created_at` DESC "+
		"LIMIT ? OFFSET ?", limit, offset)

	if err != nil {
		return nil, err
	}

	defer rows.Close()

	qsts := make([]*Quest, 0)

	for rows.Next() {
		qst := new(Quest)
		err := rows.Scan(
			&qst.ID,
			&qst.Title,
			&qst.PictureURL,
			&qst.InvolvementID,
			&qst.QuestTypeID,
			&qst.AuthorID,
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

func (db *DB) QuestCreate(quest *Quest) (string, error) {
	return create(db, `INSERT INTO quest (id, title, picture_url, involvement_id, quest_type_id, author_id)
        VALUES(?, ?, ?, ?, ?, ?)`,
		quest.Title,
		quest.PictureURL,
		quest.InvolvementID,
		quest.QuestTypeID,
		quest.AuthorID)
}
