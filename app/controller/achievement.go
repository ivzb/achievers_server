package controller

import (
	"github.com/ivzb/achievers_server/app/model"
	"github.com/ivzb/achievers_server/app/shared/consts"
	"github.com/ivzb/achievers_server/app/shared/env"
	"github.com/ivzb/achievers_server/app/shared/form"
	"github.com/ivzb/achievers_server/app/shared/request"
	"github.com/ivzb/achievers_server/app/shared/response"
	"github.com/ivzb/achievers_server/app/shared/validator"
)

func AchievementCreate(env *env.Env) *response.Message {
	if !request.IsMethod(env.Request, consts.POST) {
		return response.MethodNotAllowed()
	}

	ach := &model.Achievement{}
	err := form.ModelValue(env.Request, ach)

	if err != nil {
		return response.BadRequest(err.Error())
	}

	err = validator.Validate(*ach)

	if err != nil {
		return response.BadRequest(err.Error())
	}

	involvementExists, err := env.DB.Involvement().Exists(ach.InvolvementID)

	if err != nil {
		env.Log.Error(err)
		return response.InternalServerError()
	}

	if !involvementExists {
		return response.NotFound(consts.InvolvementID)
	}

	ach.UserID = env.UserID

	id, err := env.DB.Achievement().Create(ach)

	if err != nil || id == "" {
		env.Log.Error(err)
		return response.InternalServerError()
	}

	return response.Created(
		consts.Achievement,
		id)
}

func AchievementSingle(env *env.Env) *response.Message {
	if !request.IsMethod(env.Request, consts.GET) {
		return response.MethodNotAllowed()
	}

	achID, respErr := getFormString(env, consts.ID, env.DB.Achievement())

	if respErr != nil {
		return respErr
	}

	ach, err := env.DB.Achievement().Single(achID)

	if err != nil {
		env.Log.Error(err)
		return response.InternalServerError()
	}

	return response.Ok(
		consts.Achievement,
		1,
		ach)
}

func AchievementsLast(env *env.Env) *response.Message {
	if !request.IsMethod(env.Request, consts.GET) {
		return response.MethodNotAllowed()
	}

	id, err := env.DB.Achievement().LastID()

	if err != nil {
		env.Log.Error(err)
		return response.InternalServerError()
	}

	achs, err := env.DB.Achievement().After(id)

	if err != nil {
		env.Log.Error(err)
		return response.InternalServerError()
	}

	return response.Ok(
		consts.Achievements,
		len(achs),
		achs)
}

func AchievementsAfter(env *env.Env) *response.Message {
	if !request.IsMethod(env.Request, consts.GET) {
		return response.MethodNotAllowed()
	}

	id, respErr := getFormString(env, consts.ID, env.DB.Achievement())

	if respErr != nil {
		return respErr
	}

	achs, err := env.DB.Achievement().After(id)

	if err != nil {
		env.Log.Error(err)
		return response.InternalServerError()
	}

	return response.Ok(
		consts.Achievements,
		len(achs),
		achs)
}

func AchievementsByQuestIDLast(env *env.Env) *response.Message {
	if !request.IsMethod(env.Request, consts.GET) {
		return response.MethodNotAllowed()
	}

	qstID, respErr := getFormString(env, consts.QuestID, env.DB.Quest())

	if respErr != nil {
		return respErr
	}

	id, err := env.DB.Achievement().LastIDByQuestID(qstID)

	if err != nil {
		env.Log.Error(err)
		return response.InternalServerError()
	}

	achs, err := env.DB.Achievement().AfterByQuestID(qstID, id)

	if err != nil {
		env.Log.Error(err)
		return response.InternalServerError()
	}

	return response.Ok(
		consts.Achievements,
		len(achs),
		achs)
}

func AchievementsByQuestIDAfter(env *env.Env) *response.Message {
	if !request.IsMethod(env.Request, consts.GET) {
		return response.MethodNotAllowed()
	}

	qstID, respErr := getFormString(env, consts.QuestID, env.DB.Quest())

	if respErr != nil {
		return respErr
	}

	id, respErr := getFormString(env, consts.ID, env.DB.Achievement())

	if respErr != nil {
		return respErr
	}

	achs, err := env.DB.Achievement().AfterByQuestID(qstID, id)

	if err != nil {
		env.Log.Error(err)
		return response.InternalServerError()
	}

	return response.Ok(
		consts.Achievements,
		len(achs),
		achs)
}
