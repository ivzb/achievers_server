package controller

import (
	"fmt"

	"github.com/ivzb/achievers_server/app/db"
	"github.com/ivzb/achievers_server/app/shared/consts"
	"github.com/ivzb/achievers_server/app/shared/env"
	"github.com/ivzb/achievers_server/app/shared/form"
	"github.com/ivzb/achievers_server/app/shared/response"
)

func getFormString(env *env.Env, key string, exister db.Exister) (string, *response.Message) {
	value, err := form.StringValue(env.Request, key)

	if err != nil {
		return "", response.BadRequest(fmt.Sprintf(consts.FormatMissing, key))
	}

	exists, err := exister.Exists(value)

	if err != nil {
		env.Log.Error(err)
		return "", response.InternalServerError()
	}

	if !exists {
		return "", response.NotFound(key)
	}

	return value, nil
}
