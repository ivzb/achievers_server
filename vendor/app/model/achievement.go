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

func (db *DB) AllAchievements() ([]*Achievement, error) {
	rows, err := db.Query("SELECT * FROM achievement")
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
