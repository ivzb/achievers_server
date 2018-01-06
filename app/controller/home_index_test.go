package controller

import (
	"encoding/json"
	"fmt"
	"net/http"
)

var homeIndexTests = []*test{
	constructHomeIndexTest(&testInput{
		purpose:            "invalid request method",
		requestMethod:      POST,
		responseType:       Core,
		responseStatusCode: http.StatusMethodNotAllowed,
		responseMessage:    methodNotAllowed,
		form:               &map[string]string{},
	}),
	constructHomeIndexTest(&testInput{
		purpose:            "welcome",
		requestMethod:      GET,
		responseType:       Retrieve,
		responseStatusCode: http.StatusOK,
		responseMessage:    fmt.Sprintf(formatFound, home),
		form:               achievementsByQuestIDForm(),
	}),
}

func constructHomeIndexTest(testInput *testInput) *test {
	responseResults, _ := json.Marshal(welcome)

	return constructTest(HomeIndex, testInput, responseResults)
}
