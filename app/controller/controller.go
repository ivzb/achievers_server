package controller

import (
	"fmt"

	"github.com/ivzb/achievers_server/app/model"
	"github.com/ivzb/achievers_server/app/shared/consts"
	"github.com/ivzb/achievers_server/app/shared/form"
	"github.com/ivzb/achievers_server/app/shared/response"
)

func getFormString(env *model.Env, key string) (string, *response.Message) {
	value, err := form.StringValue(env.Request, key)

	if err != nil {
		return "", response.BadRequest(fmt.Sprintf(consts.FormatMissing, key))
	}

	exists, err := env.DB.AchievementExists(value)

	if err != nil {
		env.Log.Error(err)
		return "", response.InternalServerError()
	}

	if !exists {
		return "", response.NotFound(key)
	}

	return value, nil
}
