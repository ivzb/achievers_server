package model

import (
	"time"
)

type Reward struct {
	ID string `json:"id"`

	Title       string `json:"title"`
	Description string `json:"description"`
	PictureURL  string `json:"picture_url"`

	RewardTypeID uint8  `json:"reward_type_id"`
	UserID       string `json:"user_id"`

	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	DeletedAt time.Time `json:"deleted_at"`
}

func (db *DB) RewardExists(id string) (bool, error) {
	return exists(db, "reward", "id", id)
}

func (db *DB) RewardSingle(id string) (*Reward, error) {
	rwd := new(Reward)

	rwd.ID = id

	row := db.QueryRow("SELECT `title`, `description`, `picture_url`, `reward_type_id`, `user_id`, `created_at`, `updated_at`, `deleted_at` "+
		"FROM reward "+
		"WHERE id = ? "+
		"LIMIT 1", id)

	err := row.Scan(
		&rwd.Title,
		&rwd.Description,
		&rwd.PictureURL,
		&rwd.RewardTypeID,
		&rwd.UserID,
		&rwd.CreatedAt,
		&rwd.UpdatedAt,
		&rwd.DeletedAt)

	if err != nil {
		return nil, err
	}

	return rwd, nil
}

func (db *DB) RewardsAll(page int) ([]*Reward, error) {
	offset := limit * page

	rows, err := db.Query("SELECT `id`, `title`, `description`, `picture_url`, `reward_type_id`, `user_id`, `created_at`, `updated_at`, `deleted_at` "+
		"FROM reward "+
		"ORDER BY `created_at` DESC "+
		"LIMIT ? OFFSET ?", limit, offset)

	if err != nil {
		return nil, err
	}

	defer rows.Close()

	rwds := make([]*Reward, 0)

	for rows.Next() {
		rwd := new(Reward)
		err := rows.Scan(
			&rwd.ID,
			&rwd.Title,
			&rwd.Description,
			&rwd.PictureURL,
			&rwd.RewardTypeID,
			&rwd.UserID,
			&rwd.CreatedAt,
			&rwd.UpdatedAt,
			&rwd.DeletedAt)

		if err != nil {
			return nil, err
		}

		rwds = append(rwds, rwd)
	}

	if err = rows.Err(); err != nil {
		return nil, err
	}

	return rwds, nil
}

func (db *DB) RewardCreate(reward *Reward) (string, error) {
	return create(db, `INSERT INTO reward(id, title, description, picture_url, reward_type_id, user_id)
        VALUES(?, ?, ?, ?, ?, ?)`,
		reward.Title,
		reward.Description,
		reward.PictureURL,
		reward.RewardTypeID,
		reward.UserID)
}
