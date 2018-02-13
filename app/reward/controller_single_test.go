package reward

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/ivzb/achievers_server/app/db/mock"
	"github.com/ivzb/achievers_server/app/db/mock/generate"
	"github.com/ivzb/achievers_server/app/shared/consts"
	"github.com/ivzb/achievers_server/app/shared/test"
)

func rewardSingleForm() *map[string]string {
	return &map[string]string{
		consts.ID: test.MockID,
	}
}

var rewardSingleTests = []*test.Test{
	constructRewardSingleTest(&test.TestInput{
		Purpose:            "invalid Request method",
		RequestMethod:      consts.POST,
		ResponseType:       test.Core,
		ResponseStatusCode: http.StatusMethodNotAllowed,
		ResponseMessage:    consts.MethodNotAllowed,
		Form:               &map[string]string{},
	}),
	constructRewardSingleTest(&test.TestInput{
		Purpose:            "missing id",
		RequestMethod:      consts.GET,
		ResponseType:       test.Core,
		ResponseStatusCode: http.StatusBadRequest,
		ResponseMessage:    fmt.Sprintf(consts.FormatMissing, consts.ID),
		Form:               test.MapWithout(rewardSingleForm(), consts.ID),
		DB:                 &mock.DB{},
	}),
	constructRewardSingleTest(&test.TestInput{
		Purpose:            "reward exists db error",
		RequestMethod:      consts.GET,
		ResponseType:       test.Core,
		ResponseStatusCode: http.StatusInternalServerError,
		ResponseMessage:    consts.FriendlyErrorMessage,
		Form:               rewardSingleForm(),
		DB: &mock.DB{
			RewardMock: mock.Reward{
				ExistsMock: mock.RewardExists{Err: test.MockDbErr},
			},
		},
	}),
	constructRewardSingleTest(&test.TestInput{
		Purpose:            "reward does not exist",
		RequestMethod:      consts.GET,
		ResponseType:       test.Core,
		ResponseStatusCode: http.StatusNotFound,
		ResponseMessage:    fmt.Sprintf(consts.FormatNotFound, consts.ID),
		Form:               rewardSingleForm(),
		DB: &mock.DB{
			RewardMock: mock.Reward{
				ExistsMock: mock.RewardExists{Bool: false},
			},
		},
	}),
	constructRewardSingleTest(&test.TestInput{
		Purpose:            "reward single db error",
		RequestMethod:      consts.GET,
		ResponseType:       test.Core,
		ResponseStatusCode: http.StatusInternalServerError,
		ResponseMessage:    consts.FriendlyErrorMessage,
		Form:               rewardSingleForm(),
		DB: &mock.DB{
			RewardMock: mock.Reward{
				ExistsMock: mock.RewardExists{Bool: true},
				SingleMock: mock.RewardSingle{Err: test.MockDbErr},
			},
		},
	}),
	constructRewardSingleTest(&test.TestInput{
		Purpose:            "reward single OK",
		RequestMethod:      consts.GET,
		ResponseType:       test.Retrieve,
		ResponseStatusCode: http.StatusOK,
		ResponseMessage:    fmt.Sprintf(consts.FormatFound, consts.Reward),
		Form:               rewardSingleForm(),
		DB: &mock.DB{
			RewardMock: mock.Reward{
				ExistsMock: mock.RewardExists{Bool: true},
				SingleMock: mock.RewardSingle{Rwd: generate.Reward()},
			},
		},
	}),
}

func constructRewardSingleTest(testInput *test.TestInput) *test.Test {
	ResponseResults, _ := json.Marshal(generate.Reward())

	return test.ConstructTest(Single, testInput, ResponseResults)
}
