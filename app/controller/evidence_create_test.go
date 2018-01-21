package controller

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/ivzb/achievers_server/app/db/mock"
	"github.com/ivzb/achievers_server/app/shared/consts"
)

func evidenceCreateForm() *map[string]string {
	return &map[string]string{
		consts.Title:            mockTitle,
		consts.PictureURL:       mockPictureURL,
		consts.URL:              mockURL,
		consts.MultimediaTypeID: mockMultimediaTypeID,
		consts.AchievementID:    mockAchievementID,
	}
}

var evidenceCreateTests = []*test{
	constructEvidenceCreateTest(&testInput{
		purpose:            "invalid request method",
		requestMethod:      consts.GET,
		responseType:       Core,
		responseStatusCode: http.StatusMethodNotAllowed,
		responseMessage:    consts.MethodNotAllowed,
		form:               &map[string]string{},
	}),
	constructEvidenceCreateTest(&testInput{
		purpose:            "former error",
		requestMethod:      consts.POST,
		responseType:       Core,
		responseStatusCode: http.StatusBadRequest,
		responseMessage:    "content-type of request is incorrect",
		form:               &map[string]string{},
		removeHeaders:      true,
	}),
	constructEvidenceCreateTest(&testInput{
		purpose:            "missing form consts.Title",
		requestMethod:      consts.POST,
		responseType:       Core,
		responseStatusCode: http.StatusBadRequest,
		responseMessage:    fmt.Sprintf(consts.FormatMissing, consts.Title),
		form:               mapWithout(evidenceCreateForm(), consts.Title),
	}),
	constructEvidenceCreateTest(&testInput{
		purpose:            "missing form pictureconsts.URL",
		requestMethod:      consts.POST,
		responseType:       Core,
		responseStatusCode: http.StatusBadRequest,
		responseMessage:    fmt.Sprintf(consts.FormatMissing, consts.PictureURL),
		form:               mapWithout(evidenceCreateForm(), consts.PictureURL),
	}),
	constructEvidenceCreateTest(&testInput{
		purpose:            "missing form url",
		requestMethod:      consts.POST,
		responseType:       Core,
		responseStatusCode: http.StatusBadRequest,
		responseMessage:    fmt.Sprintf(consts.FormatMissing, consts.URL),
		form:               mapWithout(evidenceCreateForm(), consts.URL),
	}),
	constructEvidenceCreateTest(&testInput{
		purpose:            "missing form multimedia_type_id",
		requestMethod:      consts.POST,
		responseType:       Core,
		responseStatusCode: http.StatusBadRequest,
		responseMessage:    fmt.Sprintf(consts.FormatMissing, consts.MultimediaTypeID),
		form:               mapWithout(evidenceCreateForm(), consts.MultimediaTypeID),
	}),
	constructEvidenceCreateTest(&testInput{
		purpose:            "missing form achievement_id",
		requestMethod:      consts.POST,
		responseType:       Core,
		responseStatusCode: http.StatusBadRequest,
		responseMessage:    fmt.Sprintf(consts.FormatMissing, consts.AchievementID),
		form:               mapWithout(evidenceCreateForm(), consts.AchievementID),
	}),
	constructEvidenceCreateTest(&testInput{
		purpose:            "multimediaType exists db error",
		requestMethod:      consts.POST,
		responseType:       Core,
		responseStatusCode: http.StatusInternalServerError,
		responseMessage:    consts.FriendlyErrorMessage,
		form:               evidenceCreateForm(),
		db: &mock.DB{
			MultimediaTypeExistsMock: mock.MultimediaTypeExists{Err: mockDbErr},
		},
	}),
	constructEvidenceCreateTest(&testInput{
		purpose:            "multimediaType does not exist",
		requestMethod:      consts.POST,
		responseType:       Core,
		responseStatusCode: http.StatusNotFound,
		responseMessage:    fmt.Sprintf(consts.FormatNotFound, consts.MultimediaTypeID),
		form:               evidenceCreateForm(),
		db: &mock.DB{
			MultimediaTypeExistsMock: mock.MultimediaTypeExists{Bool: false},
		},
	}),
	constructEvidenceCreateTest(&testInput{
		purpose:            "achievement exists db error",
		requestMethod:      consts.POST,
		responseType:       Core,
		responseStatusCode: http.StatusInternalServerError,
		responseMessage:    consts.FriendlyErrorMessage,
		form:               evidenceCreateForm(),
		db: &mock.DB{
			MultimediaTypeExistsMock: mock.MultimediaTypeExists{Bool: true},
			AchievementExistsMock:    mock.AchievementExists{Err: mockDbErr},
		},
	}),
	constructEvidenceCreateTest(&testInput{
		purpose:            "achievement does not exist",
		requestMethod:      consts.POST,
		responseType:       Core,
		responseStatusCode: http.StatusNotFound,
		responseMessage:    fmt.Sprintf(consts.FormatNotFound, consts.AchievementID),
		form:               evidenceCreateForm(),
		db: &mock.DB{
			MultimediaTypeExistsMock: mock.MultimediaTypeExists{Bool: true},
			AchievementExistsMock:    mock.AchievementExists{Bool: false},
		},
	}),
	constructEvidenceCreateTest(&testInput{
		purpose:            "evidence create db error",
		requestMethod:      consts.POST,
		responseType:       Core,
		responseStatusCode: http.StatusInternalServerError,
		responseMessage:    consts.FriendlyErrorMessage,
		form:               evidenceCreateForm(),
		db: &mock.DB{
			MultimediaTypeExistsMock: mock.MultimediaTypeExists{Bool: true},
			AchievementExistsMock:    mock.AchievementExists{Bool: true},
			EvidenceCreateMock:       mock.EvidenceCreate{Err: mockDbErr},
		},
	}),
	constructEvidenceCreateTest(&testInput{
		purpose:            "evidence create ok",
		requestMethod:      consts.POST,
		responseType:       Retrieve,
		responseStatusCode: http.StatusCreated,
		responseMessage:    fmt.Sprintf(consts.FormatCreated, consts.Evidence),
		form:               evidenceCreateForm(),
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
