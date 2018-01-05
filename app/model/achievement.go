package model

import (
	"time"
)

type Achievement struct {
	ID string `json:"id"`

	Title       string `json:"title"`
	Description string `json:"description"`
	PictureURL  string `json:"picture_url"`

	InvolvementID string `json:"involvement_id"`
	AuthorID      string `json:"author_id"`

	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	DeletedAt time.Time `json:"deleted_at"`
}

func (db *DB) AchievementExists(id string) (bool, error) {
	return exists(db, "achievement", "id", id)
}

func (db *DB) AchievementSingle(id string) (*Achievement, error) {
	ach := new(Achievement)

	ach.ID = id

	row := db.QueryRow("SELECT `title`, `description`, `picture_url`, `involvement_id`, `author_id`, `created_at`, `updated_at`, `deleted_at` "+
		"FROM achievement "+
		"WHERE id = ? "+
		"LIMIT 1", id)

	err := row.Scan(
		&ach.Title,
		&ach.Description,
		&ach.PictureURL,
		&ach.InvolvementID,
		&ach.AuthorID,
		&ach.CreatedAt,
		&ach.UpdatedAt,
		&ach.DeletedAt)

	if err != nil {
		return nil, err
	}

	return ach, nil
}

func (db *DB) AchievementsByQuestID(questID string, page int) ([]*Achievement, error) {
	offset := limit * page

	rows, err := db.Query("SELECT `a`.`id`, `a`.`title`, `a`.`description`, `a`.`picture_url`, `a`.`involvement_id`, `a`.`author_id`, `a`.`created_at`, `a`.`updated_at`, `a`.`deleted_at` "+
		"FROM achievement AS a "+
		"INNER JOIN quest_achievement as qa "+
		"ON a.id = qa.achievement_id "+
		"WHERE qa.quest_id = ? "+
		"ORDER BY `a`.`created_at` DESC "+
		"LIMIT ? OFFSET ?", questID, limit, offset)

	if err != nil {
		return nil, err
	}

	defer rows.Close()

	achs := make([]*Achievement, 0)

	for rows.Next() {
		ach := new(Achievement)
		err := rows.Scan(
			&ach.ID,
			&ach.Title,
			&ach.Description,
			&ach.PictureURL,
			&ach.InvolvementID,
			&ach.AuthorID,
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

func (db *DB) AchievementsAll(page int) ([]*Achievement, error) {
	offset := limit * page

	rows, err := db.Query("SELECT `id`, `title`, `description`, `picture_url`, `involvement_id`, `author_id`, `created_at`, `updated_at`, `deleted_at` "+
		"FROM achievement "+
		"ORDER BY `created_at` DESC "+
		"LIMIT ? OFFSET ?", limit, offset)

	if err != nil {
		return nil, err
	}

	defer rows.Close()

	achs := make([]*Achievement, 0)

	for rows.Next() {
		ach := new(Achievement)
		err := rows.Scan(
			&ach.ID,
			&ach.Title,
			&ach.Description,
			&ach.PictureURL,
			&ach.InvolvementID,
			&ach.AuthorID,
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

func (db *DB) AchievementCreate(achievement *Achievement) (string, error) {
	return create(db, `INSERT INTO achievement (id, title, description, picture_url, involvement_id, author_id)
        VALUES(?, ?, ?, ?, ?, ?)`,
		achievement.Title,
		achievement.Description,
		achievement.PictureURL,
		achievement.InvolvementID,
		achievement.AuthorID)
}
