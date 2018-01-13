package controller

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/ivzb/achievers_server/app/model/mock"
	"github.com/ivzb/achievers_server/app/shared/consts"
)

func rewardSingleForm() *map[string]string {
	return &map[string]string{
		consts.ID: mockID,
	}
}

var rewardSingleTests = []*test{
	constructRewardSingleTest(&testInput{
		purpose:            "invalid request method",
		requestMethod:      consts.POST,
		responseType:       Core,
		responseStatusCode: http.StatusMethodNotAllowed,
		responseMessage:    consts.MethodNotAllowed,
		form:               &map[string]string{},
	}),
	constructRewardSingleTest(&testInput{
		purpose:            "missing id",
		requestMethod:      consts.GET,
		responseType:       Core,
		responseStatusCode: http.StatusBadRequest,
		responseMessage:    fmt.Sprintf(consts.FormatMissing, consts.ID),
		form:               mapWithout(rewardSingleForm(), consts.ID),
	}),
	constructRewardSingleTest(&testInput{
		purpose:            "reward exists db error",
		requestMethod:      consts.GET,
		responseType:       Core,
		responseStatusCode: http.StatusInternalServerError,
		responseMessage:    consts.FriendlyErrorMessage,
		form:               rewardSingleForm(),
		db: &mock.DB{
			RewardExistsMock: mock.RewardExists{Err: mockDbErr},
		},
	}),
	constructRewardSingleTest(&testInput{
		purpose:            "reward does not exist",
		requestMethod:      consts.GET,
		responseType:       Core,
		responseStatusCode: http.StatusNotFound,
		responseMessage:    fmt.Sprintf(consts.FormatNotFound, consts.Reward),
		form:               rewardSingleForm(),
		db: &mock.DB{
			RewardExistsMock: mock.RewardExists{Bool: false},
		},
	}),
	constructRewardSingleTest(&testInput{
		purpose:            "reward single db error",
		requestMethod:      consts.GET,
		responseType:       Core,
		responseStatusCode: http.StatusInternalServerError,
		responseMessage:    consts.FriendlyErrorMessage,
		form:               rewardSingleForm(),
		db: &mock.DB{
			RewardExistsMock: mock.RewardExists{Bool: true},
			RewardSingleMock: mock.RewardSingle{Err: mockDbErr},
		},
	}),
	constructRewardSingleTest(&testInput{
		purpose:            "reward single OK",
		requestMethod:      consts.GET,
		responseType:       Retrieve,
		responseStatusCode: http.StatusOK,
		responseMessage:    fmt.Sprintf(consts.FormatFound, consts.Reward),
		form:               rewardSingleForm(),
		db: &mock.DB{
			RewardExistsMock: mock.RewardExists{Bool: true},
			RewardSingleMock: mock.RewardSingle{Rwd: mock.Reward()},
		},
	}),
}

func constructRewardSingleTest(testInput *testInput) *test {
	responseResults, _ := json.Marshal(mock.Reward())

	return constructTest(RewardSingle, testInput, responseResults)
}
