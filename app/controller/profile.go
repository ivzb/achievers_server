package controller

import (
	"github.com/ivzb/achievers_server/app/model"
	"github.com/ivzb/achievers_server/app/shared/consts"
	"github.com/ivzb/achievers_server/app/shared/form"
	"github.com/ivzb/achievers_server/app/shared/response"
)

func ProfileMe(env *model.Env) *response.Message {
	if env.Request.Method != consts.GET {
		return response.MethodNotAllowed()
	}

	prfl, err := env.DB.ProfileByUserID(env.UserID)

	if err != nil {
		env.Log.Error(err)
		return response.InternalServerError()
	}

	return response.Ok(
		consts.Profile,
		1,
		prfl)
}

func ProfileSingle(env *model.Env) *response.Message {
	if env.Request.Method != consts.GET {
		return response.MethodNotAllowed()
	}

	prflID, err := form.StringValue(env.Request, consts.ID)

	if err != nil {
		return response.BadRequest(err.Error())
	}

	exists, err := env.DB.ProfileExists(prflID)

	if err != nil {
		env.Log.Error(err)
		return response.InternalServerError()
	}

	if !exists {
		return response.NotFound(consts.Profile)
	}

	prfl, err := env.DB.ProfileSingle(prflID)

	if err != nil {
		env.Log.Error(err)
		return response.InternalServerError()
	}

	return response.Ok(
		consts.Profile,
		1,
		prfl)
}
