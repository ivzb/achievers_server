package db

import (
	"errors"
	"strconv"

	"github.com/ivzb/achievers_server/app/model"
	"github.com/ivzb/achievers_server/app/shared/consts"
)

type QuestAchievementer interface {
	Exists(args ...interface{}) (bool, error)
	Create(qstAch interface{}) (string, error)
}

type QuestAchievement struct {
	*Context
}

func (db *DB) QuestAchievement() QuestAchievementer {
	return &QuestAchievement{
		newContext(db, consts.QuestAchievement, new(model.QuestAchievement)),
	}
}

func (ctx *QuestAchievement) Exists(args ...interface{}) (bool, error) {
	if len(args) != 2 {
		return false, errors.New("two arguments wanted, got " + strconv.Itoa(len(args)))
	}

	qstID := args[0].(string)
	achID := args[1].(string)

	//return ctx.existsMultiple(keys, values)
	return ctx.exists(&model.QuestAchievement{
		QuestID:       qstID,
		AchievementID: achID,
	})
}

func (ctx *QuestAchievement) Create(qstAch interface{}) (string, error) {
	return ctx.create(qstAch)
}
