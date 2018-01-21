package controller

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/ivzb/achievers_server/app/db/mock"
	"github.com/ivzb/achievers_server/app/shared/consts"
)

func achievementCreateForm() *map[string]string {
	return &map[string]string{
		consts.Title:         mockTitle,
		consts.Description:   mockDescription,
		consts.PictureURL:    mockPictureURL,
		consts.InvolvementID: mockInvolvementID,
	}
}

var achievementCreateTests = []*test{
	constructAchievementCreateTest(&testInput{
		purpose:            "invalid request method",
		requestMethod:      consts.GET,
		responseType:       Core,
		responseStatusCode: http.StatusMethodNotAllowed,
		responseMessage:    consts.MethodNotAllowed,
		form:               &map[string]string{},
	}),
	constructAchievementCreateTest(&testInput{
		purpose:            "former error",
		requestMethod:      consts.POST,
		responseType:       Core,
		responseStatusCode: http.StatusBadRequest,
		responseMessage:    "content-type of request is incorrect",
		form:               &map[string]string{},
		removeHeaders:      true,
	}),
	constructAchievementCreateTest(&testInput{
		purpose:            "missing form consts.Title",
		requestMethod:      consts.POST,
		responseType:       Core,
		responseStatusCode: http.StatusBadRequest,
		responseMessage:    fmt.Sprintf(consts.FormatMissing, consts.Title),
		form:               mapWithout(achievementCreateForm(), consts.Title),
	}),
	constructAchievementCreateTest(&testInput{
		purpose:            "missing form consts.Description",
		requestMethod:      consts.POST,
		responseType:       Core,
		responseStatusCode: http.StatusBadRequest,
		responseMessage:    fmt.Sprintf(consts.FormatMissing, consts.Description),
		form:               mapWithout(achievementCreateForm(), consts.Description),
	}),
	constructAchievementCreateTest(&testInput{
		purpose:            "missing form picture_url",
		requestMethod:      consts.POST,
		responseType:       Core,
		responseStatusCode: http.StatusBadRequest,
		responseMessage:    fmt.Sprintf(consts.FormatMissing, consts.PictureURL),
		form:               mapWithout(achievementCreateForm(), consts.PictureURL),
	}),
	constructAchievementCreateTest(&testInput{
		purpose:            "missing form involvement_id",
		requestMethod:      consts.POST,
		responseType:       Core,
		responseStatusCode: http.StatusBadRequest,
		responseMessage:    fmt.Sprintf(consts.FormatMissing, consts.InvolvementID),
		form:               mapWithout(achievementCreateForm(), consts.InvolvementID),
	}),
	constructAchievementCreateTest(&testInput{
		purpose:            "involvement exists db error",
		requestMethod:      consts.POST,
		responseType:       Core,
		responseStatusCode: http.StatusInternalServerError,
		responseMessage:    consts.FriendlyErrorMessage,
		form:               achievementCreateForm(),
		db: &mock.DB{
			InvolvementExistsMock: mock.InvolvementExists{Err: mockDbErr},
		},
	}),
	constructAchievementCreateTest(&testInput{
		purpose:            "involvement does not exist",
		requestMethod:      consts.POST,
		responseType:       Core,
		responseStatusCode: http.StatusNotFound,
		responseMessage:    fmt.Sprintf(consts.FormatNotFound, consts.InvolvementID),
		form:               achievementCreateForm(),
		db: &mock.DB{
			InvolvementExistsMock: mock.InvolvementExists{Bool: false},
		},
	}),
	constructAchievementCreateTest(&testInput{
		purpose:            "achievement create db error",
		requestMethod:      consts.POST,
		responseType:       Core,
		responseStatusCode: http.StatusInternalServerError,
		responseMessage:    consts.FriendlyErrorMessage,
		form:               achievementCreateForm(),
		db: &mock.DB{
			InvolvementExistsMock: mock.InvolvementExists{Bool: true},
			AchievementCreateMock: mock.AchievementCreate{Err: mockDbErr},
		},
	}),
	constructAchievementCreateTest(&testInput{
		purpose:            "achievement create ok",
		requestMethod:      consts.POST,
		responseType:       Retrieve,
		responseStatusCode: http.StatusCreated,
		responseMessage:    fmt.Sprintf(consts.FormatCreated, consts.Achievement),
		form:               achievementCreateForm(),
		db: &mock.DB{
			InvolvementExistsMock: mock.InvolvementExists{Bool: true},
			AchievementCreateMock: mock.AchievementCreate{ID: mockID},
		},
	}),
}

func constructAchievementCreateTest(testInput *testInput) *test {
	responseResults, _ := json.Marshal(mockID)
	return constructTest(AchievementCreate, testInput, responseResults)
}
