package db

import (
	"github.com/ivzb/achievers_server/app/model"
)

type Rewarder interface {
	Exists(id string) (bool, error)
	Single(id string) (*model.Reward, error)
	Create(reward *model.Reward) (string, error)

	LastID() (string, error)
	After(afterID string) ([]*model.Reward, error)
}

type Reward struct {
	db         *DB
	table      string
	selectArgs string
}

func (db *DB) Reward() Rewarder {
	return &Reward{
		db:         db,
		table:      "reward",
		selectArgs: "id, title, description, picture_url, reward_type_id, user_id, created_at, updated_at, deleted_at",
	}
}

func (*Reward) scan(row sqlScanner) (*model.Reward, error) {
	rwd := new(model.Reward)

	err := row.Scan(
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

	return rwd, nil
}

func (ctx *Reward) Exists(id string) (bool, error) {
	return exists(ctx.db, "reward", "id", id)
}

func (ctx *Reward) Single(id string) (*model.Reward, error) {
	row := ctx.db.QueryRow("SELECT "+ctx.selectArgs+
		" FROM "+ctx.table+
		" WHERE id = $1 "+
		" LIMIT 1", id)

	return ctx.scan(row)
}

func (ctx *Reward) Create(reward *model.Reward) (string, error) {
	return create(ctx.db, `INSERT INTO reward(title, description, picture_url, reward_type_id, user_id)
		VALUES($1, $2, $3, $4, $5)`,
		reward.Title,
		reward.Description,
		reward.PictureURL,
		reward.RewardTypeID,
		reward.UserID)
}

func (ctx *Reward) LastID() (string, error) {
	return lastID(ctx.db, ctx.table)
}

func (ctx *Reward) After(afterID string) ([]*model.Reward, error) {
	rows, err := ctx.db.Query("SELECT "+ctx.selectArgs+
		" FROM reward "+
		" WHERE created_at <= "+
		"  (SELECT created_at"+
		"   FROM reward"+
		"   WHERE id = $1)"+
		" ORDER BY created_at DESC"+
		" LIMIT $2", afterID, limit)

	if err != nil {
		return nil, err
	}

	defer rows.Close()

	rwds := make([]*model.Reward, 0)

	for rows.Next() {
		rwd, err := ctx.scan(rows)

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
