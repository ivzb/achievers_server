package db

import (
	"github.com/ivzb/achievers_server/app/model"
)

type Rewarder interface {
	Exists(id string) (bool, error)
	Single(id string) (interface{}, error)
	Create(reward *model.Reward) (string, error)

	LastID() (string, error)
	After(id string) ([]interface{}, error)
}

type Reward struct {
	*Context
}

func (db *DB) Reward() Rewarder {
	return &Reward{
		newContext(db, "reward", &model.Reward{}),
	}
}

func (*Reward) scan(row sqlScanner) (interface{}, error) {
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
	return ctx.exists("id", id)
}

func (ctx *Reward) Single(id string) (interface{}, error) {
	return ctx.single(id, ctx.scan)
}

func (ctx *Reward) Create(reward *model.Reward) (string, error) {
	return ctx.create(
		reward.Title,
		reward.Description,
		reward.PictureURL,
		reward.RewardTypeID,
		reward.UserID)
}

func (ctx *Reward) LastID() (string, error) {
	return ctx.lastID()
}

func (ctx *Reward) After(id string) ([]interface{}, error) {
	return ctx.after(id, ctx.scan)
}
