package controller

import (
	"fmt"
	"log"
	"net/http"
	"strconv"

	"github.com/ivzb/achievers_server/app/model"
	"github.com/ivzb/achievers_server/app/shared/response"
)

func EvidencesIndex(
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

	evds, err := env.DB.EvidencesAll(pg)

	if err != nil {
		return response.InternalServerError(friendlyErrorMessage)
	}

	if len(evds) == 0 {
		return response.NotFound(fmt.Sprintf(formatNotFound, page))
	}

	return response.Ok(
		fmt.Sprintf(formatFound, evidences),
		len(evds),
		evds)
}

func EvidenceSingle(
	env *model.Env,
	w http.ResponseWriter,
	r *http.Request) response.Message {

	if r.Method != "GET" {
		return response.MethodNotAllowed(methodNotAllowed)
	}

	evdID := r.FormValue(id)

	if evdID == "" {
		return response.BadRequest(fmt.Sprintf(formatMissing, id))
	}

	exists, err := env.DB.EvidenceExists(evdID)

	if err != nil {
		return response.InternalServerError(friendlyErrorMessage)
	}

	if !exists {
		return response.NotFound(fmt.Sprintf(formatNotFound, evidence))
	}

	evd, err := env.DB.EvidenceSingle(evdID)

	if err != nil {
		return response.InternalServerError(friendlyErrorMessage)
	}

	if evd == nil {
		return response.NotFound(fmt.Sprintf(formatNotFound, evidence))
	}

	return response.Ok(
		fmt.Sprintf(formatFound, evidence),
		1,
		evd)
}

func EvidenceCreate(
	env *model.Env,
	w http.ResponseWriter,
	r *http.Request) response.Message {

	if r.Method != "POST" {
		return response.MethodNotAllowed(methodNotAllowed)
	}

	evd := &model.Evidence{}
	err := env.Former.Map(r, evd)

	if err != nil {
		return response.BadRequest(err.Error())
	}

	if evd.Description == "" {
		return response.BadRequest(fmt.Sprintf(formatMissing, description))
	}

	if evd.PreviewURL == "" {
		return response.BadRequest(fmt.Sprintf(formatMissing, previewURL))
	}

	if evd.URL == "" {
		return response.BadRequest(fmt.Sprintf(formatMissing, _url))
	}

	if evd.MultimediaTypeID == 0 {
		return response.BadRequest(fmt.Sprintf(formatMissing, multimediaTypeID))
	}

	if evd.AchievementID == "" {
		return response.BadRequest(fmt.Sprintf(formatMissing, achievementID))
	}

	multimediaTypeExist, err := env.DB.MultimediaTypeExists(evd.MultimediaTypeID)

	if err != nil {
		return response.InternalServerError(friendlyErrorMessage)
	}

	if !multimediaTypeExist {
		return response.BadRequest(fmt.Sprintf(formatNotFound, multimediaType))
	}

	achievementExist, err := env.DB.AchievementExists(evd.AchievementID)

	if err != nil {
		return response.InternalServerError(friendlyErrorMessage)
	}

	if !achievementExist {
		return response.BadRequest(fmt.Sprintf(formatNotFound, achievement))
	}

	evd.AuthorID = env.UserId

	id, err := env.DB.EvidenceCreate(evd)

	if err != nil || id == "" {
		log.Println(err)
		return response.InternalServerError(friendlyErrorMessage)
	}

	return response.Ok(
		fmt.Sprintf(formatCreated, evidence),
		1,
		id)
}
