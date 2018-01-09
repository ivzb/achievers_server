package controller

import (
	"fmt"

	"github.com/ivzb/achievers_server/app/model"
	"github.com/ivzb/achievers_server/app/shared/form"
	"github.com/ivzb/achievers_server/app/shared/request"
	"github.com/ivzb/achievers_server/app/shared/response"
)

func QuestsIndex(env *model.Env) *response.Message {
	if !request.IsMethod(env.Request, GET) {
		return response.MethodNotAllowed()
	}

	pg, err := form.IntValue(env.Request, page)

	if err != nil {
		return response.BadRequest(err.Error())
	}

	if pg < 0 {
		return response.BadRequest(fmt.Sprintf(formatInvalid, page))
	}

	qsts, err := env.DB.QuestsAll(pg)

	if err != nil {
		return response.InternalServerError()
	}

	if len(qsts) == 0 {
		return response.NotFound(page)
	}

	return response.Ok(
		quests,
		len(qsts),
		qsts)
}

func QuestSingle(env *model.Env) *response.Message {
	if env.Request.Method != "GET" {
		return response.MethodNotAllowed()
	}

	qstID, err := form.StringValue(env.Request, id)

	if err != nil {
		return response.BadRequest(err.Error())
	}

	exists, err := env.DB.QuestExists(qstID)

	if err != nil {
		return response.InternalServerError()
	}

	if !exists {
		return response.NotFound(quest)
	}

	qst, err := env.DB.QuestSingle(qstID)

	if err != nil {
		return response.InternalServerError()
	}

	return response.Ok(
		quest,
		1,
		qst)
}

func QuestCreate(env *model.Env) *response.Message {
	if env.Request.Method != "POST" {
		return response.MethodNotAllowed()
	}

	qst := &model.Quest{}
	err := form.ModelValue(env.Request, qst)

	if err != nil {
		return response.BadRequest(err.Error())
	}

	if qst.Title == "" {
		return response.BadRequest(fmt.Sprintf(formatMissing, title))
	}

	if qst.PictureURL == "" {
		return response.BadRequest(fmt.Sprintf(formatMissing, pictureURL))
	}

	if qst.InvolvementID == "" {
		return response.BadRequest(fmt.Sprintf(formatMissing, involvementID))
	}

	if qst.QuestTypeID == 0 {
		return response.BadRequest(fmt.Sprintf(formatMissing, questTypeID))
	}

	involvementExists, err := env.DB.InvolvementExists(qst.InvolvementID)

	if err != nil {
		return response.InternalServerError()
	}

	if !involvementExists {
		return response.NotFound(involvementID)
	}

	questTypeExists, err := env.DB.QuestTypeExists(qst.QuestTypeID)

	if err != nil {
		return response.InternalServerError()
	}

	if !questTypeExists {
		return response.NotFound(questTypeID)
	}

	qst.AuthorID = env.UserID

	id, err := env.DB.QuestCreate(qst)

	if err != nil || id == "" {
		return response.InternalServerError()
	}

	return response.Created(
		quest,
		id)
}
