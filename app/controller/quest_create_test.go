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
		requestMethod:      get,
		responseType:       Core,
		responseStatusCode: http.StatusMethodNotAllowed,
		responseMessage:    methodNotAllowed,
		form:               &map[string]string{},
	}),
	constructQuestCreateTest(&testInput{
		purpose:            "former error",
		requestMethod:      post,
		responseType:       Core,
		responseStatusCode: http.StatusBadRequest,
		responseMessage:    "former error",
		form:               &map[string]string{},
		former:             &mock.Former{MapMock: mock.Map{Err: mockFormerErr}},
	}),
	constructQuestCreateTest(&testInput{
		purpose:            "missing form title",
		requestMethod:      post,
		responseType:       Core,
		responseStatusCode: http.StatusBadRequest,
		responseMessage:    fmt.Sprintf(formatMissing, title),
		form:               mapWithout(questCreateForm(), title),
		former:             &mock.Former{},
	}),
	constructQuestCreateTest(&testInput{
		purpose:            "missing form picture_url",
		requestMethod:      post,
		responseType:       Core,
		responseStatusCode: http.StatusBadRequest,
		responseMessage:    fmt.Sprintf(formatMissing, pictureURL),
		form:               mapWithout(questCreateForm(), pictureURL),
		former:             &mock.Former{},
	}),
	constructQuestCreateTest(&testInput{
		purpose:            "missing form involvement_id",
		requestMethod:      post,
		responseType:       Core,
		responseStatusCode: http.StatusBadRequest,
		responseMessage:    fmt.Sprintf(formatMissing, involvementID),
		form:               mapWithout(questCreateForm(), involvementID),
		former:             &mock.Former{},
	}),
	constructQuestCreateTest(&testInput{
		purpose:            "missing form quest_type_id",
		requestMethod:      post,
		responseType:       Core,
		responseStatusCode: http.StatusBadRequest,
		responseMessage:    fmt.Sprintf(formatMissing, questTypeID),
		form:               mapWithout(questCreateForm(), questTypeID),
		former:             &mock.Former{},
	}),
	constructQuestCreateTest(&testInput{
		purpose:            "involvement exists db error",
		requestMethod:      post,
		responseType:       Core,
		responseStatusCode: http.StatusInternalServerError,
		responseMessage:    friendlyErrorMessage,
		form:               questCreateForm(),
		former:             &mock.Former{},
		db: &mock.DB{
			InvolvementExistsMock: mock.InvolvementExists{Err: mockDbErr},
		},
	}),
	constructQuestCreateTest(&testInput{
		purpose:            "involvement does not exist",
		requestMethod:      post,
		responseType:       Core,
		responseStatusCode: http.StatusNotFound,
		responseMessage:    fmt.Sprintf(formatNotFound, involvementID),
		form:               questCreateForm(),
		former:             &mock.Former{},
		db: &mock.DB{
			InvolvementExistsMock: mock.InvolvementExists{Bool: false},
		},
	}),
	constructQuestCreateTest(&testInput{
		purpose:            "quest_type_id exists db error",
		requestMethod:      post,
		responseType:       Core,
		responseStatusCode: http.StatusInternalServerError,
		responseMessage:    friendlyErrorMessage,
		form:               questCreateForm(),
		former:             &mock.Former{},
		db: &mock.DB{
			InvolvementExistsMock: mock.InvolvementExists{Bool: true},
			QuestTypeExistsMock:   mock.QuestTypeExists{Err: mockDbErr},
		},
	}),
	constructQuestCreateTest(&testInput{
		purpose:            "quest_type_id does not exist",
		requestMethod:      post,
		responseType:       Core,
		responseStatusCode: http.StatusNotFound,
		responseMessage:    fmt.Sprintf(formatNotFound, questTypeID),
		form:               questCreateForm(),
		former:             &mock.Former{},
		db: &mock.DB{
			InvolvementExistsMock: mock.InvolvementExists{Bool: true},
			QuestTypeExistsMock:   mock.QuestTypeExists{Bool: false},
		},
	}),
	constructQuestCreateTest(&testInput{
		purpose:            "quest create db error",
		requestMethod:      post,
		responseType:       Core,
		responseStatusCode: http.StatusInternalServerError,
		responseMessage:    friendlyErrorMessage,
		form:               questCreateForm(),
		former:             &mock.Former{},
		db: &mock.DB{
			InvolvementExistsMock: mock.InvolvementExists{Bool: true},
			QuestTypeExistsMock:   mock.QuestTypeExists{Bool: true},
			QuestCreateMock:       mock.QuestCreate{Err: mockDbErr},
		},
	}),
	constructQuestCreateTest(&testInput{
		purpose:            "quest create ok",
		requestMethod:      post,
		responseType:       Retrieve,
		responseStatusCode: http.StatusOK,
		responseMessage:    fmt.Sprintf(formatCreated, quest),
		form:               questCreateForm(),
		former:             &mock.Former{},
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
