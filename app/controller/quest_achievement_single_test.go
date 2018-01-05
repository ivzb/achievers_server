package controller

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/ivzb/achievers_server/app/model/mock"
)

func questAchievementSingleForm() *map[string]string {
	return &map[string]string{
		questID:       mockID,
		achievementID: mockID,
	}
}

var questAchievementSingleTests = []*test{
	constructQuestAchievementSingleTest(&testInput{
		purpose:            "invalid request method",
		requestMethod:      post,
		responseType:       Core,
		responseStatusCode: http.StatusMethodNotAllowed,
		responseMessage:    methodNotAllowed,
		form:               &map[string]string{},
	}),
	constructQuestAchievementSingleTest(&testInput{
		purpose:            "missing questID",
		requestMethod:      get,
		responseType:       Core,
		responseStatusCode: http.StatusBadRequest,
		responseMessage:    fmt.Sprintf(formatMissing, questID),
		form:               mapWithout(questAchievementSingleForm(), questID),
	}),
	constructQuestAchievementSingleTest(&testInput{
		purpose:            "missing achievementID",
		requestMethod:      get,
		responseType:       Core,
		responseStatusCode: http.StatusBadRequest,
		responseMessage:    fmt.Sprintf(formatMissing, achievementID),
		form:               mapWithout(questAchievementSingleForm(), achievementID),
	}),
	constructQuestAchievementSingleTest(&testInput{
		purpose:            "quest_id exists db error",
		requestMethod:      get,
		responseType:       Core,
		responseStatusCode: http.StatusInternalServerError,
		responseMessage:    friendlyErrorMessage,
		form:               questAchievementSingleForm(),
		db: &mock.DB{
			QuestExistsMock: mock.QuestExists{Err: mockDbErr},
		},
	}),
	constructQuestAchievementSingleTest(&testInput{
		purpose:            "quest_id does not exist",
		requestMethod:      get,
		responseType:       Core,
		responseStatusCode: http.StatusNotFound,
		responseMessage:    fmt.Sprintf(formatNotFound, questID),
		form:               questAchievementSingleForm(),
		db: &mock.DB{
			QuestExistsMock: mock.QuestExists{Bool: false},
		},
	}),
	constructQuestAchievementSingleTest(&testInput{
		purpose:            "achievement_id exists db error",
		requestMethod:      get,
		responseType:       Core,
		responseStatusCode: http.StatusInternalServerError,
		responseMessage:    friendlyErrorMessage,
		form:               questAchievementSingleForm(),
		db: &mock.DB{
			QuestExistsMock:       mock.QuestExists{Bool: true},
			AchievementExistsMock: mock.AchievementExists{Err: mockDbErr},
		},
	}),
	constructQuestAchievementSingleTest(&testInput{
		purpose:            "achievement_id does not exist",
		requestMethod:      get,
		responseType:       Core,
		responseStatusCode: http.StatusNotFound,
		responseMessage:    fmt.Sprintf(formatNotFound, achievementID),
		form:               questAchievementSingleForm(),
		db: &mock.DB{
			QuestExistsMock:       mock.QuestExists{Bool: true},
			AchievementExistsMock: mock.AchievementExists{Bool: false},
		},
	}),
	constructQuestAchievementSingleTest(&testInput{
		purpose:            "quest_achievement exists db error",
		requestMethod:      get,
		responseType:       Core,
		responseStatusCode: http.StatusInternalServerError,
		responseMessage:    friendlyErrorMessage,
		form:               questAchievementSingleForm(),
		db: &mock.DB{
			QuestExistsMock:            mock.QuestExists{Bool: true},
			AchievementExistsMock:      mock.AchievementExists{Bool: true},
			QuestAchievementExistsMock: mock.QuestAchievementExists{Err: mockDbErr},
		},
	}),
	constructQuestAchievementSingleTest(&testInput{
		purpose:            "quest_achievement does not exist",
		requestMethod:      get,
		responseType:       Core,
		responseStatusCode: http.StatusNotFound,
		responseMessage:    fmt.Sprintf(formatNotFound, questAchievement),
		form:               questAchievementSingleForm(),
		db: &mock.DB{
			QuestExistsMock:            mock.QuestExists{Bool: true},
			AchievementExistsMock:      mock.AchievementExists{Bool: true},
			QuestAchievementExistsMock: mock.QuestAchievementExists{Bool: false},
		},
	}),
	constructQuestAchievementSingleTest(&testInput{
		purpose:            "quest_achievement single db error",
		requestMethod:      get,
		responseType:       Core,
		responseStatusCode: http.StatusInternalServerError,
		responseMessage:    friendlyErrorMessage,
		form:               questAchievementSingleForm(),
		db: &mock.DB{
			QuestExistsMock:            mock.QuestExists{Bool: true},
			AchievementExistsMock:      mock.AchievementExists{Bool: true},
			QuestAchievementExistsMock: mock.QuestAchievementExists{Bool: true},
			QuestAchievementSingleMock: mock.QuestAchievementSingle{Err: mockDbErr},
		},
	}),
	constructQuestAchievementSingleTest(&testInput{
		purpose:            "quest_achievement single OK",
		requestMethod:      get,
		responseType:       Retrieve,
		responseStatusCode: http.StatusOK,
		responseMessage:    fmt.Sprintf(formatFound, questAchievement),
		form:               questAchievementSingleForm(),
		db: &mock.DB{
			QuestExistsMock:            mock.QuestExists{Bool: true},
			AchievementExistsMock:      mock.AchievementExists{Bool: true},
			QuestAchievementExistsMock: mock.QuestAchievementExists{Bool: true},
			QuestAchievementSingleMock: mock.QuestAchievementSingle{QstAch: mock.QuestAchievement()},
		},
	}),
}

func constructQuestAchievementSingleTest(testInput *testInput) *test {
	responseResults, _ := json.Marshal(mock.QuestAchievement())

	return constructTest(QuestAchievementSingle, testInput, responseResults)
}
