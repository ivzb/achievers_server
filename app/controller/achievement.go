package controller

import (
	"fmt"

	"github.com/ivzb/achievers_server/app/model"
	"github.com/ivzb/achievers_server/app/shared/consts"
	"github.com/ivzb/achievers_server/app/shared/form"
	"github.com/ivzb/achievers_server/app/shared/request"
	"github.com/ivzb/achievers_server/app/shared/response"
)

func AchievementsIndex(env *model.Env) *response.Message {
	if !request.IsMethod(env.Request, consts.GET) {
		return response.MethodNotAllowed()
	}

	pg, err := form.IntValue(env.Request, consts.Page)

	if err != nil {
		return response.BadRequest(err.Error())
	}

	if pg < 0 {
		return response.BadRequest(fmt.Sprintf(consts.FormatInvalid, consts.Page))
	}

	achs, err := env.DB.AchievementsAll(pg)

	if err != nil {
		env.Log.Error(err)
		return response.InternalServerError()
	}

	if len(achs) == 0 {
		return response.NotFound(consts.Page)
	}

	return response.Ok(
		consts.Achievements,
		len(achs),
		achs)
}

func AchievementsByQuestID(env *model.Env) *response.Message {
	if !request.IsMethod(env.Request, consts.GET) {
		return response.MethodNotAllowed()
	}

	pg, err := form.IntValue(env.Request, consts.Page)

	if err != nil {
		return response.BadRequest(err.Error())
	}

	if pg < 0 {
		return response.BadRequest(fmt.Sprintf(consts.FormatInvalid, consts.Page))
	}

	qstID, err := form.StringValue(env.Request, consts.ID)

	if err != nil {
		return response.BadRequest(err.Error())
	}

	qstExists, err := env.DB.QuestExists(qstID)

	if err != nil {
		env.Log.Error(err)
		return response.InternalServerError()
	} else if !qstExists {
		return response.NotFound(consts.ID)
	}

	achs, err := env.DB.AchievementsByQuestID(qstID, pg)

	if err != nil {
		env.Log.Error(err)
		return response.InternalServerError()
	} else if len(achs) == 0 {
		return response.NotFound(consts.Page)
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
		return response.BadRequest(err.Error())
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
