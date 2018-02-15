package controller

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/ivzb/achievers_server/app/db/mock"
	"github.com/ivzb/achievers_server/app/db/mock/generate"
	"github.com/ivzb/achievers_server/app/shared/consts"
)

var rewardsLatestArgs = []string{"9"}

var rewardsLatestTests = []*test{
	constructRewardsLatestTest(&testInput{
		purpose:            "invalid request method",
		requestMethod:      consts.POST,
		responseType:       Core,
		responseStatusCode: http.StatusMethodNotAllowed,
		responseMessage:    consts.MethodNotAllowed,
		form:               &map[string]string{},
		args:               rewardsLatestArgs,
	}),
	constructRewardsLatestTest(&testInput{
		purpose:            "id RewardsLastID db error",
		requestMethod:      consts.GET,
		responseType:       Core,
		responseStatusCode: http.StatusInternalServerError,
		responseMessage:    consts.FriendlyErrorMessage,
		db: &mock.DB{
			RewardMock: mock.Reward{
				LastIDMock: mock.RewardsLastID{Err: mockDbErr},
			},
		},
		args: rewardsLatestArgs,
	}),
	constructRewardsLatestTest(&testInput{
		purpose:            "RewardsAfterMock db error",
		requestMethod:      consts.GET,
		responseType:       Core,
		responseStatusCode: http.StatusInternalServerError,
		responseMessage:    consts.FriendlyErrorMessage,
		db: &mock.DB{
			RewardMock: mock.Reward{
				LastIDMock: mock.RewardsLastID{ID: mockID},
				AfterMock:  mock.RewardsAfter{Err: mockDbErr},
			},
		},
		args: rewardsLatestArgs,
	}),
	constructRewardsLatestTest(&testInput{
		purpose:            "no results",
		requestMethod:      consts.GET,
		responseType:       Core,
		responseStatusCode: http.StatusOK,
		responseMessage:    fmt.Sprintf(consts.FormatFound, consts.Rewards),
		db: &mock.DB{
			RewardMock: mock.Reward{
				LastIDMock: mock.RewardsLastID{ID: mockID},
				AfterMock:  mock.RewardsAfter{Rwds: generate.Rewards(0)},
			},
		},
		args: []string{"0"},
	}),
	constructRewardsLatestTest(&testInput{
		purpose:            "4 results",
		requestMethod:      consts.GET,
		responseType:       Retrieve,
		responseStatusCode: http.StatusOK,
		responseMessage:    fmt.Sprintf(consts.FormatFound, consts.Rewards),
		db: &mock.DB{
			RewardMock: mock.Reward{
				LastIDMock: mock.RewardsLastID{ID: mockID},
				AfterMock:  mock.RewardsAfter{Rwds: generate.Rewards(4)},
			},
		},
		args: []string{"4"},
	}),
	constructRewardsLatestTest(&testInput{
		purpose:            "9 results",
		requestMethod:      consts.GET,
		responseType:       Retrieve,
		responseStatusCode: http.StatusOK,
		responseMessage:    fmt.Sprintf(consts.FormatFound, consts.Rewards),
		db: &mock.DB{
			RewardMock: mock.Reward{
				LastIDMock: mock.RewardsLastID{ID: mockID},
				AfterMock:  mock.RewardsAfter{Rwds: generate.Rewards(9)},
			},
		},
		args: []string{"9"},
	}),
}

func constructRewardsLatestTest(testInput *testInput) *test {
	rewardsSize, err := strconv.Atoi(testInput.args[0])

	var responseResults []byte

	if err == nil {
		responseResults, _ = json.Marshal(generate.Rewards(rewardsSize))
	}

	return constructTest(RewardsLast, testInput, responseResults)
}