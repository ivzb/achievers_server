package controller

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/ivzb/achievers_server/app/model/mock"
)

func rewardSingleForm() *map[string]string {
	return &map[string]string{
		id: mockID,
	}
}

var rewardSingleTests = []*test{
	constructRewardSingleTest(&testInput{
		purpose:            "invalid request method",
		requestMethod:      POST,
		responseType:       Core,
		responseStatusCode: http.StatusMethodNotAllowed,
		responseMessage:    methodNotAllowed,
		form:               &map[string]string{},
	}),
	constructRewardSingleTest(&testInput{
		purpose:            "missing id",
		requestMethod:      GET,
		responseType:       Core,
		responseStatusCode: http.StatusBadRequest,
		responseMessage:    fmt.Sprintf(formatMissing, id),
		form:               mapWithout(rewardSingleForm(), id),
	}),
	constructRewardSingleTest(&testInput{
		purpose:            "reward exists db error",
		requestMethod:      GET,
		responseType:       Core,
		responseStatusCode: http.StatusInternalServerError,
		responseMessage:    friendlyErrorMessage,
		form:               rewardSingleForm(),
		db: &mock.DB{
			RewardExistsMock: mock.RewardExists{Err: mockDbErr},
		},
	}),
	constructRewardSingleTest(&testInput{
		purpose:            "reward does not exist",
		requestMethod:      GET,
		responseType:       Core,
		responseStatusCode: http.StatusNotFound,
		responseMessage:    fmt.Sprintf(formatNotFound, reward),
		form:               rewardSingleForm(),
		db: &mock.DB{
			RewardExistsMock: mock.RewardExists{Bool: false},
		},
	}),
	constructRewardSingleTest(&testInput{
		purpose:            "reward single db error",
		requestMethod:      GET,
		responseType:       Core,
		responseStatusCode: http.StatusInternalServerError,
		responseMessage:    friendlyErrorMessage,
		form:               rewardSingleForm(),
		db: &mock.DB{
			RewardExistsMock: mock.RewardExists{Bool: true},
			RewardSingleMock: mock.RewardSingle{Err: mockDbErr},
		},
	}),
	constructRewardSingleTest(&testInput{
		purpose:            "reward single OK",
		requestMethod:      GET,
		responseType:       Retrieve,
		responseStatusCode: http.StatusOK,
		responseMessage:    fmt.Sprintf(formatFound, reward),
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
