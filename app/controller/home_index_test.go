package controller

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/ivzb/achievers_server/app/shared/consts"
)

var homeIndexTests = []*test{
	constructHomeIndexTest(&testInput{
		purpose:            "invalid request method",
		requestMethod:      consts.POST,
		responseType:       Core,
		responseStatusCode: http.StatusMethodNotAllowed,
		responseMessage:    consts.MethodNotAllowed,
		form:               &map[string]string{},
	}),
	constructHomeIndexTest(&testInput{
		purpose:            "welcome",
		requestMethod:      consts.GET,
		responseType:       Retrieve,
		responseStatusCode: http.StatusOK,
		responseMessage:    fmt.Sprintf(consts.FormatFound, consts.Home),
		form:               achievementsByQuestIDForm(),
	}),
}

func constructHomeIndexTest(testInput *testInput) *test {
	responseResults, _ := json.Marshal(consts.Welcome)

	return constructTest(HomeIndex, testInput, responseResults)
}
