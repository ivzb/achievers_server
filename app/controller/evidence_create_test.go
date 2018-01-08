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
		requestMethod:      GET,
		responseType:       Core,
		responseStatusCode: http.StatusMethodNotAllowed,
		responseMessage:    methodNotAllowed,
		form:               &map[string]string{},
	}),
	constructEvidenceCreateTest(&testInput{
		purpose:            "former error",
		requestMethod:      POST,
		responseType:       Core,
		responseStatusCode: http.StatusBadRequest,
		responseMessage:    "former error",
		form:               &map[string]string{},
		former: mock.Form{
			MapMock: mock.Map{Err: mockFormerErr},
		},
	}),
	constructEvidenceCreateTest(&testInput{
		purpose:            "missing form title",
		requestMethod:      POST,
		responseType:       Core,
		responseStatusCode: http.StatusBadRequest,
		responseMessage:    fmt.Sprintf(formatMissing, title),
		form:               mapWithout(evidenceCreateForm(), title),
		former:             mock.Form{},
	}),
	constructEvidenceCreateTest(&testInput{
		purpose:            "missing form picture_url",
		requestMethod:      POST,
		responseType:       Core,
		responseStatusCode: http.StatusBadRequest,
		responseMessage:    fmt.Sprintf(formatMissing, pictureURL),
		form:               mapWithout(evidenceCreateForm(), pictureURL),
		former:             mock.Form{},
	}),
	constructEvidenceCreateTest(&testInput{
		purpose:            "missing form url",
		requestMethod:      POST,
		responseType:       Core,
		responseStatusCode: http.StatusBadRequest,
		responseMessage:    fmt.Sprintf(formatMissing, _url),
		form:               mapWithout(evidenceCreateForm(), _url),
		former:             mock.Form{},
	}),
	constructEvidenceCreateTest(&testInput{
		purpose:            "missing form multimedia_type_id",
		requestMethod:      POST,
		responseType:       Core,
		responseStatusCode: http.StatusBadRequest,
		responseMessage:    fmt.Sprintf(formatMissing, multimediaTypeID),
		form:               mapWithout(evidenceCreateForm(), multimediaTypeID),
		former:             mock.Form{},
	}),
	constructEvidenceCreateTest(&testInput{
		purpose:            "missing form achievement_id",
		requestMethod:      POST,
		responseType:       Core,
		responseStatusCode: http.StatusBadRequest,
		responseMessage:    fmt.Sprintf(formatMissing, achievementID),
		form:               mapWithout(evidenceCreateForm(), achievementID),
		former:             mock.Form{},
	}),
	constructEvidenceCreateTest(&testInput{
		purpose:            "multimediaType exists db error",
		requestMethod:      POST,
		responseType:       Core,
		responseStatusCode: http.StatusInternalServerError,
		responseMessage:    friendlyErrorMessage,
		form:               evidenceCreateForm(),
		former:             mock.Form{},
		db: &mock.DB{
			MultimediaTypeExistsMock: mock.MultimediaTypeExists{Err: mockDbErr},
		},
	}),
	constructEvidenceCreateTest(&testInput{
		purpose:            "multimediaType does not exist",
		requestMethod:      POST,
		responseType:       Core,
		responseStatusCode: http.StatusNotFound,
		responseMessage:    fmt.Sprintf(formatNotFound, multimediaTypeID),
		form:               evidenceCreateForm(),
		former:             mock.Form{},
		db: &mock.DB{
			MultimediaTypeExistsMock: mock.MultimediaTypeExists{Bool: false},
		},
	}),
	constructEvidenceCreateTest(&testInput{
		purpose:            "achievement exists db error",
		requestMethod:      POST,
		responseType:       Core,
		responseStatusCode: http.StatusInternalServerError,
		responseMessage:    friendlyErrorMessage,
		form:               evidenceCreateForm(),
		former:             mock.Form{},
		db: &mock.DB{
			MultimediaTypeExistsMock: mock.MultimediaTypeExists{Bool: true},
			AchievementExistsMock:    mock.AchievementExists{Err: mockDbErr},
		},
	}),
	constructEvidenceCreateTest(&testInput{
		purpose:            "achievement does not exist",
		requestMethod:      POST,
		responseType:       Core,
		responseStatusCode: http.StatusNotFound,
		responseMessage:    fmt.Sprintf(formatNotFound, achievementID),
		form:               evidenceCreateForm(),
		former:             mock.Form{},
		db: &mock.DB{
			MultimediaTypeExistsMock: mock.MultimediaTypeExists{Bool: true},
			AchievementExistsMock:    mock.AchievementExists{Bool: false},
		},
	}),
	constructEvidenceCreateTest(&testInput{
		purpose:            "evidence create db error",
		requestMethod:      POST,
		responseType:       Core,
		responseStatusCode: http.StatusInternalServerError,
		responseMessage:    friendlyErrorMessage,
		form:               evidenceCreateForm(),
		former:             mock.Form{},
		db: &mock.DB{
			MultimediaTypeExistsMock: mock.MultimediaTypeExists{Bool: true},
			AchievementExistsMock:    mock.AchievementExists{Bool: true},
			EvidenceCreateMock:       mock.EvidenceCreate{Err: mockDbErr},
		},
	}),
	constructEvidenceCreateTest(&testInput{
		purpose:            "evidence create ok",
		requestMethod:      POST,
		responseType:       Retrieve,
		responseStatusCode: http.StatusOK,
		responseMessage:    fmt.Sprintf(formatCreated, evidence),
		form:               evidenceCreateForm(),
		former:             mock.Form{},
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
