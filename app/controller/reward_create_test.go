package controller

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/ivzb/achievers_server/app/model/mock"
)

func rewardCreateForm() *map[string]string {
	return &map[string]string{
		title:        mockTitle,
		description:  mockDescription,
		pictureURL:   mockPictureURL,
		rewardTypeID: mockRewardTypeID,
	}
}

var rewardCreateTests = []*test{
	constructRewardCreateTest(&testInput{
		purpose:            "invalid request method",
		requestMethod:      GET,
		responseType:       Core,
		responseStatusCode: http.StatusMethodNotAllowed,
		responseMessage:    methodNotAllowed,
		form:               &map[string]string{},
	}),
	constructRewardCreateTest(&testInput{
		purpose:            "former error",
		requestMethod:      POST,
		responseType:       Core,
		responseStatusCode: http.StatusBadRequest,
		responseMessage:    "former error",
		form:               &map[string]string{},
		former:             mock.Form{MapMock: mock.Map{Err: mockFormerErr}},
	}),
	constructRewardCreateTest(&testInput{
		purpose:            "missing form title",
		requestMethod:      POST,
		responseType:       Core,
		responseStatusCode: http.StatusBadRequest,
		responseMessage:    fmt.Sprintf(formatMissing, title),
		form:               mapWithout(rewardCreateForm(), title),
		former:             mock.Form{},
	}),
	constructRewardCreateTest(&testInput{
		purpose:            "missing form description",
		requestMethod:      POST,
		responseType:       Core,
		responseStatusCode: http.StatusBadRequest,
		responseMessage:    fmt.Sprintf(formatMissing, description),
		form:               mapWithout(rewardCreateForm(), description),
		former:             mock.Form{},
	}),
	constructRewardCreateTest(&testInput{
		purpose:            "missing form picture_url",
		requestMethod:      POST,
		responseType:       Core,
		responseStatusCode: http.StatusBadRequest,
		responseMessage:    fmt.Sprintf(formatMissing, pictureURL),
		form:               mapWithout(rewardCreateForm(), pictureURL),
		former:             mock.Form{},
	}),
	constructRewardCreateTest(&testInput{
		purpose:            "missing form reward_type_id",
		requestMethod:      POST,
		responseType:       Core,
		responseStatusCode: http.StatusBadRequest,
		responseMessage:    fmt.Sprintf(formatMissing, rewardTypeID),
		form:               mapWithout(rewardCreateForm(), rewardTypeID),
		former:             mock.Form{},
	}),
	constructRewardCreateTest(&testInput{
		purpose:            "reward_type exists db error",
		requestMethod:      POST,
		responseType:       Core,
		responseStatusCode: http.StatusInternalServerError,
		responseMessage:    friendlyErrorMessage,
		form:               rewardCreateForm(),
		former:             mock.Form{},
		db: &mock.DB{
			RewardTypeExistsMock: mock.RewardTypeExists{Err: mockDbErr},
		},
	}),
	constructRewardCreateTest(&testInput{
		purpose:            "reward_type does not exist",
		requestMethod:      POST,
		responseType:       Core,
		responseStatusCode: http.StatusNotFound,
		responseMessage:    fmt.Sprintf(formatNotFound, rewardTypeID),
		form:               rewardCreateForm(),
		former:             mock.Form{},
		db: &mock.DB{
			RewardTypeExistsMock: mock.RewardTypeExists{Bool: false},
		},
	}),
	constructRewardCreateTest(&testInput{
		purpose:            "reward create db error",
		requestMethod:      POST,
		responseType:       Core,
		responseStatusCode: http.StatusInternalServerError,
		responseMessage:    friendlyErrorMessage,
		form:               rewardCreateForm(),
		former:             mock.Form{},
		db: &mock.DB{
			RewardTypeExistsMock: mock.RewardTypeExists{Bool: true},
			RewardCreateMock:     mock.RewardCreate{Err: mockDbErr},
		},
	}),
	constructRewardCreateTest(&testInput{
		purpose:            "reward create ok",
		requestMethod:      POST,
		responseType:       Retrieve,
		responseStatusCode: http.StatusOK,
		responseMessage:    fmt.Sprintf(formatCreated, reward),
		form:               rewardCreateForm(),
		former:             mock.Form{},
		db: &mock.DB{
			RewardTypeExistsMock: mock.RewardTypeExists{Bool: true},
			RewardCreateMock:     mock.RewardCreate{ID: mockID},
		},
	}),
}

func constructRewardCreateTest(testInput *testInput) *test {
	responseResults, _ := json.Marshal(mockID)
	return constructTest(RewardCreate, testInput, responseResults)
}
