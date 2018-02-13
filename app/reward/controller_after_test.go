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

func rewardsAfterForm() *map[string]string {
	return &map[string]string{
		consts.ID: test.MockID,
	}
}

var rewardsAfterArgs = []string{"9"}

var rewardsAfterTests = []*test.Test{
	constructRewardsAfterTest(&test.TestInput{
		Purpose:            "invalid request method",
		RequestMethod:      consts.POST,
		ResponseType:       test.Core,
		ResponseStatusCode: http.StatusMethodNotAllowed,
		ResponseMessage:    consts.MethodNotAllowed,
		Form:               &map[string]string{},
		Args:               rewardsAfterArgs,
	}),
	constructRewardsAfterTest(&test.TestInput{
		Purpose:            "missing id",
		RequestMethod:      consts.GET,
		ResponseType:       test.Core,
		ResponseStatusCode: http.StatusBadRequest,
		ResponseMessage:    fmt.Sprintf(consts.FormatMissing, consts.ID),
		Form:               test.MapWithout(rewardsAfterForm(), consts.ID),
		DB: &mock.DB{
			RewardMock: mock.Reward{
				ExistsMock: mock.RewardExists{Err: test.MockDbErr},
			},
		},
		Args: rewardsAfterArgs,
	}),
	constructRewardsAfterTest(&test.TestInput{
		Purpose:            "id RewardExists db error",
		RequestMethod:      consts.GET,
		ResponseType:       test.Core,
		ResponseStatusCode: http.StatusInternalServerError,
		ResponseMessage:    consts.FriendlyErrorMessage,
		Form:               rewardsAfterForm(),
		DB: &mock.DB{
			RewardMock: mock.Reward{
				ExistsMock: mock.RewardExists{Err: test.MockDbErr},
			},
		},
		Args: rewardsAfterArgs,
	}),
	constructRewardsAfterTest(&test.TestInput{
		Purpose:            "id RewardExists does not exist",
		RequestMethod:      consts.GET,
		ResponseType:       test.Core,
		ResponseStatusCode: http.StatusNotFound,
		ResponseMessage:    fmt.Sprintf(consts.FormatNotFound, consts.ID),
		Form:               rewardsAfterForm(),
		DB: &mock.DB{
			RewardMock: mock.Reward{
				ExistsMock: mock.RewardExists{Bool: false},
			},
		},
		Args: rewardsAfterArgs,
	}),
	constructRewardsAfterTest(&test.TestInput{
		Purpose:            "RewardsAfterMock db error",
		RequestMethod:      consts.GET,
		ResponseType:       test.Core,
		ResponseStatusCode: http.StatusInternalServerError,
		ResponseMessage:    consts.FriendlyErrorMessage,
		Form:               rewardsAfterForm(),
		DB: &mock.DB{
			RewardMock: mock.Reward{
				ExistsMock: mock.RewardExists{Bool: true},
				AfterMock:  mock.RewardsAfter{Err: test.MockDbErr},
			},
		},
		Args: rewardsAfterArgs,
	}),
	constructRewardsAfterTest(&test.TestInput{
		Purpose:            "no results",
		RequestMethod:      consts.GET,
		ResponseType:       test.Core,
		ResponseStatusCode: http.StatusOK,
		ResponseMessage:    fmt.Sprintf(consts.FormatFound, consts.Rewards),
		Form:               rewardsAfterForm(),
		DB: &mock.DB{
			RewardMock: mock.Reward{
				ExistsMock: mock.RewardExists{Bool: true},
				AfterMock:  mock.RewardsAfter{Rwds: generate.Rewards(0)},
			},
		},
		Args: []string{"0"},
	}),
	constructRewardsAfterTest(&test.TestInput{
		Purpose:            "4 results",
		RequestMethod:      consts.GET,
		ResponseType:       test.Retrieve,
		ResponseStatusCode: http.StatusOK,
		ResponseMessage:    fmt.Sprintf(consts.FormatFound, consts.Rewards),
		Form:               rewardsAfterForm(),
		DB: &mock.DB{
			RewardMock: mock.Reward{
				ExistsMock: mock.RewardExists{Bool: true},
				AfterMock:  mock.RewardsAfter{Rwds: generate.Rewards(4)},
			},
		},
		Args: []string{"4"},
	}),
	constructRewardsAfterTest(&test.TestInput{
		Purpose:            "9 results",
		RequestMethod:      consts.GET,
		ResponseType:       test.Retrieve,
		ResponseStatusCode: http.StatusOK,
		ResponseMessage:    fmt.Sprintf(consts.FormatFound, consts.Rewards),
		Form:               rewardsAfterForm(),
		DB: &mock.DB{
			RewardMock: mock.Reward{
				ExistsMock: mock.RewardExists{Bool: true},
				AfterMock:  mock.RewardsAfter{Rwds: generate.Rewards(9)},
			},
		},
		Args: []string{"9"},
	}),
}

func constructRewardsAfterTest(testInput *test.TestInput) *test.Test {
	rewardsSize, err := strconv.Atoi(testInput.Args[0])

	var responseResults []byte

	if err == nil {
		responseResults, _ = json.Marshal(generate.Rewards(rewardsSize))
	}

	return test.ConstructTest(After, testInput, responseResults)
}
