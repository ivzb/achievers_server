package controller

import (
	"fmt"
	"net/http"

	"github.com/ivzb/achievers_server/app/model"
	"github.com/ivzb/achievers_server/app/shared/response"
)

func QuestAchievementSingle(
	env *model.Env,
	w http.ResponseWriter,
	r *http.Request) response.Message {

	if r.Method != "GET" {
		return response.MethodNotAllowed(methodNotAllowed)
	}

	qstID := r.FormValue(questID)

	if len(qstID) == 0 {
		return response.BadRequest(fmt.Sprintf(formatMissing, questID))
	}

	achID := r.FormValue(achievementID)

	if len(achID) == 0 {
		return response.BadRequest(fmt.Sprintf(formatMissing, achievementID))
	}

	qstExists, err := env.DB.QuestExists(qstID)

	if err != nil {
		env.Logger.Error(err)
		return response.InternalServerError(friendlyErrorMessage)
	}

	if !qstExists {
		return response.NotFound(fmt.Sprintf(formatNotFound, questID))
	}

	achExists, err := env.DB.AchievementExists(achID)

	if err != nil {
		env.Logger.Error(err)
		return response.InternalServerError(friendlyErrorMessage)
	}

	if !achExists {
		return response.NotFound(fmt.Sprintf(formatNotFound, achievementID))
	}

	exists, err := env.DB.QuestAchievementExists(qstID, achID)

	if err != nil {
		env.Logger.Error(err)
		return response.InternalServerError(friendlyErrorMessage)
	}

	if !exists {
		return response.NotFound(fmt.Sprintf(formatNotFound, questAchievement))
	}

	qstAch, err := env.DB.QuestAchievementSingle(qstID, achID)

	if err != nil {
		env.Logger.Error(err)
		return response.InternalServerError(friendlyErrorMessage)
	}

	return response.Ok(
		fmt.Sprintf(formatFound, questAchievement),
		1,
		qstAch)
}
