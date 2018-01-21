package controller

import (
	"fmt"

	"github.com/ivzb/achievers_server/app/model"
	"github.com/ivzb/achievers_server/app/shared/consts"
	"github.com/ivzb/achievers_server/app/shared/env"
	"github.com/ivzb/achievers_server/app/shared/form"
	"github.com/ivzb/achievers_server/app/shared/request"
	"github.com/ivzb/achievers_server/app/shared/response"
)

func QuestsIndex(env *env.Env) *response.Message {
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

	qsts, err := env.DB.Quest().All(pg)

	if err != nil {
		env.Log.Error(err)
		return response.InternalServerError()
	}

	if len(qsts) == 0 {
		return response.NotFound(consts.Page)
	}

	return response.Ok(
		consts.Quests,
		len(qsts),
		qsts)
}

func QuestSingle(env *env.Env) *response.Message {
	if env.Request.Method != consts.GET {
		return response.MethodNotAllowed()
	}

	qstID, err := form.StringValue(env.Request, consts.ID)

	if err != nil {
		return response.BadRequest(err.Error())
	}

	exists, err := env.DB.Quest().Exists(qstID)

	if err != nil {
		env.Log.Error(err)
		return response.InternalServerError()
	}

	if !exists {
		return response.NotFound(consts.Quest)
	}

	qst, err := env.DB.Quest().Single(qstID)

	if err != nil {
		env.Log.Error(err)
		return response.InternalServerError()
	}

	return response.Ok(
		consts.Quest,
		1,
		qst)
}

func QuestCreate(env *env.Env) *response.Message {
	if env.Request.Method != "POST" {
		return response.MethodNotAllowed()
	}

	qst := &model.Quest{}
	err := form.ModelValue(env.Request, qst)

	if err != nil {
		return response.BadRequest(err.Error())
	}

	if qst.Title == "" {
		return response.BadRequest(fmt.Sprintf(consts.FormatMissing, consts.Title))
	}

	if qst.PictureURL == "" {
		return response.BadRequest(fmt.Sprintf(consts.FormatMissing, consts.PictureURL))
	}

	if qst.InvolvementID == "" {
		return response.BadRequest(fmt.Sprintf(consts.FormatMissing, consts.InvolvementID))
	}

	if qst.QuestTypeID == 0 {
		return response.BadRequest(fmt.Sprintf(consts.FormatMissing, consts.QuestTypeID))
	}

	involvementExists, err := env.DB.Involvement().Exists(qst.InvolvementID)

	if err != nil {
		env.Log.Error(err)
		return response.InternalServerError()
	}

	if !involvementExists {
		return response.NotFound(consts.InvolvementID)
	}

	questTypeExists, err := env.DB.QuestType().Exists(qst.QuestTypeID)

	if err != nil {
		env.Log.Error(err)
		return response.InternalServerError()
	}

	if !questTypeExists {
		return response.NotFound(consts.QuestTypeID)
	}

	qst.UserID = env.UserID

	id, err := env.DB.Quest().Create(qst)

	if err != nil || id == "" {
		env.Log.Error(err)
		return response.InternalServerError()
	}

	return response.Created(
		consts.Quest,
		id)
}
