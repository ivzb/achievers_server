package db

import (
	"reflect"
	"strings"

	"github.com/ivzb/achievers_server/app/model"
	"github.com/ivzb/achievers_server/app/shared/consts"
)

type Rewarder interface {
	Exists(id string) (bool, error)
	Single(id string) (interface{}, error)
	Create(reward interface{}) (string, error)

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

func (ctx *Reward) Exists(id string) (bool, error) {
	return ctx.exst(&model.Reward{ID: id})
}

func (ctx *Reward) Single(id string) (interface{}, error) {
	return ctx.single(id)
}

func (ctx *Reward) Create(reward interface{}) (string, error) {
	return ctx.create(reward)
}

func (ctx *Reward) LastID() (string, error) {
	return ctx.lastID()
}

func (ctx *Reward) After(id string) ([]interface{}, error) {
	return ctx.after(id)
}

// create executes passed query and args
func (ctx *Context) exst(model interface{}) (bool, error) {
	columns := strings.Split(ctx.existsArgs, ", ")
	where := whereClause(columns)
	query := "SELECT COUNT(id) FROM " + ctx.table + " WHERE " + where + " LIMIT 1"

	// instantiate struct via its type
	structInstance := reflect.ValueOf(model).Elem()
	structType := structInstance.Type()
	fieldsForExists := make([]interface{}, 0)
	tag := "exists"

	// enumerate struct fields
	for i := 0; i < structType.NumField(); i++ {
		hasTag := structType.Field(i).Tag.Get(tag)

		if len(hasTag) > 0 {
			field := structInstance.Field(i).Addr().Interface()
			fieldsForExists = append(fieldsForExists, field)
		}
	}

	var count int
	err := ctx.db.QueryRow(query, fieldsForExists...).Scan(&count)

	if err != nil {
		return false, err
	}

	return count > 0, nil
}
