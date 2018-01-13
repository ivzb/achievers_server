package controller

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/ivzb/achievers_server/app/model/mock"
	"github.com/ivzb/achievers_server/app/shared/consts"
)

func questsIndexForm() *map[string]string {
	return &map[string]string{
		consts.Page: mockPage,
	}
}

var questsIndexArgs = []string{"9"}

var questsIndexTests = []*test{
	constructQuestsIndexTest(&testInput{
		purpose:            "invalid request method",
		requestMethod:      consts.POST,
		responseType:       Core,
		responseStatusCode: http.StatusMethodNotAllowed,
		responseMessage:    consts.MethodNotAllowed,
		form:               &map[string]string{},
		args:               questsIndexArgs,
	}),
	constructQuestsIndexTest(&testInput{
		purpose:            "missing consts.Page",
		requestMethod:      consts.GET,
		responseType:       Core,
		responseStatusCode: http.StatusBadRequest,
		responseMessage:    fmt.Sprintf(consts.FormatMissing, consts.Page),
		form:               mapWithout(questsIndexForm(), consts.Page),
		args:               questsIndexArgs,
	}),
	constructQuestsIndexTest(&testInput{
		purpose:            "invalid consts.Page",
		requestMethod:      consts.GET,
		responseType:       Core,
		responseStatusCode: http.StatusBadRequest,
		responseMessage:    fmt.Sprintf(consts.FormatInvalid, consts.Page),
		form: &map[string]string{
			consts.Page: "-1",
		},
		args: questsIndexArgs,
	}),
	constructQuestsIndexTest(&testInput{
		purpose:            "db error",
		requestMethod:      consts.GET,
		responseType:       Core,
		responseStatusCode: http.StatusInternalServerError,
		responseMessage:    consts.FriendlyErrorMessage,
		form:               questsIndexForm(),
		db: &mock.DB{
			QuestsAllMock: mock.QuestsAll{Err: mockDbErr},
		},
		args: questsIndexArgs,
	}),
	constructQuestsIndexTest(&testInput{
		purpose:            "no results on consts.Page",
		requestMethod:      consts.GET,
		responseType:       Core,
		responseStatusCode: http.StatusNotFound,
		responseMessage:    fmt.Sprintf(consts.FormatNotFound, consts.Page),
		form:               questsIndexForm(),
		db: &mock.DB{
			QuestsAllMock: mock.QuestsAll{Qsts: mock.Quests(0)},
		},
		args: []string{"0"},
	}),
	constructQuestsIndexTest(&testInput{
		purpose:            "4 results on consts.Page",
		requestMethod:      consts.GET,
		responseType:       Retrieve,
		responseStatusCode: http.StatusOK,
		responseMessage:    fmt.Sprintf(consts.FormatFound, consts.Quests),
		form:               questsIndexForm(),
		db: &mock.DB{
			QuestsAllMock: mock.QuestsAll{Qsts: mock.Quests(4)},
		},
		args: []string{"4"},
	}),
	constructQuestsIndexTest(&testInput{
		purpose:            "9 results on consts.Page",
		requestMethod:      consts.GET,
		responseType:       Retrieve,
		responseStatusCode: http.StatusOK,
		responseMessage:    fmt.Sprintf(consts.FormatFound, consts.Quests),
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
