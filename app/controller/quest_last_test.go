package controller

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/ivzb/achievers_server/app/db/mock"
	"github.com/ivzb/achievers_server/app/db/mock/generate"
	"github.com/ivzb/achievers_server/app/shared/consts"
)

var questsLatestArgs = []string{"9"}

var questsLatestTests = []*test{
	constructQuestsLatestTest(&testInput{
		purpose:            "invalid request method",
		requestMethod:      consts.POST,
		responseType:       Core,
		responseStatusCode: http.StatusMethodNotAllowed,
		responseMessage:    consts.MethodNotAllowed,
		form:               &map[string]string{},
		args:               questsLatestArgs,
	}),
	constructQuestsLatestTest(&testInput{
		purpose:            "id QuestsLastID db error",
		requestMethod:      consts.GET,
		responseType:       Core,
		responseStatusCode: http.StatusInternalServerError,
		responseMessage:    consts.FriendlyErrorMessage,
		db: &mock.DB{
			QuestMock: mock.Quest{
				LastIDMock: mock.QuestsLastID{Err: mockDbErr},
			},
		},
		args: questsLatestArgs,
	}),
	constructQuestsLatestTest(&testInput{
		purpose:            "QuestsAfterMock db error",
		requestMethod:      consts.GET,
		responseType:       Core,
		responseStatusCode: http.StatusInternalServerError,
		responseMessage:    consts.FriendlyErrorMessage,
		db: &mock.DB{
			QuestMock: mock.Quest{
				LastIDMock: mock.QuestsLastID{ID: mockID},
				AfterMock:  mock.QuestsAfter{Err: mockDbErr},
			},
		},
		args: questsLatestArgs,
	}),
	constructQuestsLatestTest(&testInput{
		purpose:            "no results",
		requestMethod:      consts.GET,
		responseType:       Core,
		responseStatusCode: http.StatusOK,
		responseMessage:    fmt.Sprintf(consts.FormatFound, consts.Quests),
		db: &mock.DB{
			QuestMock: mock.Quest{
				LastIDMock: mock.QuestsLastID{ID: mockID},
				AfterMock:  mock.QuestsAfter{Qsts: generate.Quests(0)},
			},
		},
		args: []string{"0"},
	}),
	constructQuestsLatestTest(&testInput{
		purpose:            "4 results",
		requestMethod:      consts.GET,
		responseType:       Retrieve,
		responseStatusCode: http.StatusOK,
		responseMessage:    fmt.Sprintf(consts.FormatFound, consts.Quests),
		db: &mock.DB{
			QuestMock: mock.Quest{
				LastIDMock: mock.QuestsLastID{ID: mockID},
				AfterMock:  mock.QuestsAfter{Qsts: generate.Quests(4)},
			},
		},
		args: []string{"4"},
	}),
	constructQuestsLatestTest(&testInput{
		purpose:            "9 results",
		requestMethod:      consts.GET,
		responseType:       Retrieve,
		responseStatusCode: http.StatusOK,
		responseMessage:    fmt.Sprintf(consts.FormatFound, consts.Quests),
		db: &mock.DB{
			QuestMock: mock.Quest{
				LastIDMock: mock.QuestsLastID{ID: mockID},
				AfterMock:  mock.QuestsAfter{Qsts: generate.Quests(9)},
			},
		},
		args: []string{"9"},
	}),
}

func constructQuestsLatestTest(testInput *testInput) *test {
	questsSize, err := strconv.Atoi(testInput.args[0])

	var responseResults []byte

	if err == nil {
		responseResults, _ = json.Marshal(generate.Quests(questsSize))
	}

	return constructTest(QuestsLast, testInput, responseResults)
}
