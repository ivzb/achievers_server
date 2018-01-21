package controller

import (
	"github.com/ivzb/achievers_server/app/shared/consts"
	"github.com/ivzb/achievers_server/app/shared/env"
	"github.com/ivzb/achievers_server/app/shared/form"
	"github.com/ivzb/achievers_server/app/shared/response"
)

func ProfileMe(env *env.Env) *response.Message {
	if env.Request.Method != consts.GET {
		return response.MethodNotAllowed()
	}

	prfl, err := env.DB.Profile().SingleByUserID(env.UserID)

	if err != nil {
		env.Log.Error(err)
		return response.InternalServerError()
	}

	return response.Ok(
		consts.Profile,
		1,
		prfl)
}

func ProfileSingle(env *env.Env) *response.Message {
	if env.Request.Method != consts.GET {
		return response.MethodNotAllowed()
	}

	prflID, err := form.StringValue(env.Request, consts.ID)

	if err != nil {
		return response.BadRequest(err.Error())
	}

	exists, err := env.DB.Profile().Exists(prflID)

	if err != nil {
		env.Log.Error(err)
		return response.InternalServerError()
	}

	if !exists {
		return response.NotFound(consts.Profile)
	}

	prfl, err := env.DB.Profile().Single(prflID)

	if err != nil {
		env.Log.Error(err)
		return response.InternalServerError()
	}

	return response.Ok(
		consts.Profile,
		1,
		prfl)
}
