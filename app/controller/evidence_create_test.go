package controller

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/ivzb/achievers_server/app/model/mock"
)

type evidenceCreateTest struct {
	purpose                string
	requestMethod          string
	responseType           int
	responseStatusCode     int
	responseMessage        string
	formerErr              error
	formDescription        string
	formPreviewURL         string
	formURL                string
	formMultimediaTypeID   string
	formAchievementID      string
	dbMultimediaTypeExists mock.MultimediaTypeExists
	dbAchievementExists    mock.AchievementExists
	dbEvidenceCreate       mock.EvidenceCreate
}

var evidenceCreateTests = []*test{
	constructEvidenceCreateTest(&evidenceCreateTest{
		purpose:                "invalid request method",
		requestMethod:          get,
		responseType:           Core,
		responseStatusCode:     http.StatusMethodNotAllowed,
		responseMessage:        methodNotAllowed,
		formerErr:              nil,
		formDescription:        "",
		formPreviewURL:         "",
		formURL:                "",
		formMultimediaTypeID:   "",
		formAchievementID:      "",
		dbMultimediaTypeExists: mock.MultimediaTypeExists{},
		dbAchievementExists:    mock.AchievementExists{},
		dbEvidenceCreate:       mock.EvidenceCreate{},
	}),
	constructEvidenceCreateTest(&evidenceCreateTest{
		purpose:                "former error",
		requestMethod:          post,
		responseType:           Core,
		responseStatusCode:     http.StatusBadRequest,
		responseMessage:        "former error",
		formerErr:              mockFormerErr,
		formDescription:        "",
		formPreviewURL:         "",
		formURL:                "",
		formMultimediaTypeID:   "",
		formAchievementID:      "",
		dbMultimediaTypeExists: mock.MultimediaTypeExists{},
		dbAchievementExists:    mock.AchievementExists{},
		dbEvidenceCreate:       mock.EvidenceCreate{},
	}),
	constructEvidenceCreateTest(&evidenceCreateTest{
		purpose:                "missing form description",
		requestMethod:          post,
		responseType:           Core,
		responseStatusCode:     http.StatusBadRequest,
		responseMessage:        fmt.Sprintf(formatMissing, description),
		formerErr:              nil,
		formDescription:        "",
		formPreviewURL:         mockPreviewURL,
		formURL:                mockURL,
		formMultimediaTypeID:   mockMultimediaTypeID,
		formAchievementID:      mockAchievementID,
		dbMultimediaTypeExists: mock.MultimediaTypeExists{},
		dbAchievementExists:    mock.AchievementExists{},
		dbEvidenceCreate:       mock.EvidenceCreate{},
	}),
	constructEvidenceCreateTest(&evidenceCreateTest{
		purpose:                "missing form preview_url",
		requestMethod:          post,
		responseType:           Core,
		responseStatusCode:     http.StatusBadRequest,
		responseMessage:        fmt.Sprintf(formatMissing, previewURL),
		formerErr:              nil,
		formDescription:        mockDescription,
		formPreviewURL:         "",
		formURL:                mockURL,
		formMultimediaTypeID:   mockMultimediaTypeID,
		formAchievementID:      mockAchievementID,
		dbMultimediaTypeExists: mock.MultimediaTypeExists{},
		dbAchievementExists:    mock.AchievementExists{},
		dbEvidenceCreate:       mock.EvidenceCreate{},
	}),
	constructEvidenceCreateTest(&evidenceCreateTest{
		purpose:                "missing form url",
		requestMethod:          post,
		responseType:           Core,
		responseStatusCode:     http.StatusBadRequest,
		responseMessage:        fmt.Sprintf(formatMissing, _url),
		formerErr:              nil,
		formDescription:        mockDescription,
		formPreviewURL:         mockPreviewURL,
		formURL:                "",
		formMultimediaTypeID:   mockMultimediaTypeID,
		formAchievementID:      mockAchievementID,
		dbMultimediaTypeExists: mock.MultimediaTypeExists{},
		dbAchievementExists:    mock.AchievementExists{},
		dbEvidenceCreate:       mock.EvidenceCreate{},
	}),
	constructEvidenceCreateTest(&evidenceCreateTest{
		purpose:                "missing form multimedia_type_id",
		requestMethod:          post,
		responseType:           Core,
		responseStatusCode:     http.StatusBadRequest,
		responseMessage:        fmt.Sprintf(formatMissing, multimediaTypeID),
		formerErr:              nil,
		formDescription:        mockDescription,
		formPreviewURL:         mockPreviewURL,
		formURL:                mockURL,
		formMultimediaTypeID:   "",
		formAchievementID:      mockAchievementID,
		dbMultimediaTypeExists: mock.MultimediaTypeExists{},
		dbAchievementExists:    mock.AchievementExists{},
		dbEvidenceCreate:       mock.EvidenceCreate{},
	}),
	constructEvidenceCreateTest(&evidenceCreateTest{
		purpose:                "missing form achievement_id",
		requestMethod:          post,
		responseType:           Core,
		responseStatusCode:     http.StatusBadRequest,
		responseMessage:        fmt.Sprintf(formatMissing, achievementID),
		formerErr:              nil,
		formDescription:        mockDescription,
		formPreviewURL:         mockPreviewURL,
		formURL:                mockURL,
		formMultimediaTypeID:   mockMultimediaTypeID,
		formAchievementID:      "",
		dbMultimediaTypeExists: mock.MultimediaTypeExists{},
		dbAchievementExists:    mock.AchievementExists{},
		dbEvidenceCreate:       mock.EvidenceCreate{},
	}),
	constructEvidenceCreateTest(&evidenceCreateTest{
		purpose:                "multimediaType exists db error",
		requestMethod:          post,
		responseType:           Core,
		responseStatusCode:     http.StatusInternalServerError,
		responseMessage:        friendlyErrorMessage,
		formerErr:              nil,
		formDescription:        mockDescription,
		formPreviewURL:         mockPreviewURL,
		formURL:                mockURL,
		formMultimediaTypeID:   mockMultimediaTypeID,
		formAchievementID:      mockAchievementID,
		dbMultimediaTypeExists: mock.MultimediaTypeExists{Err: mockDbErr},
		dbAchievementExists:    mock.AchievementExists{},
		dbEvidenceCreate:       mock.EvidenceCreate{},
	}),
	constructEvidenceCreateTest(&evidenceCreateTest{
		purpose:                "multimediaType does not exist",
		requestMethod:          post,
		responseType:           Core,
		responseStatusCode:     http.StatusNotFound,
		responseMessage:        fmt.Sprintf(formatNotFound, multimediaType),
		formerErr:              nil,
		formDescription:        mockDescription,
		formPreviewURL:         mockPreviewURL,
		formURL:                mockURL,
		formMultimediaTypeID:   mockMultimediaTypeID,
		formAchievementID:      mockAchievementID,
		dbMultimediaTypeExists: mock.MultimediaTypeExists{Bool: false},
		dbAchievementExists:    mock.AchievementExists{},
		dbEvidenceCreate:       mock.EvidenceCreate{},
	}),
	constructEvidenceCreateTest(&evidenceCreateTest{
		purpose:                "achievement exists db error",
		requestMethod:          post,
		responseType:           Core,
		responseStatusCode:     http.StatusInternalServerError,
		responseMessage:        friendlyErrorMessage,
		formerErr:              nil,
		formDescription:        mockDescription,
		formPreviewURL:         mockPreviewURL,
		formURL:                mockURL,
		formMultimediaTypeID:   mockMultimediaTypeID,
		formAchievementID:      mockAchievementID,
		dbMultimediaTypeExists: mock.MultimediaTypeExists{Bool: true},
		dbAchievementExists:    mock.AchievementExists{Err: mockDbErr},
		dbEvidenceCreate:       mock.EvidenceCreate{},
	}),
	constructEvidenceCreateTest(&evidenceCreateTest{
		purpose:                "achievement does not exist",
		requestMethod:          post,
		responseType:           Core,
		responseStatusCode:     http.StatusNotFound,
		responseMessage:        fmt.Sprintf(formatNotFound, achievement),
		formerErr:              nil,
		formDescription:        mockDescription,
		formPreviewURL:         mockPreviewURL,
		formURL:                mockURL,
		formMultimediaTypeID:   mockMultimediaTypeID,
		formAchievementID:      mockAchievementID,
		dbMultimediaTypeExists: mock.MultimediaTypeExists{Bool: true},
		dbAchievementExists:    mock.AchievementExists{Bool: false},
		dbEvidenceCreate:       mock.EvidenceCreate{},
	}),
	constructEvidenceCreateTest(&evidenceCreateTest{
		purpose:                "evidence create db error",
		requestMethod:          post,
		responseType:           Core,
		responseStatusCode:     http.StatusInternalServerError,
		responseMessage:        friendlyErrorMessage,
		formerErr:              nil,
		formDescription:        mockDescription,
		formPreviewURL:         mockPreviewURL,
		formURL:                mockURL,
		formMultimediaTypeID:   mockMultimediaTypeID,
		formAchievementID:      mockAchievementID,
		dbMultimediaTypeExists: mock.MultimediaTypeExists{Bool: true},
		dbAchievementExists:    mock.AchievementExists{Bool: true},
		dbEvidenceCreate:       mock.EvidenceCreate{Err: mockDbErr},
	}),
	constructEvidenceCreateTest(&evidenceCreateTest{
		purpose:                "evidence create ok",
		requestMethod:          post,
		responseType:           Retrieve,
		responseStatusCode:     http.StatusOK,
		responseMessage:        fmt.Sprintf(formatCreated, evidence),
		formerErr:              nil,
		formDescription:        mockDescription,
		formPreviewURL:         mockPreviewURL,
		formURL:                mockURL,
		formMultimediaTypeID:   mockMultimediaTypeID,
		formAchievementID:      mockAchievementID,
		dbMultimediaTypeExists: mock.MultimediaTypeExists{Bool: true},
		dbAchievementExists:    mock.AchievementExists{Bool: true},
		dbEvidenceCreate:       mock.EvidenceCreate{ID: mockID},
	}),
}

func constructEvidenceCreateTest(testInput *evidenceCreateTest) *test {
	responseResults, _ := json.Marshal(mockID)

	db := &mock.DB{
		MultimediaTypeExistsMock: testInput.dbMultimediaTypeExists,
		AchievementExistsMock:    testInput.dbAchievementExists,
		EvidenceCreateMock:       testInput.dbEvidenceCreate,
	}

	logger := &mock.Logger{}

	former := &mock.Former{
		MapMock: mock.Map{Err: testInput.formerErr},
	}

	return &test{
		purpose: testInput.purpose,
		handle:  EvidenceCreate,
		request: constructTestRequest(
			testInput.requestMethod,
			constructForm(map[string]string{
				description:      testInput.formDescription,
				previewURL:       testInput.formPreviewURL,
				_url:             testInput.formURL,
				multimediaTypeID: testInput.formMultimediaTypeID,
				achievementID:    testInput.formAchievementID,
			}),
			constructEnv(db, logger, former, nil),
		),
		response: constructTestResponse(
			testInput.responseType,
			testInput.responseStatusCode,
			testInput.responseMessage,
			responseResults,
		),
	}
}
