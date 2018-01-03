package controller

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/ivzb/achievers_server/app/model/mock"
)

type achievementsIndexTest struct {
	purpose            string
	requestMethod      string
	responseType       int
	responseStatusCode int
	responseMessage    string
	mockPageSize       int
	formPage           string
	dbAchievementsAll  mock.AchievementsAll
}

var achievementsIndexTests = []*test{
	constructAchievementsIndexTest(&achievementsIndexTest{
		purpose:            "invalid request method",
		requestMethod:      post,
		responseType:       Core,
		responseStatusCode: http.StatusMethodNotAllowed,
		responseMessage:    methodNotAllowed,
	}),
	constructAchievementsIndexTest(&achievementsIndexTest{
		purpose:            "9 results on page",
		requestMethod:      get,
		responseType:       Retrieve,
		responseStatusCode: http.StatusOK,
		responseMessage:    fmt.Sprintf(formatFound, achievements),
		mockPageSize:       9,
		formPage:           "0",
		dbAchievementsAll:  mock.AchievementsAll{Achs: mock.Achievements(9)},
	}),
	constructAchievementsIndexTest(&achievementsIndexTest{
		purpose:            "4 results on page",
		requestMethod:      get,
		responseType:       Retrieve,
		responseStatusCode: http.StatusOK,
		responseMessage:    fmt.Sprintf(formatFound, achievements),
		mockPageSize:       4,
		formPage:           "1",
		dbAchievementsAll:  mock.AchievementsAll{Achs: mock.Achievements(4)},
	}),
	constructAchievementsIndexTest(&achievementsIndexTest{
		purpose:            "no results on page",
		requestMethod:      get,
		responseType:       Core,
		responseStatusCode: http.StatusNotFound,
		responseMessage:    fmt.Sprintf(formatNotFound, page),
		mockPageSize:       0,
		formPage:           "2",
		dbAchievementsAll:  mock.AchievementsAll{Achs: mock.Achievements(0)},
	}),
	constructAchievementsIndexTest(&achievementsIndexTest{
		purpose:            "missing page",
		requestMethod:      get,
		responseType:       Core,
		responseStatusCode: http.StatusBadRequest,
		responseMessage:    fmt.Sprintf(formatMissing, page),
		formPage:           "",
	}),
	constructAchievementsIndexTest(&achievementsIndexTest{
		purpose:            "invalid page",
		requestMethod:      get,
		responseType:       Core,
		responseStatusCode: http.StatusBadRequest,
		responseMessage:    fmt.Sprintf(formatInvalid, page),
		formPage:           "-1",
	}),
	constructAchievementsIndexTest(&achievementsIndexTest{
		purpose:            "db error",
		requestMethod:      get,
		responseType:       Core,
		responseStatusCode: http.StatusInternalServerError,
		responseMessage:    friendlyErrorMessage,
		formPage:           mockPage,
		dbAchievementsAll:  mock.AchievementsAll{Err: mockDbErr},
	}),
}

func constructAchievementsIndexTest(testInput *achievementsIndexTest) *test {
	responseResults, _ := json.Marshal(mock.Achievements(testInput.mockPageSize))

	db := &mock.DB{AchievementsAllMock: testInput.dbAchievementsAll}

	logger := &mock.Logger{}

	return &test{
		purpose: testInput.purpose,
		handle:  AchievementsIndex,
		request: constructTestRequest(
			testInput.requestMethod,
			constructForm(map[string]string{page: testInput.formPage}),
			constructEnv(db, logger, nil, nil),
		),
		response: constructTestResponse(
			testInput.responseType,
			testInput.responseStatusCode,
			testInput.responseMessage,
			responseResults,
		),
	}
}
