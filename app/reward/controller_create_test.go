package reward

import (
	"encoding/json"
	"fmt"
	"net/http"

	dMock "github.com/ivzb/achievers_server/app/db/mock"
	"github.com/ivzb/achievers_server/app/shared/consts"
	"github.com/ivzb/achievers_server/app/shared/test"
)

func rewardCreateForm() *map[string]string {
	return &map[string]string{
		consts.Title:        test.MockTitle,
		consts.Description:  test.MockDescription,
		consts.PictureURL:   test.MockPictureURL,
		consts.RewardTypeID: test.MockRewardTypeID,
	}
}

var rewardCreateTests = []*test.Test{
	constructRewardCreateTest(&test.TestInput{
		Purpose:            "invalid Request method",
		RequestMethod:      consts.GET,
		ResponseType:       test.Core,
		ResponseStatusCode: http.StatusMethodNotAllowed,
		ResponseMessage:    consts.MethodNotAllowed,
		Form:               &map[string]string{},
	}),
	constructRewardCreateTest(&test.TestInput{
		Purpose:            "former error",
		RequestMethod:      consts.POST,
		ResponseType:       test.Core,
		ResponseStatusCode: http.StatusBadRequest,
		ResponseMessage:    "content-type of request is incorrect",
		Form:               &map[string]string{},
		RemoveHeaders:      true,
	}),
	constructRewardCreateTest(&test.TestInput{
		Purpose:            "blank form consts.Title",
		RequestMethod:      consts.POST,
		ResponseType:       test.Core,
		ResponseStatusCode: http.StatusBadRequest,
		ResponseMessage:    fmt.Sprintf(consts.FormatBlank, consts.Title),
		Form:               test.MapWithout(rewardCreateForm(), consts.Title),
	}),
	constructRewardCreateTest(&test.TestInput{
		Purpose:            "blank form consts.Description",
		RequestMethod:      consts.POST,
		ResponseType:       test.Core,
		ResponseStatusCode: http.StatusBadRequest,
		ResponseMessage:    fmt.Sprintf(consts.FormatBlank, consts.Description),
		Form:               test.MapWithout(rewardCreateForm(), consts.Description),
	}),
	constructRewardCreateTest(&test.TestInput{
		Purpose:            "blank form picture_url",
		RequestMethod:      consts.POST,
		ResponseType:       test.Core,
		ResponseStatusCode: http.StatusBadRequest,
		ResponseMessage:    fmt.Sprintf(consts.FormatBlank, consts.PictureURL),
		Form:               test.MapWithout(rewardCreateForm(), consts.PictureURL),
	}),
	constructRewardCreateTest(&test.TestInput{
		Purpose:            "blank form reward_type_id",
		RequestMethod:      consts.POST,
		ResponseType:       test.Core,
		ResponseStatusCode: http.StatusBadRequest,
		ResponseMessage:    fmt.Sprintf(consts.FormatValidID, consts.RewardTypeID),
		Form:               test.MapWithout(rewardCreateForm(), consts.RewardTypeID),
	}),
	constructRewardCreateTest(&test.TestInput{
		Purpose:            "reward_type exists db error",
		RequestMethod:      consts.POST,
		ResponseType:       test.Core,
		ResponseStatusCode: http.StatusInternalServerError,
		ResponseMessage:    consts.FriendlyErrorMessage,
		Form:               rewardCreateForm(),
		DB: &dMock.DB{
			RewardTypeMock: dMock.RewardType{
				ExistsMock: dMock.RewardTypeExists{Err: test.MockDbErr},
			},
		},
	}),
	constructRewardCreateTest(&test.TestInput{
		Purpose:            "reward_type does not exist",
		RequestMethod:      consts.POST,
		ResponseType:       test.Core,
		ResponseStatusCode: http.StatusNotFound,
		ResponseMessage:    fmt.Sprintf(consts.FormatNotFound, consts.RewardTypeID),
		Form:               rewardCreateForm(),
		DB: &dMock.DB{
			RewardTypeMock: dMock.RewardType{
				ExistsMock: dMock.RewardTypeExists{Bool: false},
			},
		},
	}),
	constructRewardCreateTest(&test.TestInput{
		Purpose:            "reward create db error",
		RequestMethod:      consts.POST,
		ResponseType:       test.Core,
		ResponseStatusCode: http.StatusInternalServerError,
		ResponseMessage:    consts.FriendlyErrorMessage,
		Form:               rewardCreateForm(),
		DB: &dMock.DB{
			RewardTypeMock: dMock.RewardType{
				ExistsMock: dMock.RewardTypeExists{Bool: true},
			},
			RewardMock: dMock.Reward{
				CreateMock: dMock.RewardCreate{Err: test.MockDbErr},
			},
		},
	}),
	constructRewardCreateTest(&test.TestInput{
		Purpose:            "reward create ok",
		RequestMethod:      consts.POST,
		ResponseType:       test.Retrieve,
		ResponseStatusCode: http.StatusCreated,
		ResponseMessage:    fmt.Sprintf(consts.FormatCreated, consts.Reward),
		Form:               rewardCreateForm(),
		DB: &dMock.DB{
			RewardTypeMock: dMock.RewardType{
				ExistsMock: dMock.RewardTypeExists{Bool: true},
			},
			RewardMock: dMock.Reward{
				CreateMock: dMock.RewardCreate{ID: test.MockID},
			},
		},
	}),
}

func constructRewardCreateTest(testInput *test.TestInput) *test.Test {
	responseResults, _ := json.Marshal(test.MockID)
	return test.ConstructTest(Create, testInput, responseResults)
}
