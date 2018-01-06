package controller

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/ivzb/achievers_server/app/model/mock"
)

func questsIndexForm() *map[string]string {
	return &map[string]string{
		page: mockPage,
	}
}

var questsIndexArgs = []string{"9"}

var questsIndexTests = []*test{
	constructQuestsIndexTest(&testInput{
		purpose:            "invalid request method",
		requestMethod:      POST,
		responseType:       Core,
		responseStatusCode: http.StatusMethodNotAllowed,
		responseMessage:    methodNotAllowed,
		form:               &map[string]string{},
		args:               questsIndexArgs,
	}),
	constructQuestsIndexTest(&testInput{
		purpose:            "missing page",
		requestMethod:      GET,
		responseType:       Core,
		responseStatusCode: http.StatusBadRequest,
		responseMessage:    fmt.Sprintf(formatMissing, page),
		form:               mapWithout(questsIndexForm(), page),
		args:               questsIndexArgs,
	}),
	constructQuestsIndexTest(&testInput{
		purpose:            "invalid page",
		requestMethod:      GET,
		responseType:       Core,
		responseStatusCode: http.StatusBadRequest,
		responseMessage:    fmt.Sprintf(formatInvalid, page),
		form: &map[string]string{
			page: "-1",
		},
		args: questsIndexArgs,
	}),
	constructQuestsIndexTest(&testInput{
		purpose:            "db error",
		requestMethod:      GET,
		responseType:       Core,
		responseStatusCode: http.StatusInternalServerError,
		responseMessage:    friendlyErrorMessage,
		form:               questsIndexForm(),
		db: &mock.DB{
			QuestsAllMock: mock.QuestsAll{Err: mockDbErr},
		},
		args: questsIndexArgs,
	}),
	constructQuestsIndexTest(&testInput{
		purpose:            "no results on page",
		requestMethod:      GET,
		responseType:       Core,
		responseStatusCode: http.StatusNotFound,
		responseMessage:    fmt.Sprintf(formatNotFound, page),
		form:               questsIndexForm(),
		db: &mock.DB{
			QuestsAllMock: mock.QuestsAll{Qsts: mock.Quests(0)},
		},
		args: []string{"0"},
	}),
	constructQuestsIndexTest(&testInput{
		purpose:            "4 results on page",
		requestMethod:      GET,
		responseType:       Retrieve,
		responseStatusCode: http.StatusOK,
		responseMessage:    fmt.Sprintf(formatFound, quests),
		form:               questsIndexForm(),
		db: &mock.DB{
			QuestsAllMock: mock.QuestsAll{Qsts: mock.Quests(4)},
		},
		args: []string{"4"},
	}),
	constructQuestsIndexTest(&testInput{
		purpose:            "9 results on page",
		requestMethod:      GET,
		responseType:       Retrieve,
		responseStatusCode: http.StatusOK,
		responseMessage:    fmt.Sprintf(formatFound, quests),
		form:               questsIndexForm(),
		db: &mock.DB{
			QuestsAllMock: mock.QuestsAll{Qsts: mock.Quests(9)},
		},
		args: []string{"9"},
	}),
}

func constructQuestsIndexTest(testInput *testInput) *test {
	questsSize, err := strconv.Atoi(testInput.args[0])

	var responseResults []byte

	if err == nil {
		responseResults, _ = json.Marshal(mock.Quests(questsSize))
	}

	return constructTest(QuestsIndex, testInput, responseResults)
}
