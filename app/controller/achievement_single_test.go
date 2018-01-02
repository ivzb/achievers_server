package controller

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/ivzb/achievers_server/app/model/mock"
)

type achievementSingleTest struct {
	purpose             string
	requestMethod       string
	responseType        int
	responseStatusCode  int
	responseMessage     string
	formID              string
	dbAchievementExists mock.AchievementExists
	dbAchievementSingle mock.AchievementSingle
}

var achievementSingleTests = []*test{
	constructAchievementSingleTest(&achievementSingleTest{
		purpose:             "invalid request method",
		requestMethod:       post,
		responseType:        Core,
		responseStatusCode:  http.StatusMethodNotAllowed,
		responseMessage:     methodNotAllowed,
		formID:              "",
		dbAchievementExists: mock.AchievementExists{},
		dbAchievementSingle: mock.AchievementSingle{},
	}),
	constructAchievementSingleTest(&achievementSingleTest{
		purpose:             "missing id",
		requestMethod:       get,
		responseType:        Core,
		responseStatusCode:  http.StatusBadRequest,
		responseMessage:     fmt.Sprintf(formatMissing, id),
		formID:              "",
		dbAchievementExists: mock.AchievementExists{},
		dbAchievementSingle: mock.AchievementSingle{},
	}),
	constructAchievementSingleTest(&achievementSingleTest{
		purpose:             "achievement exists db error",
		requestMethod:       get,
		responseType:        Core,
		responseStatusCode:  http.StatusInternalServerError,
		responseMessage:     friendlyErrorMessage,
		formID:              mockID,
		dbAchievementExists: mock.AchievementExists{Err: mockDbErr},
		dbAchievementSingle: mock.AchievementSingle{},
	}),
	constructAchievementSingleTest(&achievementSingleTest{
		purpose:             "achievement does not exist",
		requestMethod:       get,
		responseType:        Core,
		responseStatusCode:  http.StatusNotFound,
		responseMessage:     fmt.Sprintf(formatNotFound, achievement),
		formID:              mockID,
		dbAchievementExists: mock.AchievementExists{Bool: false},
		dbAchievementSingle: mock.AchievementSingle{},
	}),
	constructAchievementSingleTest(&achievementSingleTest{
		purpose:             "achievement single db error",
		requestMethod:       get,
		responseType:        Core,
		responseStatusCode:  http.StatusInternalServerError,
		responseMessage:     friendlyErrorMessage,
		formID:              mockID,
		dbAchievementExists: mock.AchievementExists{Bool: true},
		dbAchievementSingle: mock.AchievementSingle{Err: mockDbErr},
	}),
	constructAchievementSingleTest(&achievementSingleTest{
		purpose:             "achievement single OK",
		requestMethod:       get,
		responseType:        Retrieve,
		responseStatusCode:  http.StatusOK,
		responseMessage:     fmt.Sprintf(formatFound, achievement),
		formID:              mockID,
		dbAchievementExists: mock.AchievementExists{Bool: true},
		dbAchievementSingle: mock.AchievementSingle{Ach: mock.Achievement()},
	}),
}

func constructAchievementSingleTest(testInput *achievementSingleTest) *test {
	responseResults, _ := json.Marshal(mock.Achievement())

	db := &mock.DB{
		AchievementExistsMock: testInput.dbAchievementExists,
		AchievementSingleMock: testInput.dbAchievementSingle,
	}

	logger := &mock.Logger{}

	return &test{
		purpose: testInput.purpose,
		handle:  AchievementSingle,
		request: constructTestRequest(
			testInput.requestMethod,
			constructForm(map[string]string{
				id: testInput.formID,
			}),
			constructEnv(db, logger, nil),
		),
		response: constructTestResponse(
			testInput.responseType,
			testInput.responseStatusCode,
			testInput.responseMessage,
			responseResults,
		),
	}
}
