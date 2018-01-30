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

func questsAfterForm() *map[string]string {
	return &map[string]string{
		consts.ID: mockID,
	}
}

var questsAfterArgs = []string{"9"}

var questsAfterTests = []*test{
	constructQuestsAfterTest(&testInput{
		purpose:            "invalid request method",
		requestMethod:      consts.POST,
		responseType:       Core,
		responseStatusCode: http.StatusMethodNotAllowed,
		responseMessage:    consts.MethodNotAllowed,
		form:               &map[string]string{},
		args:               questsAfterArgs,
	}),
	constructQuestsAfterTest(&testInput{
		purpose:            "missing id",
		requestMethod:      consts.GET,
		responseType:       Core,
		responseStatusCode: http.StatusBadRequest,
		responseMessage:    fmt.Sprintf(consts.FormatMissing, consts.ID),
		form:               mapWithout(questsAfterForm(), consts.ID),
		db: &mock.DB{
			QuestMock: mock.Quest{
				ExistsMock: mock.QuestExists{Err: mockDbErr},
			},
		},
		args: questsAfterArgs,
	}),
	constructQuestsAfterTest(&testInput{
		purpose:            "id QuestExists db error",
		requestMethod:      consts.GET,
		responseType:       Core,
		responseStatusCode: http.StatusInternalServerError,
		responseMessage:    consts.FriendlyErrorMessage,
		form:               questsAfterForm(),
		db: &mock.DB{
			QuestMock: mock.Quest{
				ExistsMock: mock.QuestExists{Err: mockDbErr},
			},
		},
		args: questsAfterArgs,
	}),
	constructQuestsAfterTest(&testInput{
		purpose:            "id QuestExists does not exist",
		requestMethod:      consts.GET,
		responseType:       Core,
		responseStatusCode: http.StatusNotFound,
		responseMessage:    fmt.Sprintf(consts.FormatNotFound, consts.ID),
		form:               questsAfterForm(),
		db: &mock.DB{
			QuestMock: mock.Quest{
				ExistsMock: mock.QuestExists{Bool: false},
			},
		},
		args: questsAfterArgs,
	}),
	constructQuestsAfterTest(&testInput{
		purpose:            "QuestsAfterMock db error",
		requestMethod:      consts.GET,
		responseType:       Core,
		responseStatusCode: http.StatusInternalServerError,
		responseMessage:    consts.FriendlyErrorMessage,
		form:               questsAfterForm(),
		db: &mock.DB{
			QuestMock: mock.Quest{
				ExistsMock: mock.QuestExists{Bool: true},
				AfterMock:  mock.QuestsAfter{Err: mockDbErr},
			},
		},
		args: questsAfterArgs,
	}),
	constructQuestsAfterTest(&testInput{
		purpose:            "no results",
		requestMethod:      consts.GET,
		responseType:       Core,
		responseStatusCode: http.StatusOK,
		responseMessage:    fmt.Sprintf(consts.FormatFound, consts.Quests),
		form:               questsAfterForm(),
		db: &mock.DB{
			QuestMock: mock.Quest{
				ExistsMock: mock.QuestExists{Bool: true},
				AfterMock:  mock.QuestsAfter{Qsts: generate.Quests(0)},
			},
		},
		args: []string{"0"},
	}),
	constructQuestsAfterTest(&testInput{
		purpose:            "4 results",
		requestMethod:      consts.GET,
		responseType:       Retrieve,
		responseStatusCode: http.StatusOK,
		responseMessage:    fmt.Sprintf(consts.FormatFound, consts.Quests),
		form:               questsAfterForm(),
		db: &mock.DB{
			QuestMock: mock.Quest{
				ExistsMock: mock.QuestExists{Bool: true},
				AfterMock:  mock.QuestsAfter{Qsts: generate.Quests(4)},
			},
		},
		args: []string{"4"},
	}),
	constructQuestsAfterTest(&testInput{
		purpose:            "9 results",
		requestMethod:      consts.GET,
		responseType:       Retrieve,
		responseStatusCode: http.StatusOK,
		responseMessage:    fmt.Sprintf(consts.FormatFound, consts.Quests),
		form:               questsAfterForm(),
		db: &mock.DB{
			QuestMock: mock.Quest{
				ExistsMock: mock.QuestExists{Bool: true},
				AfterMock:  mock.QuestsAfter{Qsts: generate.Quests(9)},
			},
		},
		args: []string{"9"},
	}),
}

func constructQuestsAfterTest(testInput *testInput) *test {
	questsSize, err := strconv.Atoi(testInput.args[0])

	var responseResults []byte

	if err == nil {
		responseResults, _ = json.Marshal(generate.Quests(questsSize))
	}

	return constructTest(QuestsAfter, testInput, responseResults)
}
