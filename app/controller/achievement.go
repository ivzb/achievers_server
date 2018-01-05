package controller

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/ivzb/achievers_server/app/model"
	"github.com/ivzb/achievers_server/app/shared/response"
)

func AchievementsIndex(
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

	achs, err := env.DB.AchievementsAll(pg)

	if err != nil {
		env.Logger.Error(err)
		return response.InternalServerError(friendlyErrorMessage)
	}

	if len(achs) == 0 {
		return response.NotFound(fmt.Sprintf(formatNotFound, page))
	}

	return response.Ok(
		fmt.Sprintf(formatFound, achievements),
		len(achs),
		achs)
}

func AchievementsByQuestID(
	env *model.Env,
	w http.ResponseWriter,
	r *http.Request) response.Message {

	if r.Method != "GET" {
		return response.MethodNotAllowed(methodNotAllowed)
	}

	pg, err := strconv.Atoi(r.FormValue(page))

	if err != nil {
		return response.BadRequest(fmt.Sprintf(formatMissing, page))
	}

	if pg < 0 {
		return response.BadRequest(fmt.Sprintf(formatInvalid, page))
	}

	qstID := r.FormValue(id)

	if qstID == "" {
		return response.BadRequest(fmt.Sprintf(formatMissing, id))
	}

	qstExists, err := env.DB.QuestExists(qstID)

	if err != nil {
		env.Logger.Error(err)
		return response.InternalServerError(friendlyErrorMessage)
	}

	if !qstExists {
		return response.NotFound(fmt.Sprintf(formatNotFound, id))
	}

	achs, err := env.DB.AchievementsByQuestID(qstID, pg)

	if err != nil {
		env.Logger.Error(err)
		return response.InternalServerError(friendlyErrorMessage)
	}

	if len(achs) == 0 {
		return response.NotFound(fmt.Sprintf(formatNotFound, page))
	}

	return response.Ok(
		fmt.Sprintf(formatFound, achievements),
		len(achs),
		achs)
}

func AchievementSingle(
	env *model.Env,
	w http.ResponseWriter,
	r *http.Request) response.Message {

	if r.Method != "GET" {
		return response.MethodNotAllowed(methodNotAllowed)
	}

	achid := r.FormValue(id)

	if achid == "" {
		return response.BadRequest(fmt.Sprintf(formatMissing, id))
	}

	exists, err := env.DB.AchievementExists(achid)

	if err != nil {
		env.Logger.Error(err)
		return response.InternalServerError(friendlyErrorMessage)
	}

	if !exists {
		return response.NotFound(fmt.Sprintf(formatNotFound, achievement))
	}

	ach, err := env.DB.AchievementSingle(achid)

	if err != nil {
		env.Logger.Error(err)
		return response.InternalServerError(friendlyErrorMessage)
	}

	return response.Ok(
		fmt.Sprintf(formatFound, achievement),
		1,
		ach)
}

func AchievementCreate(
	env *model.Env,
	w http.ResponseWriter,
	r *http.Request) response.Message {

	if r.Method != "POST" {
		return response.MethodNotAllowed(methodNotAllowed)
	}

	ach := &model.Achievement{}
	err := env.Former.Map(r, ach)

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
		env.Logger.Error(err)
		return response.InternalServerError(friendlyErrorMessage)
	}

	if !involvementExists {
		return response.NotFound(fmt.Sprintf(formatNotFound, involvementID))
	}

	ach.AuthorID = env.UserId

	id, err := env.DB.AchievementCreate(ach)

	if err != nil || id == "" {
		return response.InternalServerError(friendlyErrorMessage)
	}

	return response.Ok(
		fmt.Sprintf(formatCreated, achievement),
		1,
		id)
}
