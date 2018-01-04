package controller

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/ivzb/achievers_server/app/model/mock"
)

func evidenceCreateForm() *map[string]string {
	return &map[string]string{
		title:            mockTitle,
		pictureURL:       mockPictureURL,
		_url:             mockURL,
		multimediaTypeID: mockMultimediaTypeID,
		achievementID:    mockAchievementID,
	}
}

var evidenceCreateTests = []*test{
	constructEvidenceCreateTest(&testInput{
		purpose:            "invalid request method",
		requestMethod:      get,
		responseType:       Core,
		responseStatusCode: http.StatusMethodNotAllowed,
		responseMessage:    methodNotAllowed,
		form:               &map[string]string{},
	}),
	constructEvidenceCreateTest(&testInput{
		purpose:            "former error",
		requestMethod:      post,
		responseType:       Core,
		responseStatusCode: http.StatusBadRequest,
		responseMessage:    "former error",
		form:               &map[string]string{},
		former: &mock.Former{
			MapMock: mock.Map{Err: mockFormerErr},
		},
	}),
	constructEvidenceCreateTest(&testInput{
		purpose:            "missing form title",
		requestMethod:      post,
		responseType:       Core,
		responseStatusCode: http.StatusBadRequest,
		responseMessage:    fmt.Sprintf(formatMissing, title),
		form:               mapWithout(evidenceCreateForm(), title),
		former:             &mock.Former{},
	}),
	constructEvidenceCreateTest(&testInput{
		purpose:            "missing form picture_url",
		requestMethod:      post,
		responseType:       Core,
		responseStatusCode: http.StatusBadRequest,
		responseMessage:    fmt.Sprintf(formatMissing, pictureURL),
		form:               mapWithout(evidenceCreateForm(), pictureURL),
		former:             &mock.Former{},
	}),
	constructEvidenceCreateTest(&testInput{
		purpose:            "missing form url",
		requestMethod:      post,
		responseType:       Core,
		responseStatusCode: http.StatusBadRequest,
		responseMessage:    fmt.Sprintf(formatMissing, _url),
		form:               mapWithout(evidenceCreateForm(), _url),
		former:             &mock.Former{},
	}),
	constructEvidenceCreateTest(&testInput{
		purpose:            "missing form multimedia_type_id",
		requestMethod:      post,
		responseType:       Core,
		responseStatusCode: http.StatusBadRequest,
		responseMessage:    fmt.Sprintf(formatMissing, multimediaTypeID),
		form:               mapWithout(evidenceCreateForm(), multimediaTypeID),
		former:             &mock.Former{},
	}),
	constructEvidenceCreateTest(&testInput{
		purpose:            "missing form achievement_id",
		requestMethod:      post,
		responseType:       Core,
		responseStatusCode: http.StatusBadRequest,
		responseMessage:    fmt.Sprintf(formatMissing, achievementID),
		form:               mapWithout(evidenceCreateForm(), achievementID),
		former:             &mock.Former{},
	}),
	constructEvidenceCreateTest(&testInput{
		purpose:            "multimediaType exists db error",
		requestMethod:      post,
		responseType:       Core,
		responseStatusCode: http.StatusInternalServerError,
		responseMessage:    friendlyErrorMessage,
		form:               evidenceCreateForm(),
		former:             &mock.Former{},
		db: &mock.DB{
			MultimediaTypeExistsMock: mock.MultimediaTypeExists{Err: mockDbErr},
		},
	}),
	constructEvidenceCreateTest(&testInput{
		purpose:            "multimediaType does not exist",
		requestMethod:      post,
		responseType:       Core,
		responseStatusCode: http.StatusNotFound,
		responseMessage:    fmt.Sprintf(formatNotFound, multimediaTypeID),
		form:               evidenceCreateForm(),
		former:             &mock.Former{},
		db: &mock.DB{
			MultimediaTypeExistsMock: mock.MultimediaTypeExists{Bool: false},
		},
	}),
	constructEvidenceCreateTest(&testInput{
		purpose:            "achievement exists db error",
		requestMethod:      post,
		responseType:       Core,
		responseStatusCode: http.StatusInternalServerError,
		responseMessage:    friendlyErrorMessage,
		form:               evidenceCreateForm(),
		former:             &mock.Former{},
		db: &mock.DB{
			MultimediaTypeExistsMock: mock.MultimediaTypeExists{Bool: true},
			AchievementExistsMock:    mock.AchievementExists{Err: mockDbErr},
		},
	}),
	constructEvidenceCreateTest(&testInput{
		purpose:            "achievement does not exist",
		requestMethod:      post,
		responseType:       Core,
		responseStatusCode: http.StatusNotFound,
		responseMessage:    fmt.Sprintf(formatNotFound, achievementID),
		form:               evidenceCreateForm(),
		former:             &mock.Former{},
		db: &mock.DB{
			MultimediaTypeExistsMock: mock.MultimediaTypeExists{Bool: true},
			AchievementExistsMock:    mock.AchievementExists{Bool: false},
		},
	}),
	constructEvidenceCreateTest(&testInput{
		purpose:            "evidence create db error",
		requestMethod:      post,
		responseType:       Core,
		responseStatusCode: http.StatusInternalServerError,
		responseMessage:    friendlyErrorMessage,
		form:               evidenceCreateForm(),
		former:             &mock.Former{},
		db: &mock.DB{
			MultimediaTypeExistsMock: mock.MultimediaTypeExists{Bool: true},
			AchievementExistsMock:    mock.AchievementExists{Bool: true},
			EvidenceCreateMock:       mock.EvidenceCreate{Err: mockDbErr},
		},
	}),
	constructEvidenceCreateTest(&testInput{
		purpose:            "evidence create ok",
		requestMethod:      post,
		responseType:       Retrieve,
		responseStatusCode: http.StatusOK,
		responseMessage:    fmt.Sprintf(formatCreated, evidence),
		form:               evidenceCreateForm(),
		former:             &mock.Former{},
		db: &mock.DB{
			MultimediaTypeExistsMock: mock.MultimediaTypeExists{Bool: true},
			AchievementExistsMock:    mock.AchievementExists{Bool: true},
			EvidenceCreateMock:       mock.EvidenceCreate{ID: mockID},
		},
	}),
}

func constructEvidenceCreateTest(testInput *testInput) *test {
	responseResults, _ := json.Marshal(mockID)

	return constructTest(EvidenceCreate, testInput, responseResults)
}
