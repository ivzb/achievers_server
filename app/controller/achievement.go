package controller

import (
	"fmt"

	"github.com/ivzb/achievers_server/app/model"
	"github.com/ivzb/achievers_server/app/shared/response"
)

func AchievementsIndex(env *model.Env) *response.Message {
	if !env.Request.IsMethod(GET) {
		return response.MethodNotAllowed()
	}

	pg, err := env.Request.Form.IntValue(page)

	if err != nil {
		return response.BadRequest(err.Error())
	}

	if pg < 0 {
		return response.BadRequest(fmt.Sprintf(formatInvalid, page))
	}

	achs, err := env.DB.AchievementsAll(pg)

	if err != nil {
		env.Log.Error(err)
		return response.InternalServerError()
	}

	if len(achs) == 0 {
		return response.NotFound(page)
	}

	return response.Ok(
		achievements,
		len(achs),
		achs)
}

func AchievementsByQuestID(env *model.Env) *response.Message {
	if !env.Request.IsMethod(GET) {
		return response.MethodNotAllowed()
	}

	pg, err := env.Request.Form.IntValue(page)

	if err != nil {
		return response.BadRequest(err.Error())
	}

	if pg < 0 {
		return response.BadRequest(fmt.Sprintf(formatInvalid, page))
	}

	qstID, err := env.Request.Form.StringValue(id)

	if err != nil {
		return response.BadRequest(err.Error())
	}

	qstExists, err := env.DB.QuestExists(qstID)

	if err != nil {
		env.Log.Error(err)
		return response.InternalServerError()
	} else if !qstExists {
		return response.NotFound(id)
	}

	achs, err := env.DB.AchievementsByQuestID(qstID, pg)

	if err != nil {
		env.Log.Error(err)
		return response.InternalServerError()
	} else if len(achs) == 0 {
		return response.NotFound(page)
	}

	return response.Ok(
		achievements,
		len(achs),
		achs)
}

func AchievementSingle(env *model.Env) *response.Message {
	if !env.Request.IsMethod(GET) {
		return response.MethodNotAllowed()
	}

	achID, err := env.Request.Form.StringValue(id)

	if err != nil {
		return response.BadRequest(err.Error())
	}

	achExists, err := env.DB.AchievementExists(achID)

	if err != nil {
		env.Log.Error(err)
		return response.InternalServerError()
	}

	if !achExists {
		return response.NotFound(achievement)
	}

	ach, err := env.DB.AchievementSingle(achID)

	if err != nil {
		env.Log.Error(err)
		return response.InternalServerError()
	}

	return response.Ok(
		achievement,
		1,
		ach)
}

func AchievementCreate(env *model.Env) *response.Message {
	if !env.Request.IsMethod(POST) {
		return response.MethodNotAllowed()
	}

	ach := &model.Achievement{}
	err := env.Request.Form.Map(ach)

	if err != nil {
		return response.BadRequest(err.Error())
	}

	if ach.Title == "" {
		return response.BadRequest(fmt.Sprintf(formatMissing, title))
	}

	if ach.Description == "" {
		return response.BadRequest(fmt.Sprintf(formatMissing, description))
	}

	if ach.PictureURL == "" {
		return response.BadRequest(fmt.Sprintf(formatMissing, pictureURL))
	}

	if ach.InvolvementID == "" {
		return response.BadRequest(fmt.Sprintf(formatMissing, involvementID))
	}

	involvementExists, err := env.DB.InvolvementExists(ach.InvolvementID)

	if err != nil {
		env.Log.Error(err)
		return response.InternalServerError()
	} else if !involvementExists {
		return response.NotFound(involvementID)
	}

	ach.AuthorID = env.Request.UserID

	id, err := env.DB.AchievementCreate(ach)

	if err != nil || id == "" {
		return response.InternalServerError()
	}

	return response.Created(
		achievement,
		id)
}
