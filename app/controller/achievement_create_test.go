package controller

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/ivzb/achievers_server/app/model/mock"
)

func achievementCreateForm() *map[string]string {
	return &map[string]string{
		title:         mockTitle,
		description:   mockDescription,
		pictureURL:    mockPictureURL,
		involvementID: mockInvolvementID,
	}
}

var achievementCreateTests = []*test{
	constructAchievementCreateTest(&testInput{
		purpose:            "invalid request method",
		requestMethod:      GET,
		responseType:       Core,
		responseStatusCode: http.StatusMethodNotAllowed,
		responseMessage:    methodNotAllowed,
		form:               &map[string]string{},
	}),
	constructAchievementCreateTest(&testInput{
		purpose:            "former error",
		requestMethod:      POST,
		responseType:       Core,
		responseStatusCode: http.StatusBadRequest,
		responseMessage:    "content-type of request is incorrect",
		form:               &map[string]string{},
		removeHeaders:      true,
	}),
	constructAchievementCreateTest(&testInput{
		purpose:            "missing form title",
		requestMethod:      POST,
		responseType:       Core,
		responseStatusCode: http.StatusBadRequest,
		responseMessage:    fmt.Sprintf(formatMissing, title),
		form:               mapWithout(achievementCreateForm(), title),
	}),
	constructAchievementCreateTest(&testInput{
		purpose:            "missing form description",
		requestMethod:      POST,
		responseType:       Core,
		responseStatusCode: http.StatusBadRequest,
		responseMessage:    fmt.Sprintf(formatMissing, description),
		form:               mapWithout(achievementCreateForm(), description),
	}),
	constructAchievementCreateTest(&testInput{
		purpose:            "missing form picture_url",
		requestMethod:      POST,
		responseType:       Core,
		responseStatusCode: http.StatusBadRequest,
		responseMessage:    fmt.Sprintf(formatMissing, pictureURL),
		form:               mapWithout(achievementCreateForm(), pictureURL),
	}),
	constructAchievementCreateTest(&testInput{
		purpose:            "missing form involvement_id",
		requestMethod:      POST,
		responseType:       Core,
		responseStatusCode: http.StatusBadRequest,
		responseMessage:    fmt.Sprintf(formatMissing, involvementID),
		form:               mapWithout(achievementCreateForm(), involvementID),
	}),
	constructAchievementCreateTest(&testInput{
		purpose:            "involvement exists db error",
		requestMethod:      POST,
		responseType:       Core,
		responseStatusCode: http.StatusInternalServerError,
		responseMessage:    friendlyErrorMessage,
		form:               achievementCreateForm(),
		db: &mock.DB{
			InvolvementExistsMock: mock.InvolvementExists{Err: mockDbErr},
		},
	}),
	constructAchievementCreateTest(&testInput{
		purpose:            "involvement does not exist",
		requestMethod:      POST,
		responseType:       Core,
		responseStatusCode: http.StatusNotFound,
		responseMessage:    fmt.Sprintf(formatNotFound, involvementID),
		form:               achievementCreateForm(),
		db: &mock.DB{
			InvolvementExistsMock: mock.InvolvementExists{Bool: false},
		},
	}),
	constructAchievementCreateTest(&testInput{
		purpose:            "achievement create db error",
		requestMethod:      POST,
		responseType:       Core,
		responseStatusCode: http.StatusInternalServerError,
		responseMessage:    friendlyErrorMessage,
		form:               achievementCreateForm(),
		db: &mock.DB{
			InvolvementExistsMock: mock.InvolvementExists{Bool: true},
			AchievementCreateMock: mock.AchievementCreate{Err: mockDbErr},
		},
	}),
	constructAchievementCreateTest(&testInput{
		purpose:            "achievement create ok",
		requestMethod:      POST,
		responseType:       Retrieve,
		responseStatusCode: http.StatusCreated,
		responseMessage:    fmt.Sprintf(formatCreated, achievement),
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
