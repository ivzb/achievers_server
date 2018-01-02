package controller

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/ivzb/achievers_server/app/model/mock"
)

type achievementCreateTest struct {
	purpose             string
	requestMethod       string
	responseType        int
	responseStatusCode  int
	responseMessage     string
	formerErr           error
	formTitle           string
	formDescription     string
	formPictureURL      string
	formInvolvementID   string
	dbInvolvementExists mock.InvolvementExists
	dbAchievementCreate mock.AchievementCreate
}

var achievementCreateTests = []*test{
	constructAchievementCreateTest(&achievementCreateTest{
		purpose:             "invalid request method",
		requestMethod:       get,
		responseType:        Core,
		responseStatusCode:  http.StatusMethodNotAllowed,
		responseMessage:     methodNotAllowed,
		formerErr:           nil,
		formTitle:           "",
		formDescription:     "",
		formPictureURL:      "",
		formInvolvementID:   "",
		dbInvolvementExists: mock.InvolvementExists{},
		dbAchievementCreate: mock.AchievementCreate{},
	}),
	constructAchievementCreateTest(&achievementCreateTest{
		purpose:             "former error",
		requestMethod:       post,
		responseType:        Core,
		responseStatusCode:  http.StatusBadRequest,
		responseMessage:     "former error",
		formerErr:           mockFormerErr,
		formTitle:           "",
		formDescription:     "",
		formPictureURL:      "",
		formInvolvementID:   "",
		dbInvolvementExists: mock.InvolvementExists{},
		dbAchievementCreate: mock.AchievementCreate{},
	}),
	constructAchievementCreateTest(&achievementCreateTest{
		purpose:             "missing form title",
		requestMethod:       post,
		responseType:        Core,
		responseStatusCode:  http.StatusBadRequest,
		responseMessage:     fmt.Sprintf(formatMissing, title),
		formerErr:           nil,
		formTitle:           "",
		formDescription:     mockDescription,
		formPictureURL:      mockPictureURL,
		formInvolvementID:   mockInvolvementID,
		dbInvolvementExists: mock.InvolvementExists{},
		dbAchievementCreate: mock.AchievementCreate{},
	}),
	constructAchievementCreateTest(&achievementCreateTest{
		purpose:             "missing form description",
		requestMethod:       post,
		responseType:        Core,
		responseStatusCode:  http.StatusBadRequest,
		responseMessage:     fmt.Sprintf(formatMissing, description),
		formerErr:           nil,
		formTitle:           mockTitle,
		formDescription:     "",
		formPictureURL:      mockPictureURL,
		formInvolvementID:   mockInvolvementID,
		dbInvolvementExists: mock.InvolvementExists{},
		dbAchievementCreate: mock.AchievementCreate{},
	}),
	constructAchievementCreateTest(&achievementCreateTest{
		purpose:             "missing form picture_url",
		requestMethod:       post,
		responseType:        Core,
		responseStatusCode:  http.StatusBadRequest,
		responseMessage:     fmt.Sprintf(formatMissing, pictureURL),
		formerErr:           nil,
		formTitle:           mockTitle,
		formDescription:     mockDescription,
		formPictureURL:      "",
		formInvolvementID:   mockInvolvementID,
		dbInvolvementExists: mock.InvolvementExists{},
		dbAchievementCreate: mock.AchievementCreate{},
	}),
	constructAchievementCreateTest(&achievementCreateTest{
		purpose:             "missing form involvement_id",
		requestMethod:       post,
		responseType:        Core,
		responseStatusCode:  http.StatusBadRequest,
		responseMessage:     fmt.Sprintf(formatMissing, involvementID),
		formerErr:           nil,
		formTitle:           mockTitle,
		formDescription:     mockDescription,
		formPictureURL:      mockPictureURL,
		formInvolvementID:   "",
		dbInvolvementExists: mock.InvolvementExists{},
		dbAchievementCreate: mock.AchievementCreate{},
	}),
	constructAchievementCreateTest(&achievementCreateTest{
		purpose:             "involvement exists db error",
		requestMethod:       post,
		responseType:        Core,
		responseStatusCode:  http.StatusInternalServerError,
		responseMessage:     friendlyErrorMessage,
		formerErr:           nil,
		formTitle:           mockTitle,
		formDescription:     mockDescription,
		formPictureURL:      mockPictureURL,
		formInvolvementID:   mockInvolvementID,
		dbInvolvementExists: mock.InvolvementExists{Err: mockDbErr},
		dbAchievementCreate: mock.AchievementCreate{},
	}),
	constructAchievementCreateTest(&achievementCreateTest{
		purpose:             "involvement does not exist",
		requestMethod:       post,
		responseType:        Core,
		responseStatusCode:  http.StatusNotFound,
		responseMessage:     fmt.Sprintf(formatNotFound, involvement),
		formerErr:           nil,
		formTitle:           mockTitle,
		formDescription:     mockDescription,
		formPictureURL:      mockPictureURL,
		formInvolvementID:   mockInvolvementID,
		dbInvolvementExists: mock.InvolvementExists{Bool: false},
		dbAchievementCreate: mock.AchievementCreate{},
	}),
	constructAchievementCreateTest(&achievementCreateTest{
		purpose:             "achievement create db error",
		requestMethod:       post,
		responseType:        Core,
		responseStatusCode:  http.StatusInternalServerError,
		responseMessage:     friendlyErrorMessage,
		formerErr:           nil,
		formTitle:           mockTitle,
		formDescription:     mockDescription,
		formPictureURL:      mockPictureURL,
		formInvolvementID:   mockInvolvementID,
		dbInvolvementExists: mock.InvolvementExists{Bool: true},
		dbAchievementCreate: mock.AchievementCreate{Err: mockDbErr},
	}),
	constructAchievementCreateTest(&achievementCreateTest{
		purpose:             "achievement create ok",
		requestMethod:       post,
		responseType:        Retrieve,
		responseStatusCode:  http.StatusOK,
		responseMessage:     fmt.Sprintf(formatCreated, achievement),
		formerErr:           nil,
		formTitle:           mockTitle,
		formDescription:     mockDescription,
		formPictureURL:      mockPictureURL,
		formInvolvementID:   mockInvolvementID,
		dbInvolvementExists: mock.InvolvementExists{Bool: true},
		dbAchievementCreate: mock.AchievementCreate{ID: mockID},
	}),
}

func constructAchievementCreateTest(testInput *achievementCreateTest) *test {
	responseResults, _ := json.Marshal(mockID)

	db := &mock.DB{
		InvolvementExistsMock: testInput.dbInvolvementExists,
		AchievementCreateMock: testInput.dbAchievementCreate,
	}

	logger := &mock.Logger{}

	former := &mock.Former{
		MapMock: mock.Map{Err: testInput.formerErr},
	}

	return &test{
		purpose: testInput.purpose,
		handle:  AchievementCreate,
		request: constructTestRequest(
			testInput.requestMethod,
			constructForm(map[string]string{
				title:         testInput.formTitle,
				description:   testInput.formDescription,
				pictureURL:    testInput.formPictureURL,
				involvementID: testInput.formInvolvementID,
			}),
			constructEnv(db, logger, former),
		),
		response: constructTestResponse(
			testInput.responseType,
			testInput.responseStatusCode,
			testInput.responseMessage,
			responseResults,
		),
	}
}
