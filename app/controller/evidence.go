package controller

import (
	"github.com/ivzb/achievers_server/app"
	"github.com/ivzb/achievers_server/app/model"
	"github.com/ivzb/achievers_server/app/shared/consts"
	"github.com/ivzb/achievers_server/app/shared/controller"
	"github.com/ivzb/achievers_server/app/shared/form"
	"github.com/ivzb/achievers_server/app/shared/request"
	"github.com/ivzb/achievers_server/app/shared/response"
	"github.com/ivzb/achievers_server/app/shared/validator"
)

func EvidencesLast(env *app.Env) *response.Message {
	if !request.IsMethod(env.Request, consts.GET) {
		return response.MethodNotAllowed()
	}

	id, err := env.DB.Evidence().LastID()

	if err != nil {
		env.Log.Error(err)
		return response.InternalServerError()
	}

	evds, err := env.DB.Evidence().After(id)

	if err != nil {
		env.Log.Error(err)
		return response.InternalServerError()
	}

	return response.Ok(
		consts.Evidences,
		len(evds),
		evds)
}

func EvidencesAfter(env *app.Env) *response.Message {
	if !request.IsMethod(env.Request, consts.GET) {
		return response.MethodNotAllowed()
	}

	id, respErr := controller.GetFormString(env, consts.ID, env.DB.Evidence())

	if respErr != nil {
		return respErr
	}

	evds, err := env.DB.Evidence().After(id)

	if err != nil {
		env.Log.Error(err)
		return response.InternalServerError()
	}

	return response.Ok(
		consts.Evidences,
		len(evds),
		evds)
}

func EvidenceSingle(env *app.Env) *response.Message {
	if env.Request.Method != consts.GET {
		return response.MethodNotAllowed()
	}

	id, respErr := controller.GetFormString(env, consts.ID, env.DB.Evidence())

	if respErr != nil {
		return respErr
	}

	evd, err := env.DB.Evidence().Single(id)

	if err != nil {
		return response.InternalServerError()
	}

	return response.Ok(
		consts.Evidence,
		1,
		evd)
}

func EvidenceCreate(env *app.Env) *response.Message {
	if env.Request.Method != "POST" {
		return response.MethodNotAllowed()
	}

	evd := &model.Evidence{}
	err := form.ModelValue(env.Request, evd)

	if err != nil {
		return response.BadRequest(err.Error())
	}

	err = validator.Validate(*evd)

	if err != nil {
		return response.BadRequest(err.Error())
	}

	multimediaTypeExist, err := env.DB.MultimediaType().Exists(evd.MultimediaTypeID)

	if err != nil {
		return response.InternalServerError()
	}

	if !multimediaTypeExist {
		return response.NotFound(consts.MultimediaTypeID)
	}

	achievementExist, err := env.DB.Achievement().Exists(evd.AchievementID)

	if err != nil {
		return response.InternalServerError()
	}

	if !achievementExist {
		return response.NotFound(consts.AchievementID)
	}

	evd.UserID = env.UserID

	id, err := env.DB.Evidence().Create(evd)

	if err != nil || id == "" {
		return response.InternalServerError()
	}

	return response.Created(
		consts.Evidence,
		id)
}
