package controller

import (
	"fmt"

	"github.com/ivzb/achievers_server/app/model"
	"github.com/ivzb/achievers_server/app/shared/consts"
	"github.com/ivzb/achievers_server/app/shared/form"
	"github.com/ivzb/achievers_server/app/shared/request"
	"github.com/ivzb/achievers_server/app/shared/response"
)

func AchievementsLast(env *model.Env) *response.Message {
	if !request.IsMethod(env.Request, consts.GET) {
		return response.MethodNotAllowed()
	}

	id, err := env.DB.AchievementsLastID()

	if err != nil {
		env.Log.Error(err)
		return response.InternalServerError()
	}

	achs, err := env.DB.AchievementsAfter(id)

	if err != nil {
		env.Log.Error(err)
		return response.InternalServerError()
	}

	return response.Ok(
		consts.Achievements,
		len(achs),
		achs)
}

func AchievementsAfter(env *model.Env) *response.Message {
	if !request.IsMethod(env.Request, consts.GET) {
		return response.MethodNotAllowed()
	}

	id, respErr := getFormString(env, consts.AfterID)

	if respErr != nil {
		return respErr
	}

	//id, err := form.StringValue(env.Request, consts.AfterID)

	//if err != nil {
	//return response.BadRequest(fmt.Sprintf(consts.FormatMissing, consts.AfterID))
	//}

	//exists, err := env.DB.AchievementExists(id)

	//if err != nil {
	//env.Log.Error(err)
	//return response.InternalServerError()
	//}

	//if !exists {
	//return response.NotFound(consts.ID)
	//}

	achs, err := env.DB.AchievementsAfter(id)

	if err != nil {
		env.Log.Error(err)
		return response.InternalServerError()
	}

	return response.Ok(
		consts.Achievements,
		len(achs),
		achs)
}

func AchievementsByQuestIDLast(env *model.Env) *response.Message {
	if !request.IsMethod(env.Request, consts.GET) {
		return response.MethodNotAllowed()
	}

	qstID, err := form.StringValue(env.Request, consts.QuestID)

	if err != nil {
		return response.BadRequest(fmt.Sprintf(consts.FormatMissing, consts.QuestID))
	}

	qstExists, err := env.DB.QuestExists(qstID)

	if err != nil {
		env.Log.Error(err)
		return response.InternalServerError()
	}

	if !qstExists {
		return response.NotFound(consts.QuestID)
	}

	afterID, err := env.DB.AchievementsByQuestIDLastID(qstID)

	if err != nil {
		env.Log.Error(err)
		return response.InternalServerError()
	}

	achs, err := env.DB.AchievementsByQuestIDAfter(qstID, afterID)

	if err != nil {
		env.Log.Error(err)
		return response.InternalServerError()
	}

	return response.Ok(
		consts.Achievements,
		len(achs),
		achs)
}

func AchievementsByQuestIDAfter(env *model.Env) *response.Message {
	if !request.IsMethod(env.Request, consts.GET) {
		return response.MethodNotAllowed()
	}

	qstID, err := form.StringValue(env.Request, consts.QuestID)

	if err != nil {
		return response.BadRequest(fmt.Sprintf(consts.FormatMissing, consts.QuestID))
	}

	qstExists, err := env.DB.QuestExists(qstID)

	if err != nil {
		env.Log.Error(err)
		return response.InternalServerError()
	}

	if !qstExists {
		return response.NotFound(consts.QuestID)
	}

	afterID, err := form.StringValue(env.Request, consts.AfterID)

	if err != nil {
		return response.BadRequest(fmt.Sprintf(consts.FormatMissing, consts.AfterID))
	}

	exists, err := env.DB.AchievementExists(afterID)

	if err != nil {
		env.Log.Error(err)
		return response.InternalServerError()
	}

	if !exists {
		return response.NotFound(consts.AfterID)
	}

	achs, err := env.DB.AchievementsByQuestIDAfter(qstID, afterID)

	if err != nil {
		env.Log.Error(err)
		return response.InternalServerError()
	}

	return response.Ok(
		consts.Achievements,
		len(achs),
		achs)
}

func AchievementSingle(env *model.Env) *response.Message {
	if !request.IsMethod(env.Request, consts.GET) {
		return response.MethodNotAllowed()
	}

	achID, err := form.StringValue(env.Request, consts.ID)

	if err != nil {
		return response.BadRequest(fmt.Sprintf(consts.FormatMissing, consts.ID))
	}

	achExists, err := env.DB.AchievementExists(achID)

	if err != nil {
		env.Log.Error(err)
		return response.InternalServerError()
	}

	if !achExists {
		return response.NotFound(consts.Achievement)
	}

	ach, err := env.DB.AchievementSingle(achID)

	if err != nil {
		env.Log.Error(err)
		return response.InternalServerError()
	}

	return response.Ok(
		consts.Achievement,
		1,
		ach)
}

func AchievementCreate(env *model.Env) *response.Message {
	if !request.IsMethod(env.Request, consts.POST) {
		return response.MethodNotAllowed()
	}

	ach := &model.Achievement{}
	err := form.ModelValue(env.Request, ach)

	if err != nil {
		return response.BadRequest(err.Error())
	}

	if ach.Title == "" {
		return response.BadRequest(fmt.Sprintf(consts.FormatMissing, consts.Title))
	}

	if ach.Description == "" {
		return response.BadRequest(fmt.Sprintf(consts.FormatMissing, consts.Description))
	}

	if ach.PictureURL == "" {
		return response.BadRequest(fmt.Sprintf(consts.FormatMissing, consts.PictureURL))
	}

	if ach.InvolvementID == "" {
		return response.BadRequest(fmt.Sprintf(consts.FormatMissing, consts.InvolvementID))
	}

	involvementExists, err := env.DB.InvolvementExists(ach.InvolvementID)

	if err != nil {
		env.Log.Error(err)
		return response.InternalServerError()
	} else if !involvementExists {
		return response.NotFound(consts.InvolvementID)
	}

	ach.UserID = env.UserID

	id, err := env.DB.AchievementCreate(ach)

	if err != nil || id == "" {
		return response.InternalServerError()
	}

	return response.Created(
		consts.Achievement,
		id)
}
