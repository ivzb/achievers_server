package controller

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/ivzb/achievers_server/app/db/mock"
	"github.com/ivzb/achievers_server/app/shared/consts"
)

func questCreateForm() *map[string]string {
	return &map[string]string{
		consts.Title:         mockTitle,
		consts.PictureURL:    mockPictureURL,
		consts.InvolvementID: mockInvolvementID,
		consts.QuestTypeID:   mockQuestTypeID,
	}
}

var questCreateTests = []*test{
	constructQuestCreateTest(&testInput{
		purpose:            "invalid request method",
		requestMethod:      consts.GET,
		responseType:       Core,
		responseStatusCode: http.StatusMethodNotAllowed,
		responseMessage:    consts.MethodNotAllowed,
		form:               &map[string]string{},
	}),
	constructQuestCreateTest(&testInput{
		purpose:            "former error",
		requestMethod:      consts.POST,
		responseType:       Core,
		responseStatusCode: http.StatusBadRequest,
		responseMessage:    "content-type of request is incorrect",
		form:               &map[string]string{},
		removeHeaders:      true,
	}),
	constructQuestCreateTest(&testInput{
		purpose:            "missing form consts.Title",
		requestMethod:      consts.POST,
		responseType:       Core,
		responseStatusCode: http.StatusBadRequest,
		responseMessage:    fmt.Sprintf(consts.FormatBlank, consts.Title),
		form:               mapWithout(questCreateForm(), consts.Title),
	}),
	constructQuestCreateTest(&testInput{
		purpose:            "missing form picture_url",
		requestMethod:      consts.POST,
		responseType:       Core,
		responseStatusCode: http.StatusBadRequest,
		responseMessage:    fmt.Sprintf(consts.FormatBlank, consts.PictureURL),
		form:               mapWithout(questCreateForm(), consts.PictureURL),
	}),
	constructQuestCreateTest(&testInput{
		purpose:            "missing form involvement_id",
		requestMethod:      consts.POST,
		responseType:       Core,
		responseStatusCode: http.StatusBadRequest,
		responseMessage:    fmt.Sprintf(consts.FormatValidID, consts.InvolvementID),
		form:               mapWithout(questCreateForm(), consts.InvolvementID),
	}),
	constructQuestCreateTest(&testInput{
		purpose:            "missing form quest_type_id",
		requestMethod:      consts.POST,
		responseType:       Core,
		responseStatusCode: http.StatusBadRequest,
		responseMessage:    fmt.Sprintf(consts.FormatValidID, consts.QuestTypeID),
		form:               mapWithout(questCreateForm(), consts.QuestTypeID),
	}),
	constructQuestCreateTest(&testInput{
		purpose:            "involvement exists db error",
		requestMethod:      consts.POST,
		responseType:       Core,
		responseStatusCode: http.StatusInternalServerError,
		responseMessage:    consts.FriendlyErrorMessage,
		form:               questCreateForm(),
		db: &mock.DB{
			InvolvementMock: mock.Involvement{
				ExistsMock: mock.InvolvementExists{Err: mockDbErr},
			},
		},
	}),
	constructQuestCreateTest(&testInput{
		purpose:            "involvement does not exist",
		requestMethod:      consts.POST,
		responseType:       Core,
		responseStatusCode: http.StatusNotFound,
		responseMessage:    fmt.Sprintf(consts.FormatNotFound, consts.InvolvementID),
		form:               questCreateForm(),
		db: &mock.DB{
			InvolvementMock: mock.Involvement{
				ExistsMock: mock.InvolvementExists{Bool: false},
			},
		},
	}),
	constructQuestCreateTest(&testInput{
		purpose:            "quest_type_id exists db error",
		requestMethod:      consts.POST,
		responseType:       Core,
		responseStatusCode: http.StatusInternalServerError,
		responseMessage:    consts.FriendlyErrorMessage,
		form:               questCreateForm(),
		db: &mock.DB{
			InvolvementMock: mock.Involvement{
				ExistsMock: mock.InvolvementExists{Bool: true},
			},
			QuestTypeMock: mock.QuestType{
				ExistsMock: mock.QuestTypeExists{Err: mockDbErr},
			},
		},
	}),
	constructQuestCreateTest(&testInput{
		purpose:            "quest_type_id does not exist",
		requestMethod:      consts.POST,
		responseType:       Core,
		responseStatusCode: http.StatusNotFound,
		responseMessage:    fmt.Sprintf(consts.FormatNotFound, consts.QuestTypeID),
		form:               questCreateForm(),
		db: &mock.DB{
			InvolvementMock: mock.Involvement{
				ExistsMock: mock.InvolvementExists{Bool: true},
			},
			QuestTypeMock: mock.QuestType{
				ExistsMock: mock.QuestTypeExists{Bool: false},
			},
		},
	}),
	constructQuestCreateTest(&testInput{
		purpose:            "quest create db error",
		requestMethod:      consts.POST,
		responseType:       Core,
		responseStatusCode: http.StatusInternalServerError,
		responseMessage:    consts.FriendlyErrorMessage,
		form:               questCreateForm(),
		db: &mock.DB{
			InvolvementMock: mock.Involvement{
				ExistsMock: mock.InvolvementExists{Bool: true},
			},
			QuestTypeMock: mock.QuestType{
				ExistsMock: mock.QuestTypeExists{Bool: true},
			},
			QuestMock: mock.Quest{
				CreateMock: mock.QuestCreate{Err: mockDbErr},
			},
		},
	}),
	constructQuestCreateTest(&testInput{
		purpose:            "quest create ok",
		requestMethod:      consts.POST,
		responseType:       Retrieve,
		responseStatusCode: http.StatusCreated,
		responseMessage:    fmt.Sprintf(consts.FormatCreated, consts.Quest),
		form:               questCreateForm(),
		db: &mock.DB{
			InvolvementMock: mock.Involvement{
				ExistsMock: mock.InvolvementExists{Bool: true},
			},
			QuestTypeMock: mock.QuestType{
				ExistsMock: mock.QuestTypeExists{Bool: true},
			},
			QuestMock: mock.Quest{
				CreateMock: mock.QuestCreate{ID: mockID},
			},
		},
	}),
}

func constructQuestCreateTest(testInput *testInput) *test {
	responseResults, _ := json.Marshal(mockID)
	return constructTest(QuestCreate, testInput, responseResults)
}
