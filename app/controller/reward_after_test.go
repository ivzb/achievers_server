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

func rewardsAfterForm() *map[string]string {
	return &map[string]string{
		consts.ID: mockID,
	}
}

var rewardsAfterArgs = []string{"9"}

var rewardsAfterTests = []*test{
	constructRewardsAfterTest(&testInput{
		purpose:            "invalid request method",
		requestMethod:      consts.POST,
		responseType:       Core,
		responseStatusCode: http.StatusMethodNotAllowed,
		responseMessage:    consts.MethodNotAllowed,
		form:               &map[string]string{},
		args:               rewardsAfterArgs,
	}),
	constructRewardsAfterTest(&testInput{
		purpose:            "missing id",
		requestMethod:      consts.GET,
		responseType:       Core,
		responseStatusCode: http.StatusBadRequest,
		responseMessage:    fmt.Sprintf(consts.FormatMissing, consts.ID),
		form:               mapWithout(rewardsAfterForm(), consts.ID),
		db: &mock.DB{
			RewardMock: mock.Reward{
				ExistsMock: mock.RewardExists{Err: mockDbErr},
			},
		},
		args: rewardsAfterArgs,
	}),
	constructRewardsAfterTest(&testInput{
		purpose:            "id RewardExists db error",
		requestMethod:      consts.GET,
		responseType:       Core,
		responseStatusCode: http.StatusInternalServerError,
		responseMessage:    consts.FriendlyErrorMessage,
		form:               rewardsAfterForm(),
		db: &mock.DB{
			RewardMock: mock.Reward{
				ExistsMock: mock.RewardExists{Err: mockDbErr},
			},
		},
		args: rewardsAfterArgs,
	}),
	constructRewardsAfterTest(&testInput{
		purpose:            "id RewardExists does not exist",
		requestMethod:      consts.GET,
		responseType:       Core,
		responseStatusCode: http.StatusNotFound,
		responseMessage:    fmt.Sprintf(consts.FormatNotFound, consts.ID),
		form:               rewardsAfterForm(),
		db: &mock.DB{
			RewardMock: mock.Reward{
				ExistsMock: mock.RewardExists{Bool: false},
			},
		},
		args: rewardsAfterArgs,
	}),
	constructRewardsAfterTest(&testInput{
		purpose:            "RewardsAfterMock db error",
		requestMethod:      consts.GET,
		responseType:       Core,
		responseStatusCode: http.StatusInternalServerError,
		responseMessage:    consts.FriendlyErrorMessage,
		form:               rewardsAfterForm(),
		db: &mock.DB{
			RewardMock: mock.Reward{
				ExistsMock: mock.RewardExists{Bool: true},
				AfterMock:  mock.RewardsAfter{Err: mockDbErr},
			},
		},
		args: rewardsAfterArgs,
	}),
	constructRewardsAfterTest(&testInput{
		purpose:            "no results",
		requestMethod:      consts.GET,
		responseType:       Core,
		responseStatusCode: http.StatusOK,
		responseMessage:    fmt.Sprintf(consts.FormatFound, consts.Rewards),
		form:               rewardsAfterForm(),
		db: &mock.DB{
			RewardMock: mock.Reward{
				ExistsMock: mock.RewardExists{Bool: true},
				AfterMock:  mock.RewardsAfter{Rwds: generate.Rewards(0)},
			},
		},
		args: []string{"0"},
	}),
	constructRewardsAfterTest(&testInput{
		purpose:            "4 results",
		requestMethod:      consts.GET,
		responseType:       Retrieve,
		responseStatusCode: http.StatusOK,
		responseMessage:    fmt.Sprintf(consts.FormatFound, consts.Rewards),
		form:               rewardsAfterForm(),
		db: &mock.DB{
			RewardMock: mock.Reward{
				ExistsMock: mock.RewardExists{Bool: true},
				AfterMock:  mock.RewardsAfter{Rwds: generate.Rewards(4)},
			},
		},
		args: []string{"4"},
	}),
	constructRewardsAfterTest(&testInput{
		purpose:            "9 results",
		requestMethod:      consts.GET,
		responseType:       Retrieve,
		responseStatusCode: http.StatusOK,
		responseMessage:    fmt.Sprintf(consts.FormatFound, consts.Rewards),
		form:               rewardsAfterForm(),
		db: &mock.DB{
			RewardMock: mock.Reward{
				ExistsMock: mock.RewardExists{Bool: true},
				AfterMock:  mock.RewardsAfter{Rwds: generate.Rewards(9)},
			},
		},
		args: []string{"9"},
	}),
}

func constructRewardsAfterTest(testInput *testInput) *test {
	rewardsSize, err := strconv.Atoi(testInput.args[0])

	var responseResults []byte

	if err == nil {
		responseResults, _ = json.Marshal(generate.Rewards(rewardsSize))
	}

	return constructTest(RewardsAfter, testInput, responseResults)
}