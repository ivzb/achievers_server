package db

import "github.com/ivzb/achievers_server/app/model"

type Rewarder interface {
	Exists(id string) (bool, error)
	Single(id string) (*model.Reward, error)
	All(page int) ([]*model.Reward, error)
	Create(reward *model.Reward) (string, error)
}

type Reward struct {
	db *DB
}

func (db *DB) Reward() Rewarder {
	return &Reward{db}
}

func (ctx *Reward) Exists(id string) (bool, error) {
	return exists(ctx.db, "reward", "id", id)
}

func (ctx *Reward) Single(id string) (*model.Reward, error) {
	rwd := new(model.Reward)

	rwd.ID = id

	row := ctx.db.QueryRow("SELECT title, description, picture_url, reward_type_id, user_id, created_at, updated_at, deleted_at "+
		"FROM reward "+
		"WHERE id = $1 "+
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

func (ctx *Reward) All(page int) ([]*model.Reward, error) {
	offset := limit * page

	rows, err := ctx.db.Query("SELECT id, title, description, picture_url, reward_type_id, user_id, created_at, updated_at, deleted_at "+
		"FROM reward "+
		"ORDER BY created_at DESC "+
		"LIMIT $1 OFFSET $2", limit, offset)

	if err != nil {
		return nil, err
	}

	defer rows.Close()

	rwds := make([]*model.Reward, 0)

	for rows.Next() {
		rwd := new(model.Reward)
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

func (ctx *Reward) Create(reward *model.Reward) (string, error) {
	return create(ctx.db, `INSERT INTO reward(id, title, description, picture_url, reward_type_id, user_id)
		VALUES($1, $2, $3, $4, $5, $6)`,
		reward.Title,
		reward.Description,
		reward.PictureURL,
		reward.RewardTypeID,
		reward.UserID)
}
