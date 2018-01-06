package controller

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/ivzb/achievers_server/app/model/mock"
)

func questAchievementCreateForm() *map[string]string {
	return &map[string]string{
		questID:       mockID,
		achievementID: mockID,
	}
}

var questAchievementCreateTests = []*test{
	constructQuestAchievementCreateTest(&testInput{
		purpose:            "invalid request method",
		requestMethod:      GET,
		responseType:       Core,
		responseStatusCode: http.StatusMethodNotAllowed,
		responseMessage:    methodNotAllowed,
		form:               &map[string]string{},
	}),
	constructQuestAchievementCreateTest(&testInput{
		purpose:            "former error",
		requestMethod:      POST,
		responseType:       Core,
		responseStatusCode: http.StatusBadRequest,
		responseMessage:    "former error",
		form:               &map[string]string{},
		former:             &mock.Former{MapMock: mock.Map{Err: mockFormerErr}},
	}),
	constructQuestAchievementCreateTest(&testInput{
		purpose:            "missing form quest_id",
		requestMethod:      POST,
		responseType:       Core,
		responseStatusCode: http.StatusBadRequest,
		responseMessage:    fmt.Sprintf(formatMissing, questID),
		form:               mapWithout(questAchievementCreateForm(), questID),
		former:             &mock.Former{},
	}),
	constructQuestAchievementCreateTest(&testInput{
		purpose:            "missing form achievement_id",
		requestMethod:      POST,
		responseType:       Core,
		responseStatusCode: http.StatusBadRequest,
		responseMessage:    fmt.Sprintf(formatMissing, achievementID),
		form:               mapWithout(questAchievementCreateForm(), achievementID),
		former:             &mock.Former{},
	}),
	constructQuestAchievementCreateTest(&testInput{
		purpose:            "quest exists db error",
		requestMethod:      POST,
		responseType:       Core,
		responseStatusCode: http.StatusInternalServerError,
		responseMessage:    friendlyErrorMessage,
		form:               questAchievementCreateForm(),
		former:             &mock.Former{},
		db: &mock.DB{
			QuestExistsMock: mock.QuestExists{Err: mockDbErr},
		},
	}),
	constructQuestAchievementCreateTest(&testInput{
		purpose:            "quest_id does not exist",
		requestMethod:      POST,
		responseType:       Core,
		responseStatusCode: http.StatusNotFound,
		responseMessage:    fmt.Sprintf(formatNotFound, questID),
		form:               questAchievementCreateForm(),
		former:             &mock.Former{},
		db: &mock.DB{
			QuestExistsMock: mock.QuestExists{Bool: false},
		},
	}),
	constructQuestAchievementCreateTest(&testInput{
		purpose:            "achievement exists db error",
		requestMethod:      POST,
		responseType:       Core,
		responseStatusCode: http.StatusInternalServerError,
		responseMessage:    friendlyErrorMessage,
		form:               questAchievementCreateForm(),
		former:             &mock.Former{},
		db: &mock.DB{
			QuestExistsMock:       mock.QuestExists{Bool: true},
			AchievementExistsMock: mock.AchievementExists{Err: mockDbErr},
		},
	}),
	constructQuestAchievementCreateTest(&testInput{
		purpose:            "achievement_id does not exist",
		requestMethod:      POST,
		responseType:       Core,
		responseStatusCode: http.StatusNotFound,
		responseMessage:    fmt.Sprintf(formatNotFound, achievementID),
		form:               questAchievementCreateForm(),
		former:             &mock.Former{},
		db: &mock.DB{
			QuestExistsMock:       mock.QuestExists{Bool: true},
			AchievementExistsMock: mock.AchievementExists{Bool: false},
		},
	}),
	constructQuestAchievementCreateTest(&testInput{
		purpose:            "quest_achievement exists db error",
		requestMethod:      POST,
		responseType:       Core,
		responseStatusCode: http.StatusInternalServerError,
		responseMessage:    friendlyErrorMessage,
		form:               questAchievementCreateForm(),
		former:             &mock.Former{},
		db: &mock.DB{
			QuestExistsMock:            mock.QuestExists{Bool: true},
			AchievementExistsMock:      mock.AchievementExists{Bool: true},
			QuestAchievementExistsMock: mock.QuestAchievementExists{Err: mockDbErr},
		},
	}),
	constructQuestAchievementCreateTest(&testInput{
		purpose:            "quest_achievement already exists",
		requestMethod:      POST,
		responseType:       Core,
		responseStatusCode: http.StatusBadRequest,
		responseMessage:    fmt.Sprintf(formatAlreadyExists, questAchievement),
		form:               questAchievementCreateForm(),
		former:             &mock.Former{},
		db: &mock.DB{
			QuestExistsMock:            mock.QuestExists{Bool: true},
			AchievementExistsMock:      mock.AchievementExists{Bool: true},
			QuestAchievementExistsMock: mock.QuestAchievementExists{Bool: true},
		},
	}),
	constructQuestAchievementCreateTest(&testInput{
		purpose:            "quest_achievement create db error",
		requestMethod:      POST,
		responseType:       Core,
		responseStatusCode: http.StatusInternalServerError,
		responseMessage:    friendlyErrorMessage,
		form:               questAchievementCreateForm(),
		former:             &mock.Former{},
		db: &mock.DB{
			QuestExistsMock:            mock.QuestExists{Bool: true},
			AchievementExistsMock:      mock.AchievementExists{Bool: true},
			QuestAchievementExistsMock: mock.QuestAchievementExists{Bool: false},
			QuestAchievementCreateMock: mock.QuestAchievementCreate{Err: mockDbErr},
		},
	}),
	constructQuestAchievementCreateTest(&testInput{
		purpose:            "quest_achievement create ok",
		requestMethod:      POST,
		responseType:       Retrieve,
		responseStatusCode: http.StatusOK,
		responseMessage:    fmt.Sprintf(formatCreated, questAchievement),
		form:               questAchievementCreateForm(),
		former:             &mock.Former{},
		db: &mock.DB{
			QuestExistsMock:            mock.QuestExists{Bool: true},
			AchievementExistsMock:      mock.AchievementExists{Bool: true},
			QuestAchievementExistsMock: mock.QuestAchievementExists{Bool: false},
			QuestAchievementCreateMock: mock.QuestAchievementCreate{ID: mockID},
		},
	}),
}

func constructQuestAchievementCreateTest(testInput *testInput) *test {
	responseResults, _ := json.Marshal(mockID)
	return constructTest(QuestAchievementCreate, testInput, responseResults)
}
