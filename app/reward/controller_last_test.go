package reward

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/ivzb/achievers_server/app/db/mock"
	"github.com/ivzb/achievers_server/app/db/mock/generate"
	"github.com/ivzb/achievers_server/app/shared/consts"
	"github.com/ivzb/achievers_server/app/shared/test"
)

var rewardsLatestArgs = []string{"9"}

var rewardsLatestTests = []*test.Test{
	constructRewardsLatestTest(&test.TestInput{
		Purpose:            "invalid Request method",
		RequestMethod:      consts.POST,
		ResponseType:       test.Core,
		ResponseStatusCode: http.StatusMethodNotAllowed,
		ResponseMessage:    consts.MethodNotAllowed,
		Form:               &map[string]string{},
		Args:               rewardsLatestArgs,
	}),
	constructRewardsLatestTest(&test.TestInput{
		Purpose:            "id RewardsLastID db error",
		RequestMethod:      consts.GET,
		ResponseType:       test.Core,
		ResponseStatusCode: http.StatusInternalServerError,
		ResponseMessage:    consts.FriendlyErrorMessage,
		DB: &mock.DB{
			RewardMock: mock.Reward{
				LastIDMock: mock.RewardsLastID{Err: test.MockDbErr},
			},
		},
		Args: rewardsLatestArgs,
	}),
	constructRewardsLatestTest(&test.TestInput{
		Purpose:            "RewardsAfterMock db error",
		RequestMethod:      consts.GET,
		ResponseType:       test.Core,
		ResponseStatusCode: http.StatusInternalServerError,
		ResponseMessage:    consts.FriendlyErrorMessage,
		DB: &mock.DB{
			RewardMock: mock.Reward{
				LastIDMock: mock.RewardsLastID{ID: test.MockID},
				AfterMock:  mock.RewardsAfter{Err: test.MockDbErr},
			},
		},
		Args: rewardsLatestArgs,
	}),
	constructRewardsLatestTest(&test.TestInput{
		Purpose:            "no results",
		RequestMethod:      consts.GET,
		ResponseType:       test.Core,
		ResponseStatusCode: http.StatusOK,
		ResponseMessage:    fmt.Sprintf(consts.FormatFound, consts.Rewards),
		DB: &mock.DB{
			RewardMock: mock.Reward{
				LastIDMock: mock.RewardsLastID{ID: test.MockID},
				AfterMock:  mock.RewardsAfter{Rwds: generate.Rewards(0)},
			},
		},
		Args: []string{"0"},
	}),
	constructRewardsLatestTest(&test.TestInput{
		Purpose:            "4 results",
		RequestMethod:      consts.GET,
		ResponseType:       test.Retrieve,
		ResponseStatusCode: http.StatusOK,
		ResponseMessage:    fmt.Sprintf(consts.FormatFound, consts.Rewards),
		DB: &mock.DB{
			RewardMock: mock.Reward{
				LastIDMock: mock.RewardsLastID{ID: test.MockID},
				AfterMock:  mock.RewardsAfter{Rwds: generate.Rewards(4)},
			},
		},
		Args: []string{"4"},
	}),
	constructRewardsLatestTest(&test.TestInput{
		Purpose:            "9 results",
		RequestMethod:      consts.GET,
		ResponseType:       test.Retrieve,
		ResponseStatusCode: http.StatusOK,
		ResponseMessage:    fmt.Sprintf(consts.FormatFound, consts.Rewards),
		DB: &mock.DB{
			RewardMock: mock.Reward{
				LastIDMock: mock.RewardsLastID{ID: test.MockID},
				AfterMock:  mock.RewardsAfter{Rwds: generate.Rewards(9)},
			},
		},
		Args: []string{"9"},
	}),
}

func constructRewardsLatestTest(testInput *test.TestInput) *test.Test {
	rewardsSize, err := strconv.Atoi(testInput.Args[0])

	var ResponseResults []byte

	if err == nil {
		ResponseResults, _ = json.Marshal(generate.Rewards(rewardsSize))
	}

	return test.ConstructTest(Last, testInput, ResponseResults)
}
