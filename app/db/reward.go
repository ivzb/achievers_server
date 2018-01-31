package db

import (
	"log"
	"reflect"

	"github.com/ivzb/achievers_server/app/model"
	"github.com/ivzb/achievers_server/app/shared/consts"
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
		newContext(db, consts.Reward, new(model.Reward)),
	}
}

func (*Reward) scan(row sqlScanner) (interface{}, error) {
	rwd := new(model.Reward)

	// get the struct type
	modelValue := reflect.ValueOf(rwd).Elem()
	modelType := modelValue.Type()
	tag := "select"

	query := make([]interface{}, 0)

	// enumerate model fields
	for i := 0; i < modelType.NumField(); i++ {
		field := modelType.Field(i)

		key := field.Tag.Get(tag)

		if len(key) > 0 {
			log.Println(field)
			query = append(query, modelValue.Field(i).Addr().Interface())
		}
	}

	err := row.Scan(query...)

	return rwd, err
}

func (ctx *Reward) Exists(id string) (bool, error) {
	return ctx.exists(consts.ID, id)
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
