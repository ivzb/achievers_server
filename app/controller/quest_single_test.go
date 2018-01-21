package controller

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/ivzb/achievers_server/app/db/mock"
	"github.com/ivzb/achievers_server/app/shared/consts"
)

func questSingleForm() *map[string]string {
	return &map[string]string{
		consts.ID: mockID,
	}
}

var questSingleTests = []*test{
	constructQuestSingleTest(&testInput{
		purpose:            "invalid request method",
		requestMethod:      consts.POST,
		responseType:       Core,
		responseStatusCode: http.StatusMethodNotAllowed,
		responseMessage:    consts.MethodNotAllowed,
		form:               &map[string]string{},
	}),
	constructQuestSingleTest(&testInput{
		purpose:            "missing id",
		requestMethod:      consts.GET,
		responseType:       Core,
		responseStatusCode: http.StatusBadRequest,
		responseMessage:    fmt.Sprintf(consts.FormatMissing, consts.ID),
		form:               mapWithout(questSingleForm(), consts.ID),
	}),
	constructQuestSingleTest(&testInput{
		purpose:            "quest exists db error",
		requestMethod:      consts.GET,
		responseType:       Core,
		responseStatusCode: http.StatusInternalServerError,
		responseMessage:    consts.FriendlyErrorMessage,
		form:               questSingleForm(),
		db: &mock.DB{
			QuestExistsMock: mock.QuestExists{Err: mockDbErr},
		},
	}),
	constructQuestSingleTest(&testInput{
		purpose:            "quest does not exist",
		requestMethod:      consts.GET,
		responseType:       Core,
		responseStatusCode: http.StatusNotFound,
		responseMessage:    fmt.Sprintf(consts.FormatNotFound, consts.Quest),
		form:               questSingleForm(),
		db: &mock.DB{
			QuestExistsMock: mock.QuestExists{Bool: false},
		},
	}),
	constructQuestSingleTest(&testInput{
		purpose:            "quest single db error",
		requestMethod:      consts.GET,
		responseType:       Core,
		responseStatusCode: http.StatusInternalServerError,
		responseMessage:    consts.FriendlyErrorMessage,
		form:               questSingleForm(),
		db: &mock.DB{
			QuestExistsMock: mock.QuestExists{Bool: true},
			QuestSingleMock: mock.QuestSingle{Err: mockDbErr},
		},
	}),
	constructQuestSingleTest(&testInput{
		purpose:            "quest single OK",
		requestMethod:      consts.GET,
		responseType:       Retrieve,
		responseStatusCode: http.StatusOK,
		responseMessage:    fmt.Sprintf(consts.FormatFound, consts.Quest),
		form:               questSingleForm(),
		db: &mock.DB{
			QuestExistsMock: mock.QuestExists{Bool: true},
			QuestSingleMock: mock.QuestSingle{Qst: mock.Quest()},
		},
	}),
}

func constructQuestSingleTest(testInput *testInput) *test {
	responseResults, _ := json.Marshal(mock.Quest())

	return constructTest(QuestSingle, testInput, responseResults)
}
