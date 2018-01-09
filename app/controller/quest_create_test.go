package controller

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/ivzb/achievers_server/app/model/mock"
)

func questCreateForm() *map[string]string {
	return &map[string]string{
		title:         mockTitle,
		pictureURL:    mockPictureURL,
		involvementID: mockInvolvementID,
		questTypeID:   mockQuestTypeID,
	}
}

var questCreateTests = []*test{
	constructQuestCreateTest(&testInput{
		purpose:            "invalid request method",
		requestMethod:      GET,
		responseType:       Core,
		responseStatusCode: http.StatusMethodNotAllowed,
		responseMessage:    methodNotAllowed,
		form:               &map[string]string{},
	}),
	constructQuestCreateTest(&testInput{
		purpose:            "former error",
		requestMethod:      POST,
		responseType:       Core,
		responseStatusCode: http.StatusBadRequest,
		responseMessage:    "content-type of request is incorrect",
		form:               &map[string]string{},
		removeHeaders:      true,
	}),
	constructQuestCreateTest(&testInput{
		purpose:            "missing form title",
		requestMethod:      POST,
		responseType:       Core,
		responseStatusCode: http.StatusBadRequest,
		responseMessage:    fmt.Sprintf(formatMissing, title),
		form:               mapWithout(questCreateForm(), title),
	}),
	constructQuestCreateTest(&testInput{
		purpose:            "missing form picture_url",
		requestMethod:      POST,
		responseType:       Core,
		responseStatusCode: http.StatusBadRequest,
		responseMessage:    fmt.Sprintf(formatMissing, pictureURL),
		form:               mapWithout(questCreateForm(), pictureURL),
	}),
	constructQuestCreateTest(&testInput{
		purpose:            "missing form involvement_id",
		requestMethod:      POST,
		responseType:       Core,
		responseStatusCode: http.StatusBadRequest,
		responseMessage:    fmt.Sprintf(formatMissing, involvementID),
		form:               mapWithout(questCreateForm(), involvementID),
	}),
	constructQuestCreateTest(&testInput{
		purpose:            "missing form quest_type_id",
		requestMethod:      POST,
		responseType:       Core,
		responseStatusCode: http.StatusBadRequest,
		responseMessage:    fmt.Sprintf(formatMissing, questTypeID),
		form:               mapWithout(questCreateForm(), questTypeID),
	}),
	constructQuestCreateTest(&testInput{
		purpose:            "involvement exists db error",
		requestMethod:      POST,
		responseType:       Core,
		responseStatusCode: http.StatusInternalServerError,
		responseMessage:    friendlyErrorMessage,
		form:               questCreateForm(),
		db: &mock.DB{
			InvolvementExistsMock: mock.InvolvementExists{Err: mockDbErr},
		},
	}),
	constructQuestCreateTest(&testInput{
		purpose:            "involvement does not exist",
		requestMethod:      POST,
		responseType:       Core,
		responseStatusCode: http.StatusNotFound,
		responseMessage:    fmt.Sprintf(formatNotFound, involvementID),
		form:               questCreateForm(),
		db: &mock.DB{
			InvolvementExistsMock: mock.InvolvementExists{Bool: false},
		},
	}),
	constructQuestCreateTest(&testInput{
		purpose:            "quest_type_id exists db error",
		requestMethod:      POST,
		responseType:       Core,
		responseStatusCode: http.StatusInternalServerError,
		responseMessage:    friendlyErrorMessage,
		form:               questCreateForm(),
		db: &mock.DB{
			InvolvementExistsMock: mock.InvolvementExists{Bool: true},
			QuestTypeExistsMock:   mock.QuestTypeExists{Err: mockDbErr},
		},
	}),
	constructQuestCreateTest(&testInput{
		purpose:            "quest_type_id does not exist",
		requestMethod:      POST,
		responseType:       Core,
		responseStatusCode: http.StatusNotFound,
		responseMessage:    fmt.Sprintf(formatNotFound, questTypeID),
		form:               questCreateForm(),
		db: &mock.DB{
			InvolvementExistsMock: mock.InvolvementExists{Bool: true},
			QuestTypeExistsMock:   mock.QuestTypeExists{Bool: false},
		},
	}),
	constructQuestCreateTest(&testInput{
		purpose:            "quest create db error",
		requestMethod:      POST,
		responseType:       Core,
		responseStatusCode: http.StatusInternalServerError,
		responseMessage:    friendlyErrorMessage,
		form:               questCreateForm(),
		db: &mock.DB{
			InvolvementExistsMock: mock.InvolvementExists{Bool: true},
			QuestTypeExistsMock:   mock.QuestTypeExists{Bool: true},
			QuestCreateMock:       mock.QuestCreate{Err: mockDbErr},
		},
	}),
	constructQuestCreateTest(&testInput{
		purpose:            "quest create ok",
		requestMethod:      POST,
		responseType:       Retrieve,
		responseStatusCode: http.StatusCreated,
		responseMessage:    fmt.Sprintf(formatCreated, quest),
		form:               questCreateForm(),
		db: &mock.DB{
			InvolvementExistsMock: mock.InvolvementExists{Bool: true},
			QuestTypeExistsMock:   mock.QuestTypeExists{Bool: true},
			QuestCreateMock:       mock.QuestCreate{ID: mockID},
		},
	}),
}

func constructQuestCreateTest(testInput *testInput) *test {
	responseResults, _ := json.Marshal(mockID)
	return constructTest(QuestCreate, testInput, responseResults)
}
