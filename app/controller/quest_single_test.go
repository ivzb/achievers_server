package controller

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/ivzb/achievers_server/app/model/mock"
)

func questSingleForm() *map[string]string {
	return &map[string]string{
		id: mockID,
	}
}

var questSingleTests = []*test{
	constructQuestSingleTest(&testInput{
		purpose:            "invalid request method",
		requestMethod:      POST,
		responseType:       Core,
		responseStatusCode: http.StatusMethodNotAllowed,
		responseMessage:    methodNotAllowed,
		form:               &map[string]string{},
	}),
	constructQuestSingleTest(&testInput{
		purpose:            "missing id",
		requestMethod:      GET,
		responseType:       Core,
		responseStatusCode: http.StatusBadRequest,
		responseMessage:    fmt.Sprintf(formatMissing, id),
		form:               mapWithout(questSingleForm(), id),
	}),
	constructQuestSingleTest(&testInput{
		purpose:            "quest exists db error",
		requestMethod:      GET,
		responseType:       Core,
		responseStatusCode: http.StatusInternalServerError,
		responseMessage:    friendlyErrorMessage,
		form:               questSingleForm(),
		db: &mock.DB{
			QuestExistsMock: mock.QuestExists{Err: mockDbErr},
		},
	}),
	constructQuestSingleTest(&testInput{
		purpose:            "quest does not exist",
		requestMethod:      GET,
		responseType:       Core,
		responseStatusCode: http.StatusNotFound,
		responseMessage:    fmt.Sprintf(formatNotFound, quest),
		form:               questSingleForm(),
		db: &mock.DB{
			QuestExistsMock: mock.QuestExists{Bool: false},
		},
	}),
	constructQuestSingleTest(&testInput{
		purpose:            "quest single db error",
		requestMethod:      GET,
		responseType:       Core,
		responseStatusCode: http.StatusInternalServerError,
		responseMessage:    friendlyErrorMessage,
		form:               questSingleForm(),
		db: &mock.DB{
			QuestExistsMock: mock.QuestExists{Bool: true},
			QuestSingleMock: mock.QuestSingle{Err: mockDbErr},
		},
	}),
	constructQuestSingleTest(&testInput{
		purpose:            "quest single OK",
		requestMethod:      GET,
		responseType:       Retrieve,
		responseStatusCode: http.StatusOK,
		responseMessage:    fmt.Sprintf(formatFound, quest),
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
