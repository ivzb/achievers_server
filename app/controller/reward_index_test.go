package controller

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/ivzb/achievers_server/app/model/mock"
)

func rewardsIndexForm() *map[string]string {
	return &map[string]string{
		page: mockPage,
	}
}

var rewardsIndexArgs = []string{"9"}

var rewardsIndexTests = []*test{
	constructRewardsIndexTest(&testInput{
		purpose:            "invalid request method",
		requestMethod:      post,
		responseType:       Core,
		responseStatusCode: http.StatusMethodNotAllowed,
		responseMessage:    methodNotAllowed,
		form:               &map[string]string{},
		args:               rewardsIndexArgs,
	}),
	constructRewardsIndexTest(&testInput{
		purpose:            "missing page",
		requestMethod:      get,
		responseType:       Core,
		responseStatusCode: http.StatusBadRequest,
		responseMessage:    fmt.Sprintf(formatMissing, page),
		form:               mapWithout(rewardsIndexForm(), page),
		args:               rewardsIndexArgs,
	}),
	constructRewardsIndexTest(&testInput{
		purpose:            "invalid page",
		requestMethod:      get,
		responseType:       Core,
		responseStatusCode: http.StatusBadRequest,
		responseMessage:    fmt.Sprintf(formatInvalid, page),
		form: &map[string]string{
			page: "-1",
		},
		args: rewardsIndexArgs,
	}),
	constructRewardsIndexTest(&testInput{
		purpose:            "db error",
		requestMethod:      get,
		responseType:       Core,
		responseStatusCode: http.StatusInternalServerError,
		responseMessage:    friendlyErrorMessage,
		form:               rewardsIndexForm(),
		db: &mock.DB{
			RewardsAllMock: mock.RewardsAll{Err: mockDbErr},
		},
		args: rewardsIndexArgs,
	}),
	constructRewardsIndexTest(&testInput{
		purpose:            "no results on page",
		requestMethod:      get,
		responseType:       Core,
		responseStatusCode: http.StatusNotFound,
		responseMessage:    fmt.Sprintf(formatNotFound, page),
		form:               rewardsIndexForm(),
		db: &mock.DB{
			RewardsAllMock: mock.RewardsAll{Rwds: mock.Rewards(0)},
		},
		args: []string{"0"},
	}),
	constructRewardsIndexTest(&testInput{
		purpose:            "4 results on page",
		requestMethod:      get,
		responseType:       Retrieve,
		responseStatusCode: http.StatusOK,
		responseMessage:    fmt.Sprintf(formatFound, rewards),
		form:               rewardsIndexForm(),
		db: &mock.DB{
			RewardsAllMock: mock.RewardsAll{Rwds: mock.Rewards(4)},
		},
		args: []string{"4"},
	}),
	constructRewardsIndexTest(&testInput{
		purpose:            "9 results on page",
		requestMethod:      get,
		responseType:       Retrieve,
		responseStatusCode: http.StatusOK,
		responseMessage:    fmt.Sprintf(formatFound, rewards),
		form:               rewardsIndexForm(),
		db: &mock.DB{
			RewardsAllMock: mock.RewardsAll{Rwds: mock.Rewards(9)},
		},
		args: []string{"9"},
	}),
}

func constructRewardsIndexTest(testInput *testInput) *test {
	rewardsSize, err := strconv.Atoi(testInput.args[0])

	var responseResults []byte

	if err == nil {
		responseResults, _ = json.Marshal(mock.Rewards(rewardsSize))
	}

	return constructTest(RewardsIndex, testInput, responseResults)
}
