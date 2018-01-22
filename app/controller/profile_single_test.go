package controller

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/ivzb/achievers_server/app/db/mock"
	"github.com/ivzb/achievers_server/app/db/mock/generate"
	"github.com/ivzb/achievers_server/app/shared/consts"
)

func profileSingleForm() *map[string]string {
	return &map[string]string{
		consts.ID: mockID,
	}
}

var profileSingleTests = []*test{
	constructProfileSingleTest(&testInput{
		purpose:            "invalid request method",
		requestMethod:      consts.POST,
		responseType:       Core,
		responseStatusCode: http.StatusMethodNotAllowed,
		responseMessage:    consts.MethodNotAllowed,
		form:               &map[string]string{},
	}),
	constructProfileSingleTest(&testInput{
		purpose:            "missing id",
		requestMethod:      consts.GET,
		responseType:       Core,
		responseStatusCode: http.StatusBadRequest,
		responseMessage:    fmt.Sprintf(consts.FormatMissing, consts.ID),
		form:               mapWithout(profileSingleForm(), consts.ID),
	}),
	constructProfileSingleTest(&testInput{
		purpose:            "profile exists db error",
		requestMethod:      consts.GET,
		responseType:       Core,
		responseStatusCode: http.StatusInternalServerError,
		responseMessage:    consts.FriendlyErrorMessage,
		form:               profileSingleForm(),
		db: &mock.DB{
			ProfileMock: mock.Profile{
				ExistsMock: mock.ProfileExists{Err: mockDbErr},
			},
		},
	}),
	constructProfileSingleTest(&testInput{
		purpose:            "profile does not exist",
		requestMethod:      consts.GET,
		responseType:       Core,
		responseStatusCode: http.StatusNotFound,
		responseMessage:    fmt.Sprintf(consts.FormatNotFound, consts.Profile),
		form:               profileSingleForm(),
		db: &mock.DB{
			ProfileMock: mock.Profile{
				ExistsMock: mock.ProfileExists{Bool: false},
			},
		},
	}),
	constructProfileSingleTest(&testInput{
		purpose:            "profile single db error",
		requestMethod:      consts.GET,
		responseType:       Core,
		responseStatusCode: http.StatusInternalServerError,
		responseMessage:    consts.FriendlyErrorMessage,
		form:               profileSingleForm(),
		db: &mock.DB{
			ProfileMock: mock.Profile{
				ExistsMock: mock.ProfileExists{Bool: true},
				SingleMock: mock.ProfileSingle{Err: mockDbErr},
			},
		},
	}),
	constructProfileSingleTest(&testInput{
		purpose:            "profile single OK",
		requestMethod:      consts.GET,
		responseType:       Retrieve,
		responseStatusCode: http.StatusOK,
		responseMessage:    fmt.Sprintf(consts.FormatFound, consts.Profile),
		form:               profileSingleForm(),
		db: &mock.DB{
			ProfileMock: mock.Profile{
				ExistsMock: mock.ProfileExists{Bool: true},
				SingleMock: mock.ProfileSingle{Prfl: generate.Profile()},
			},
		},
	}),
}

func constructProfileSingleTest(testInput *testInput) *test {
	responseResults, _ := json.Marshal(generate.Profile())

	return constructTest(ProfileSingle, testInput, responseResults)
}
