package controller

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/ivzb/achievers_server/app/model"
	"github.com/ivzb/achievers_server/app/shared/response"
)

const (
	id            = "id"
	achievement   = "achievement"
	achievements  = "achievements"
	involvement   = "involvement"
	page          = "page"
	title         = "title"
	description   = "description"
	pictureURL    = "picture_url"
	involvementID = "involvement_id"
	authorID      = "author_id"
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

	achID := r.FormValue(id)

	if achID == "" {
		return response.BadRequest(fmt.Sprintf(formatMissing, id))
	}

	exists, err := env.DB.Exists(achievement, id, achID)

	if err != nil {
		return response.InternalServerError(friendlyErrorMessage)
	}

	if !exists {
		return response.NotFound(fmt.Sprintf(formatNotFound, achievement))
	}

	ach, err := env.DB.AchievementSingle(achID)

	if err != nil {
		return response.InternalServerError(friendlyErrorMessage)
	}

	if ach == nil {
		return response.NotFound(fmt.Sprintf(formatNotFound, achievement))
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

	ach := &model.Achievement{
		Title:         r.FormValue("title"),
		Description:   r.FormValue("description"),
		PictureUrl:    r.FormValue("picture_url"),
		InvolvementId: r.FormValue("involvement_id"),
	}

	if ach.Title == "" {
		return response.BadRequest(fmt.Sprintf(formatMissing, title))
	}

	if ach.Description == "" {
		return response.BadRequest(fmt.Sprintf(formatMissing, description))
	}

	if ach.PictureUrl == "" {
		return response.BadRequest(fmt.Sprintf(formatMissing, pictureURL))
	}

	if ach.InvolvementId == "" {
		return response.BadRequest(fmt.Sprintf(formatMissing, involvementID))
	}

	involvementExists, err := env.DB.Exists("involvement", "id", ach.InvolvementId)

	if err != nil {
		return response.InternalServerError(friendlyErrorMessage)
	}

	if !involvementExists {
		return response.BadRequest(fmt.Sprintf(formatNotFound, involvement))
	}

	ach.AuthorId = env.UserId

	id, err := env.DB.AchievementCreate(ach)

	if err != nil || id == "" {
		return response.InternalServerError(friendlyErrorMessage)
	}

	return response.Ok(
		fmt.Sprintf(formatFound, achievement),
		1,
		id)
}
