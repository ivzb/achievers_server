package controller

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/ivzb/achievers_server/app/model/mock"
)

type rewardsIndexTest struct {
	purpose            string
	requestMethod      string
	responseType       int
	responseStatusCode int
	responseMessage    string
	mockPageSize       int
	formPage           string
	dbRewardsAll       mock.RewardsAll
}

var rewardsIndexTests = []*test{
	constructRewardsIndexTest(&rewardsIndexTest{
		purpose:            "invalid request method",
		requestMethod:      post,
		responseType:       Core,
		responseStatusCode: http.StatusMethodNotAllowed,
		responseMessage:    methodNotAllowed,
	}),
	constructRewardsIndexTest(&rewardsIndexTest{
		purpose:            "9 results on page",
		requestMethod:      get,
		responseType:       Retrieve,
		responseStatusCode: http.StatusOK,
		responseMessage:    fmt.Sprintf(formatFound, rewards),
		mockPageSize:       9,
		formPage:           "0",
		dbRewardsAll:       mock.RewardsAll{Rwds: mock.Rewards(9)},
	}),
	constructRewardsIndexTest(&rewardsIndexTest{
		purpose:            "4 results on page",
		requestMethod:      get,
		responseType:       Retrieve,
		responseStatusCode: http.StatusOK,
		responseMessage:    fmt.Sprintf(formatFound, rewards),
		mockPageSize:       4,
		formPage:           "1",
		dbRewardsAll:       mock.RewardsAll{Rwds: mock.Rewards(4)},
	}),
	constructRewardsIndexTest(&rewardsIndexTest{
		purpose:            "no results on page",
		requestMethod:      get,
		responseType:       Core,
		responseStatusCode: http.StatusNotFound,
		responseMessage:    fmt.Sprintf(formatNotFound, page),
		mockPageSize:       0,
		formPage:           "2",
		dbRewardsAll:       mock.RewardsAll{Rwds: mock.Rewards(0)},
	}),
	constructRewardsIndexTest(&rewardsIndexTest{
		purpose:            "missing page",
		requestMethod:      get,
		responseType:       Core,
		responseStatusCode: http.StatusBadRequest,
		responseMessage:    fmt.Sprintf(formatMissing, page),
		formPage:           "",
	}),
	constructRewardsIndexTest(&rewardsIndexTest{
		purpose:            "invalid page",
		requestMethod:      get,
		responseType:       Core,
		responseStatusCode: http.StatusBadRequest,
		responseMessage:    fmt.Sprintf(formatInvalid, page),
		formPage:           "-1",
	}),
	constructRewardsIndexTest(&rewardsIndexTest{
		purpose:            "db error",
		requestMethod:      get,
		responseType:       Core,
		responseStatusCode: http.StatusInternalServerError,
		responseMessage:    friendlyErrorMessage,
		formPage:           mockPage,
		dbRewardsAll:       mock.RewardsAll{Err: mockDbErr},
	}),
}

func constructRewardsIndexTest(testInput *rewardsIndexTest) *test {
	responseResults, _ := json.Marshal(mock.Rewards(testInput.mockPageSize))

	db := &mock.DB{RewardsAllMock: testInput.dbRewardsAll}

	logger := &mock.Logger{}

	return &test{
		purpose: testInput.purpose,
		handle:  RewardsIndex,
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
