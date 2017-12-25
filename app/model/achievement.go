package model

import (
	"time"
)

type Achievement struct {
	Id            string    `json:"id"`
	Title         string    `json:"title"`
	Description   string    `json:"description"`
	PictureUrl    string    `json:"picture_url"`
	InvolvementId string    `json:"involvement_id"`
	AuthorId      string    `json:"author_id"`
	CreatedAt     time.Time `json:"created_at"`
	UpdatedAt     time.Time `json:"updated_at"`
	DeletedAt     time.Time `json:"deleted_at"`
}

func (db *DB) AchievementSingle(id string) (*Achievement, error) {
	ach := new(Achievement)

	ach.Id = id

	row := db.QueryRow("SELECT `title`, `description`, `picture_url`, `involvement_id`, `author_id`, `created_at`, `updated_at`, `deleted_at` "+
		"FROM achievement "+
		"WHERE id = ? "+
		"LIMIT 1", id)

	err := row.Scan(
		&ach.Title,
		&ach.Description,
		&ach.PictureUrl,
		&ach.InvolvementId,
		&ach.AuthorId,
		&ach.CreatedAt,
		&ach.UpdatedAt,
		&ach.DeletedAt)

	if err != nil {
		return nil, err
	}

	return ach, nil
}

func (db *DB) AchievementsAll(page int) ([]*Achievement, error) {
	start := pageSize * page
	end := start + pageSize
	rows, err := db.Query("SELECT `id`, `title`, `description`, `picture_url`, `involvement_id`, `author_id`, `created_at`, `updated_at`, `deleted_at` "+
		"FROM achievement "+
		"ORDER BY id DESC "+
		"LIMIT ?, ?", start, end)

	if err != nil {
		return nil, err
	}

	defer rows.Close()

	achs := make([]*Achievement, 0)

	for rows.Next() {
		ach := new(Achievement)
		err := rows.Scan(
			&ach.Id,
			&ach.Title,
			&ach.Description,
			&ach.PictureUrl,
			&ach.InvolvementId,
			&ach.AuthorId,
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