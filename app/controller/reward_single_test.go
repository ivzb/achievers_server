package controller

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/ivzb/achievers_server/app/model/mock"
)

type rewardSingleTest struct {
	purpose            string
	requestMethod      string
	responseType       int
	responseStatusCode int
	responseMessage    string
	formID             string
	dbRewardExists     mock.RewardExists
	dbRewardSingle     mock.RewardSingle
}

var rewardSingleTests = []*test{
	constructRewardSingleTest(&rewardSingleTest{
		purpose:            "invalid request method",
		requestMethod:      post,
		responseType:       Core,
		responseStatusCode: http.StatusMethodNotAllowed,
		responseMessage:    methodNotAllowed,
		formID:             "",
		dbRewardExists:     mock.RewardExists{},
		dbRewardSingle:     mock.RewardSingle{},
	}),
	constructRewardSingleTest(&rewardSingleTest{
		purpose:            "missing id",
		requestMethod:      get,
		responseType:       Core,
		responseStatusCode: http.StatusBadRequest,
		responseMessage:    fmt.Sprintf(formatMissing, id),
		formID:             "",
		dbRewardExists:     mock.RewardExists{},
		dbRewardSingle:     mock.RewardSingle{},
	}),
	constructRewardSingleTest(&rewardSingleTest{
		purpose:            "reward exists db error",
		requestMethod:      get,
		responseType:       Core,
		responseStatusCode: http.StatusInternalServerError,
		responseMessage:    friendlyErrorMessage,
		formID:             mockID,
		dbRewardExists:     mock.RewardExists{Err: mockDbErr},
		dbRewardSingle:     mock.RewardSingle{},
	}),
	constructRewardSingleTest(&rewardSingleTest{
		purpose:            "reward does not exist",
		requestMethod:      get,
		responseType:       Core,
		responseStatusCode: http.StatusNotFound,
		responseMessage:    fmt.Sprintf(formatNotFound, reward),
		formID:             mockID,
		dbRewardExists:     mock.RewardExists{Bool: false},
		dbRewardSingle:     mock.RewardSingle{},
	}),
	constructRewardSingleTest(&rewardSingleTest{
		purpose:            "reward single db error",
		requestMethod:      get,
		responseType:       Core,
		responseStatusCode: http.StatusInternalServerError,
		responseMessage:    friendlyErrorMessage,
		formID:             mockID,
		dbRewardExists:     mock.RewardExists{Bool: true},
		dbRewardSingle:     mock.RewardSingle{Err: mockDbErr},
	}),
	constructRewardSingleTest(&rewardSingleTest{
		purpose:            "reward single OK",
		requestMethod:      get,
		responseType:       Retrieve,
		responseStatusCode: http.StatusOK,
		responseMessage:    fmt.Sprintf(formatFound, reward),
		formID:             mockID,
		dbRewardExists:     mock.RewardExists{Bool: true},
		dbRewardSingle:     mock.RewardSingle{Rwd: mock.Reward()},
	}),
}

func constructRewardSingleTest(testInput *rewardSingleTest) *test {
	responseResults, _ := json.Marshal(mock.Reward())

	db := &mock.DB{
		RewardExistsMock: testInput.dbRewardExists,
		RewardSingleMock: testInput.dbRewardSingle,
	}

	logger := &mock.Logger{}

	return &test{
		purpose: testInput.purpose,
		handle:  RewardSingle,
		request: constructTestRequest(
			testInput.requestMethod,
			constructForm(map[string]string{
				id: testInput.formID,
			}),
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
