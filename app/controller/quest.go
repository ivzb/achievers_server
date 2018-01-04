package controller

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/ivzb/achievers_server/app/model"
	"github.com/ivzb/achievers_server/app/shared/response"
)

func QuestsIndex(
	env *model.Env,
	w http.ResponseWriter,
	r *http.Request) response.Message {

	if r.Method != "GET" {
		return response.MethodNotAllowed(methodNotAllowed)
	}

	pg, err := strconv.Atoi(r.FormValue("page"))

	if err != nil {
		return response.BadRequest(fmt.Sprintf(formatMissing, page))
	}

	if pg < 0 {
		return response.BadRequest(fmt.Sprintf(formatInvalid, page))
	}

	qsts, err := env.DB.QuestsAll(pg)

	if err != nil {
		return response.InternalServerError(friendlyErrorMessage)
	}

	if len(qsts) == 0 {
		return response.NotFound(fmt.Sprintf(formatNotFound, page))
	}

	return response.Ok(
		fmt.Sprintf(formatFound, quests),
		len(qsts),
		qsts)
}

func QuestSingle(
	env *model.Env,
	w http.ResponseWriter,
	r *http.Request) response.Message {

	if r.Method != "GET" {
		return response.MethodNotAllowed(methodNotAllowed)
	}

	qstID := r.FormValue(id)

	if qstID == "" {
		return response.BadRequest(fmt.Sprintf(formatMissing, id))
	}

	exists, err := env.DB.QuestExists(qstID)

	if err != nil {
		return response.InternalServerError(friendlyErrorMessage)
	}

	if !exists {
		return response.NotFound(fmt.Sprintf(formatNotFound, quest))
	}

	qst, err := env.DB.QuestSingle(qstID)

	if err != nil {
		return response.InternalServerError(friendlyErrorMessage)
	}

	return response.Ok(
		fmt.Sprintf(formatFound, quest),
		1,
		qst)
}

func QuestCreate(
	env *model.Env,
	w http.ResponseWriter,
	r *http.Request) response.Message {

	if r.Method != "POST" {
		return response.MethodNotAllowed(methodNotAllowed)
	}

	qst := &model.Quest{}
	err := env.Former.Map(r, qst)

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
		return response.InternalServerError(friendlyErrorMessage)
	}

	if !involvementExists {
		return response.NotFound(fmt.Sprintf(formatNotFound, involvementID))
	}

	questTypeExists, err := env.DB.QuestTypeExists(qst.QuestTypeID)

	if err != nil {
		return response.InternalServerError(friendlyErrorMessage)
	}

	if !questTypeExists {
		return response.NotFound(fmt.Sprintf(formatNotFound, questTypeID))
	}

	qst.AuthorID = env.UserId

	id, err := env.DB.QuestCreate(qst)

	if err != nil || id == "" {
		return response.InternalServerError(friendlyErrorMessage)
	}

	return response.Ok(
		fmt.Sprintf(formatCreated, quest),
		1,
		id)
}
