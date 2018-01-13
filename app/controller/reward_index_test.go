package controller

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/ivzb/achievers_server/app/model/mock"
	"github.com/ivzb/achievers_server/app/shared/consts"
)

func rewardsIndexForm() *map[string]string {
	return &map[string]string{
		consts.Page: mockPage,
	}
}

var rewardsIndexArgs = []string{"9"}

var rewardsIndexTests = []*test{
	constructRewardsIndexTest(&testInput{
		purpose:            "invalid request method",
		requestMethod:      consts.POST,
		responseType:       Core,
		responseStatusCode: http.StatusMethodNotAllowed,
		responseMessage:    consts.MethodNotAllowed,
		form:               &map[string]string{},
		args:               rewardsIndexArgs,
	}),
	constructRewardsIndexTest(&testInput{
		purpose:            "missing consts.Page",
		requestMethod:      consts.GET,
		responseType:       Core,
		responseStatusCode: http.StatusBadRequest,
		responseMessage:    fmt.Sprintf(consts.FormatMissing, consts.Page),
		form:               mapWithout(rewardsIndexForm(), consts.Page),
		args:               rewardsIndexArgs,
	}),
	constructRewardsIndexTest(&testInput{
		purpose:            "invalid consts.Page",
		requestMethod:      consts.GET,
		responseType:       Core,
		responseStatusCode: http.StatusBadRequest,
		responseMessage:    fmt.Sprintf(consts.FormatInvalid, consts.Page),
		form: &map[string]string{
			consts.Page: "-1",
		},
		args: rewardsIndexArgs,
	}),
	constructRewardsIndexTest(&testInput{
		purpose:            "db error",
		requestMethod:      consts.GET,
		responseType:       Core,
		responseStatusCode: http.StatusInternalServerError,
		responseMessage:    consts.FriendlyErrorMessage,
		form:               rewardsIndexForm(),
		db: &mock.DB{
			RewardsAllMock: mock.RewardsAll{Err: mockDbErr},
		},
		args: rewardsIndexArgs,
	}),
	constructRewardsIndexTest(&testInput{
		purpose:            "no results on consts.Page",
		requestMethod:      consts.GET,
		responseType:       Core,
		responseStatusCode: http.StatusNotFound,
		responseMessage:    fmt.Sprintf(consts.FormatNotFound, consts.Page),
		form:               rewardsIndexForm(),
		db: &mock.DB{
			RewardsAllMock: mock.RewardsAll{Rwds: mock.Rewards(0)},
		},
		args: []string{"0"},
	}),
	constructRewardsIndexTest(&testInput{
		purpose:            "4 results on consts.Page",
		requestMethod:      consts.GET,
		responseType:       Retrieve,
		responseStatusCode: http.StatusOK,
		responseMessage:    fmt.Sprintf(consts.FormatFound, consts.Rewards),
		form:               rewardsIndexForm(),
		db: &mock.DB{
			RewardsAllMock: mock.RewardsAll{Rwds: mock.Rewards(4)},
		},
		args: []string{"4"},
	}),
	constructRewardsIndexTest(&testInput{
		purpose:            "9 results on consts.Page",
		requestMethod:      consts.GET,
		responseType:       Retrieve,
		responseStatusCode: http.StatusOK,
		responseMessage:    fmt.Sprintf(consts.FormatFound, consts.Rewards),
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
