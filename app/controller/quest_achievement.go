package controller

import (
	"fmt"
	"net/http"

	"github.com/ivzb/achievers_server/app/model"
	"github.com/ivzb/achievers_server/app/shared/response"
)

func QuestAchievementCreate(
	env *model.Env,
	w http.ResponseWriter,
	r *http.Request) *response.Message {

	if r.Method != "POST" {
		return response.MethodNotAllowed()
	}

	qstAch := &model.QuestAchievement{}
	err := env.Form.Map(r, qstAch)

	if err != nil {
		return response.BadRequest(err.Error())
	}

	if qstAch.QuestID == "" {
		return response.BadRequest(fmt.Sprintf(formatMissing, questID))
	}

	if qstAch.AchievementID == "" {
		return response.BadRequest(fmt.Sprintf(formatMissing, achievementID))
	}

	qstExists, err := env.DB.QuestExists(qstAch.QuestID)

	if err != nil {
		env.Log.Error(err)
		return response.InternalServerError()
	}

	if !qstExists {
		return response.NotFound(fmt.Sprintf(formatNotFound, questID))
	}

	achExists, err := env.DB.AchievementExists(qstAch.AchievementID)

	if err != nil {
		env.Log.Error(err)
		return response.InternalServerError()
	}

	if !achExists {
		return response.NotFound(fmt.Sprintf(formatNotFound, achievementID))
	}

	achQstExists, err := env.DB.QuestAchievementExists(qstAch.QuestID, qstAch.AchievementID)

	if err != nil {
		env.Log.Error(err)
		return response.InternalServerError()
	}

	if achQstExists {
		return response.BadRequest(fmt.Sprintf(formatAlreadyExists, questAchievement))
	}

	qstAch.AuthorID = env.UserId

	id, err := env.DB.QuestAchievementCreate(qstAch)

	if err != nil || id == "" {
		env.Log.Error(err)
		return response.InternalServerError()
	}

	return response.Ok(
		fmt.Sprintf(formatCreated, questAchievement),
		1,
		id)
}
