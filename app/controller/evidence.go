package controller

import (
	"fmt"
	"log"

	"github.com/ivzb/achievers_server/app/model"
	"github.com/ivzb/achievers_server/app/shared/form"
	"github.com/ivzb/achievers_server/app/shared/response"
)

func EvidencesIndex(env *model.Env) *response.Message {
	if env.Request.Method != "GET" {
		return response.MethodNotAllowed()
	}

	pg, err := form.IntValue(env.Request, "page")

	if err != nil {
		return response.BadRequest(fmt.Sprintf(formatMissing, page))
	}

	if pg < 0 {
		return response.BadRequest(fmt.Sprintf(formatInvalid, page))
	}

	evds, err := env.DB.EvidencesAll(pg)

	if err != nil {
		return response.InternalServerError()
	}

	if len(evds) == 0 {
		return response.NotFound(page)
	}

	return response.Ok(
		evidences,
		len(evds),
		evds)
}

func EvidenceSingle(env *model.Env) *response.Message {
	if env.Request.Method != "GET" {
		return response.MethodNotAllowed()
	}

	evdID, err := form.StringValue(env.Request, id)

	if err != nil {
		return response.BadRequest(err.Error())
	}

	exists, err := env.DB.EvidenceExists(evdID)

	if err != nil {
		return response.InternalServerError()
	}

	if !exists {
		return response.NotFound(evidence)
	}

	evd, err := env.DB.EvidenceSingle(evdID)

	if err != nil {
		return response.InternalServerError()
	}

	return response.Ok(
		evidence,
		1,
		evd)
}

func EvidenceCreate(env *model.Env) *response.Message {
	if env.Request.Method != "POST" {
		return response.MethodNotAllowed()
	}

	evd := &model.Evidence{}
	err := form.ModelValue(env.Request, evd)

	if err != nil {
		return response.BadRequest(err.Error())
	}

	if evd.Title == "" {
		return response.BadRequest(fmt.Sprintf(formatMissing, title))
	}

	if evd.PictureURL == "" {
		return response.BadRequest(fmt.Sprintf(formatMissing, pictureURL))
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
		return response.InternalServerError()
	}

	if !multimediaTypeExist {
		return response.NotFound(multimediaTypeID)
	}

	achievementExist, err := env.DB.AchievementExists(evd.AchievementID)

	if err != nil {
		return response.InternalServerError()
	}

	if !achievementExist {
		return response.NotFound(achievementID)
	}

	evd.AuthorID = env.UserID

	id, err := env.DB.EvidenceCreate(evd)

	if err != nil || id == "" {
		log.Println(err)
		return response.InternalServerError()
	}

	return response.Created(
		evidence,
		id)
}
