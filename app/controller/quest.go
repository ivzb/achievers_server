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

func QuestsLast(env *app.Env) *response.Message {
	if !request.IsMethod(env.Request, consts.GET) {
		return response.MethodNotAllowed()
	}

	id, err := env.DB.Quest().LastID()

	if err != nil {
		env.Log.Error(err)
		return response.InternalServerError()
	}

	qsts, err := env.DB.Quest().After(id)

	if err != nil {
		env.Log.Error(err)
		return response.InternalServerError()
	}

	return response.Ok(
		consts.Quests,
		len(qsts),
		qsts)
}

func QuestsAfter(env *app.Env) *response.Message {
	if !request.IsMethod(env.Request, consts.GET) {
		return response.MethodNotAllowed()
	}

	id, respErr := controller.GetFormString(env, consts.ID, env.DB.Quest())

	if respErr != nil {
		return respErr
	}

	qsts, err := env.DB.Quest().After(id)

	if err != nil {
		env.Log.Error(err)
		return response.InternalServerError()
	}

	return response.Ok(
		consts.Quests,
		len(qsts),
		qsts)
}

func QuestSingle(env *app.Env) *response.Message {
	if env.Request.Method != consts.GET {
		return response.MethodNotAllowed()
	}

	id, respErr := controller.GetFormString(env, consts.ID, env.DB.Quest())

	if respErr != nil {
		return respErr
	}

	qst, err := env.DB.Quest().Single(id)

	if err != nil {
		env.Log.Error(err)
		return response.InternalServerError()
	}

	return response.Ok(
		consts.Quest,
		1,
		qst)
}

func QuestCreate(env *app.Env) *response.Message {
	if env.Request.Method != "POST" {
		return response.MethodNotAllowed()
	}

	qst := &model.Quest{}
	err := form.ModelValue(env.Request, qst)

	if err != nil {
		return response.BadRequest(err.Error())
	}

	err = validator.Validate(*qst)

	if err != nil {
		return response.BadRequest(err.Error())
	}

	involvementExists, err := env.DB.Involvement().Exists(qst.InvolvementID)

	if err != nil {
		env.Log.Error(err)
		return response.InternalServerError()
	}

	if !involvementExists {
		return response.NotFound(consts.InvolvementID)
	}

	questTypeExists, err := env.DB.QuestType().Exists(qst.QuestTypeID)

	if err != nil {
		env.Log.Error(err)
		return response.InternalServerError()
	}

	if !questTypeExists {
		return response.NotFound(consts.QuestTypeID)
	}

	qst.UserID = env.UserID

	id, err := env.DB.Quest().Create(qst)

	if err != nil || id == "" {
		env.Log.Error(err)
		return response.InternalServerError()
	}

	return response.Created(
		consts.Quest,
		id)
}
