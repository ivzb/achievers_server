package controller

import (
	"fmt"

	"github.com/ivzb/achievers_server/app/model"
	"github.com/ivzb/achievers_server/app/shared/response"
)

func QuestsIndex(env *model.Env) *response.Message {
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

	qsts, err := env.DB.QuestsAll(pg)

	if err != nil {
		return response.InternalServerError()
	}

	if len(qsts) == 0 {
		return response.NotFound(fmt.Sprintf(formatNotFound, page))
	}

	return response.Ok(
		fmt.Sprintf(formatFound, quests),
		len(qsts),
		qsts)
}

func QuestSingle(env *model.Env) *response.Message {
	if env.Request.Method != "GET" {
		return response.MethodNotAllowed()
	}

	qstID, err := env.Request.Form.StringValue(id)

	if err != nil {
		return response.BadRequest(err.Error())
	}

	exists, err := env.DB.QuestExists(qstID)

	if err != nil {
		return response.InternalServerError()
	}

	if !exists {
		return response.NotFound(fmt.Sprintf(formatNotFound, quest))
	}

	qst, err := env.DB.QuestSingle(qstID)

	if err != nil {
		return response.InternalServerError()
	}

	return response.Ok(
		fmt.Sprintf(formatFound, quest),
		1,
		qst)
}

func QuestCreate(env *model.Env) *response.Message {
	if env.Request.Method != "POST" {
		return response.MethodNotAllowed()
	}

	qst := &model.Quest{}
	err := env.Request.Form.Map(qst)

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
		return response.NotFound(fmt.Sprintf(formatNotFound, involvementID))
	}

	questTypeExists, err := env.DB.QuestTypeExists(qst.QuestTypeID)

	if err != nil {
		return response.InternalServerError()
	}

	if !questTypeExists {
		return response.NotFound(fmt.Sprintf(formatNotFound, questTypeID))
	}

	qst.AuthorID = env.Request.UserID

	id, err := env.DB.QuestCreate(qst)

	if err != nil || id == "" {
		return response.InternalServerError()
	}

	return response.Ok(
		fmt.Sprintf(formatCreated, quest),
		1,
		id)
}
