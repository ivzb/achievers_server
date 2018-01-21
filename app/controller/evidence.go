package controller

import (
	"fmt"
	"log"

	"github.com/ivzb/achievers_server/app/model"
	"github.com/ivzb/achievers_server/app/shared/consts"
	"github.com/ivzb/achievers_server/app/shared/env"
	"github.com/ivzb/achievers_server/app/shared/form"
	"github.com/ivzb/achievers_server/app/shared/response"
)

func EvidencesIndex(env *env.Env) *response.Message {
	if env.Request.Method != "GET" {
		return response.MethodNotAllowed()
	}

	pg, err := form.IntValue(env.Request, "page")

	if err != nil {
		return response.BadRequest(fmt.Sprintf(consts.FormatMissing, consts.Page))
	}

	if pg < 0 {
		return response.BadRequest(fmt.Sprintf(consts.FormatInvalid, consts.Page))
	}

	evds, err := env.DB.Evidence().All(pg)

	if err != nil {
		return response.InternalServerError()
	}

	if len(evds) == 0 {
		return response.NotFound(consts.Page)
	}

	return response.Ok(
		consts.Evidences,
		len(evds),
		evds)
}

func EvidenceSingle(env *env.Env) *response.Message {
	if env.Request.Method != consts.GET {
		return response.MethodNotAllowed()
	}

	evdID, err := form.StringValue(env.Request, consts.ID)

	if err != nil {
		return response.BadRequest(err.Error())
	}

	exists, err := env.DB.Evidence().Exists(evdID)

	if err != nil {
		return response.InternalServerError()
	}

	if !exists {
		return response.NotFound(consts.Evidence)
	}

	evd, err := env.DB.Evidence().Single(evdID)

	if err != nil {
		return response.InternalServerError()
	}

	return response.Ok(
		consts.Evidence,
		1,
		evd)
}

func EvidenceCreate(env *env.Env) *response.Message {
	if env.Request.Method != "POST" {
		return response.MethodNotAllowed()
	}

	evd := &model.Evidence{}
	err := form.ModelValue(env.Request, evd)

	if err != nil {
		return response.BadRequest(err.Error())
	}

	if evd.Title == "" {
		return response.BadRequest(fmt.Sprintf(consts.FormatMissing, consts.Title))
	}

	if evd.PictureURL == "" {
		return response.BadRequest(fmt.Sprintf(consts.FormatMissing, consts.PictureURL))
	}

	if evd.URL == "" {
		return response.BadRequest(fmt.Sprintf(consts.FormatMissing, consts.URL))
	}

	if evd.MultimediaTypeID == 0 {
		return response.BadRequest(fmt.Sprintf(consts.FormatMissing, consts.MultimediaTypeID))
	}

	if evd.AchievementID == "" {
		return response.BadRequest(fmt.Sprintf(consts.FormatMissing, consts.AchievementID))
	}

	multimediaTypeExist, err := env.DB.MultimediaType().Exists(evd.MultimediaTypeID)

	if err != nil {
		return response.InternalServerError()
	}

	if !multimediaTypeExist {
		return response.NotFound(consts.MultimediaTypeID)
	}

	achievementExist, err := env.DB.Achievement().Exists(evd.AchievementID)

	if err != nil {
		return response.InternalServerError()
	}

	if !achievementExist {
		return response.NotFound(consts.AchievementID)
	}

	evd.UserID = env.UserID

	id, err := env.DB.Evidence().Create(evd)

	if err != nil || id == "" {
		log.Println(err)
		return response.InternalServerError()
	}

	return response.Created(
		consts.Evidence,
		id)
}
