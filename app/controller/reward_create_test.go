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
		requestMethod:      get,
		responseType:       Core,
		responseStatusCode: http.StatusMethodNotAllowed,
		responseMessage:    methodNotAllowed,
		form:               &map[string]string{},
	}),
	constructRewardCreateTest(&testInput{
		purpose:            "former error",
		requestMethod:      post,
		responseType:       Core,
		responseStatusCode: http.StatusBadRequest,
		responseMessage:    "former error",
		form:               &map[string]string{},
		former:             &mock.Former{MapMock: mock.Map{Err: mockFormerErr}},
	}),
	constructRewardCreateTest(&testInput{
		purpose:            "missing form title",
		requestMethod:      post,
		responseType:       Core,
		responseStatusCode: http.StatusBadRequest,
		responseMessage:    fmt.Sprintf(formatMissing, title),
		form:               mapWithout(rewardCreateForm(), title),
		former:             &mock.Former{},
	}),
	constructRewardCreateTest(&testInput{
		purpose:            "missing form description",
		requestMethod:      post,
		responseType:       Core,
		responseStatusCode: http.StatusBadRequest,
		responseMessage:    fmt.Sprintf(formatMissing, description),
		form:               mapWithout(rewardCreateForm(), description),
		former:             &mock.Former{},
	}),
	constructRewardCreateTest(&testInput{
		purpose:            "missing form picture_url",
		requestMethod:      post,
		responseType:       Core,
		responseStatusCode: http.StatusBadRequest,
		responseMessage:    fmt.Sprintf(formatMissing, pictureURL),
		form:               mapWithout(rewardCreateForm(), pictureURL),
		former:             &mock.Former{},
	}),
	constructRewardCreateTest(&testInput{
		purpose:            "missing form reward_type_id",
		requestMethod:      post,
		responseType:       Core,
		responseStatusCode: http.StatusBadRequest,
		responseMessage:    fmt.Sprintf(formatMissing, rewardTypeID),
		form:               mapWithout(rewardCreateForm(), rewardTypeID),
		former:             &mock.Former{},
	}),
	constructRewardCreateTest(&testInput{
		purpose:            "reward_type exists db error",
		requestMethod:      post,
		responseType:       Core,
		responseStatusCode: http.StatusInternalServerError,
		responseMessage:    friendlyErrorMessage,
		form:               rewardCreateForm(),
		former:             &mock.Former{},
		db: &mock.DB{
			RewardTypeExistsMock: mock.RewardTypeExists{Err: mockDbErr},
		},
	}),
	constructRewardCreateTest(&testInput{
		purpose:            "reward_type does not exist",
		requestMethod:      post,
		responseType:       Core,
		responseStatusCode: http.StatusNotFound,
		responseMessage:    fmt.Sprintf(formatNotFound, rewardTypeID),
		form:               rewardCreateForm(),
		former:             &mock.Former{},
		db: &mock.DB{
			RewardTypeExistsMock: mock.RewardTypeExists{Bool: false},
		},
	}),
	constructRewardCreateTest(&testInput{
		purpose:            "reward create db error",
		requestMethod:      post,
		responseType:       Core,
		responseStatusCode: http.StatusInternalServerError,
		responseMessage:    friendlyErrorMessage,
		form:               rewardCreateForm(),
		former:             &mock.Former{},
		db: &mock.DB{
			RewardTypeExistsMock: mock.RewardTypeExists{Bool: true},
			RewardCreateMock:     mock.RewardCreate{Err: mockDbErr},
		},
	}),
	constructRewardCreateTest(&testInput{
		purpose:            "reward create ok",
		requestMethod:      post,
		responseType:       Retrieve,
		responseStatusCode: http.StatusOK,
		responseMessage:    fmt.Sprintf(formatCreated, reward),
		form:               rewardCreateForm(),
		former:             &mock.Former{},
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
