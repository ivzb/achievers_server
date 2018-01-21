package controller

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/ivzb/achievers_server/app/db/mock"
	"github.com/ivzb/achievers_server/app/shared/consts"
)

func profileMeForm() *map[string]string {
	return &map[string]string{}
}

var profileMeTests = []*test{
	constructProfileMeTest(&testInput{
		purpose:            "invalid request method",
		requestMethod:      consts.POST,
		responseType:       Core,
		responseStatusCode: http.StatusMethodNotAllowed,
		responseMessage:    consts.MethodNotAllowed,
		form:               &map[string]string{},
	}),
	constructProfileMeTest(&testInput{
		purpose:            "profile single db error",
		requestMethod:      consts.GET,
		responseType:       Core,
		responseStatusCode: http.StatusInternalServerError,
		responseMessage:    consts.FriendlyErrorMessage,
		form:               profileMeForm(),
		db: &mock.DB{
			ProfileByUserIDMock: mock.ProfileByUserID{Err: mockDbErr},
		},
	}),
	constructProfileMeTest(&testInput{
		purpose:            "profile single OK",
		requestMethod:      consts.GET,
		responseType:       Retrieve,
		responseStatusCode: http.StatusOK,
		responseMessage:    fmt.Sprintf(consts.FormatFound, consts.Profile),
		form:               profileMeForm(),
		db: &mock.DB{
			ProfileByUserIDMock: mock.ProfileByUserID{Prfl: mock.Profile()},
		},
	}),
}

func constructProfileMeTest(testInput *testInput) *test {
	responseResults, _ := json.Marshal(mock.Profile())

	return constructTest(ProfileMe, testInput, responseResults)
}
