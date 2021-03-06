package controller

import (
	"fmt"

	"github.com/ivzb/achievers_server/app"
	"github.com/ivzb/achievers_server/app/model"
	"github.com/ivzb/achievers_server/app/shared/consts"
	"github.com/ivzb/achievers_server/app/shared/form"
	"github.com/ivzb/achievers_server/app/shared/request"
	"github.com/ivzb/achievers_server/app/shared/response"
	"github.com/ivzb/achievers_server/app/shared/validator"
)

func QuestAchievementCreate(env *app.Env) *response.Message {
	if !request.IsMethod(env.Request, consts.POST) {
		return response.MethodNotAllowed()
	}

	qstAch := &model.QuestAchievement{}
	err := form.ModelValue(env.Request, qstAch)

	if err != nil {
		return response.BadRequest(err.Error())
	}

	err = validator.Validate(*qstAch)

	if err != nil {
		return response.BadRequest(err.Error())
	}

	qstExists, err := env.DB.Quest().Exists(qstAch.QuestID)

	if err != nil {
		env.Log.Error(err)
		return response.InternalServerError()
	}

	if !qstExists {
		return response.NotFound(consts.QuestID)
	}

	achExists, err := env.DB.Achievement().Exists(qstAch.AchievementID)

	if err != nil {
		env.Log.Error(err)
		return response.InternalServerError()
	}

	if !achExists {
		return response.NotFound(consts.AchievementID)
	}

	achQstExists, err := env.DB.QuestAchievement().Exists(qstAch.QuestID, qstAch.AchievementID)

	if err != nil {
		env.Log.Error(err)
		return response.InternalServerError()
	}

	if achQstExists {
		return response.BadRequest(fmt.Sprintf(consts.FormatAlreadyExists, consts.QuestAchievement))
	}

	qstAch.UserID = env.UserID

	id, err := env.DB.QuestAchievement().Create(qstAch)

	if err != nil || id == "" {
		env.Log.Error(err)
		return response.InternalServerError()
	}

	return response.Created(
		consts.QuestAchievement,
		id)
}
