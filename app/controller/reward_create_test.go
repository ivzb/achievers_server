package controller

import (
	"encoding/json"
	"fmt"
	"net/http"

	dMock "github.com/ivzb/achievers_server/app/db/mock"
	"github.com/ivzb/achievers_server/app/shared/consts"
)

func rewardCreateForm() *map[string]string {
	return &map[string]string{
		consts.Title:        mockTitle,
		consts.Description:  mockDescription,
		consts.PictureURL:   mockPictureURL,
		consts.RewardTypeID: mockRewardTypeID,
	}
}

var rewardCreateTests = []*test{
	constructRewardCreateTest(&testInput{
		purpose:            "invalid request method",
		requestMethod:      consts.GET,
		responseType:       Core,
		responseStatusCode: http.StatusMethodNotAllowed,
		responseMessage:    consts.MethodNotAllowed,
		form:               &map[string]string{},
	}),
	constructRewardCreateTest(&testInput{
		purpose:            "former error",
		requestMethod:      consts.POST,
		responseType:       Core,
		responseStatusCode: http.StatusBadRequest,
		responseMessage:    "content-type of request is incorrect",
		form:               &map[string]string{},
		removeHeaders:      true,
	}),
	constructRewardCreateTest(&testInput{
		purpose:            "blank form consts.Title",
		requestMethod:      consts.POST,
		responseType:       Core,
		responseStatusCode: http.StatusBadRequest,
		responseMessage:    fmt.Sprintf(consts.FormatBlank, consts.Title),
		form:               mapWithout(rewardCreateForm(), consts.Title),
	}),
	constructRewardCreateTest(&testInput{
		purpose:            "blank form consts.Description",
		requestMethod:      consts.POST,
		responseType:       Core,
		responseStatusCode: http.StatusBadRequest,
		responseMessage:    fmt.Sprintf(consts.FormatBlank, consts.Description),
		form:               mapWithout(rewardCreateForm(), consts.Description),
	}),
	constructRewardCreateTest(&testInput{
		purpose:            "blank form picture_url",
		requestMethod:      consts.POST,
		responseType:       Core,
		responseStatusCode: http.StatusBadRequest,
		responseMessage:    fmt.Sprintf(consts.FormatBlank, consts.PictureURL),
		form:               mapWithout(rewardCreateForm(), consts.PictureURL),
	}),
	constructRewardCreateTest(&testInput{
		purpose:            "blank form reward_type_id",
		requestMethod:      consts.POST,
		responseType:       Core,
		responseStatusCode: http.StatusBadRequest,
		responseMessage:    fmt.Sprintf(consts.FormatValidID, consts.RewardTypeID),
		form:               mapWithout(rewardCreateForm(), consts.RewardTypeID),
	}),
	constructRewardCreateTest(&testInput{
		purpose:            "reward_type exists db error",
		requestMethod:      consts.POST,
		responseType:       Core,
		responseStatusCode: http.StatusInternalServerError,
		responseMessage:    consts.FriendlyErrorMessage,
		form:               rewardCreateForm(),
		db: &dMock.DB{
			RewardTypeMock: dMock.RewardType{
				ExistsMock: dMock.RewardTypeExists{Err: mockDbErr},
			},
		},
	}),
	constructRewardCreateTest(&testInput{
		purpose:            "reward_type does not exist",
		requestMethod:      consts.POST,
		responseType:       Core,
		responseStatusCode: http.StatusNotFound,
		responseMessage:    fmt.Sprintf(consts.FormatNotFound, consts.RewardTypeID),
		form:               rewardCreateForm(),
		db: &dMock.DB{
			RewardTypeMock: dMock.RewardType{
				ExistsMock: dMock.RewardTypeExists{Bool: false},
			},
		},
	}),
	constructRewardCreateTest(&testInput{
		purpose:            "reward create db error",
		requestMethod:      consts.POST,
		responseType:       Core,
		responseStatusCode: http.StatusInternalServerError,
		responseMessage:    consts.FriendlyErrorMessage,
		form:               rewardCreateForm(),
		db: &dMock.DB{
			RewardTypeMock: dMock.RewardType{
				ExistsMock: dMock.RewardTypeExists{Bool: true},
			},
			RewardMock: dMock.Reward{
				CreateMock: dMock.RewardCreate{Err: mockDbErr},
			},
		},
	}),
	constructRewardCreateTest(&testInput{
		purpose:            "reward create ok",
		requestMethod:      consts.POST,
		responseType:       Retrieve,
		responseStatusCode: http.StatusCreated,
		responseMessage:    fmt.Sprintf(consts.FormatCreated, consts.Reward),
		form:               rewardCreateForm(),
		db: &dMock.DB{
			RewardTypeMock: dMock.RewardType{
				ExistsMock: dMock.RewardTypeExists{Bool: true},
			},
			RewardMock: dMock.Reward{
				CreateMock: dMock.RewardCreate{ID: mockID},
			},
		},
	}),
}

func constructRewardCreateTest(testInput *testInput) *test {
	responseResults, _ := json.Marshal(mockID)
	return constructTest(RewardCreate, testInput, responseResults)
}