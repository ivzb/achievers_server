package controller

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/ivzb/achievers_server/app/model/mock"
)

func userAuthForm() *map[string]string {
	return &map[string]string{
		email:    mockEmail,
		password: mockPassword,
	}
}

var userAuthTests = []*test{
	constructUserAuthTest(&testInput{
		purpose:            "invalid request method",
		requestMethod:      get,
		responseType:       Core,
		responseStatusCode: http.StatusMethodNotAllowed,
		responseMessage:    methodNotAllowed,
		form:               &map[string]string{},
	}),
	constructUserAuthTest(&testInput{
		purpose:            "former error",
		requestMethod:      post,
		responseType:       Core,
		responseStatusCode: http.StatusBadRequest,
		responseMessage:    "former error",
		form:               &map[string]string{},
		former:             &mock.Former{MapMock: mock.Map{Err: mockFormerErr}},
	}),
	constructUserAuthTest(&testInput{
		purpose:            "missing form email",
		requestMethod:      post,
		responseType:       Core,
		responseStatusCode: http.StatusBadRequest,
		responseMessage:    fmt.Sprintf(formatMissing, email),
		form:               mapWithout(userAuthForm(), email),
		former:             &mock.Former{},
	}),
	constructUserAuthTest(&testInput{
		purpose:            "missing form password",
		requestMethod:      post,
		responseType:       Core,
		responseStatusCode: http.StatusBadRequest,
		responseMessage:    fmt.Sprintf(formatMissing, password),
		form:               mapWithout(userAuthForm(), password),
		former:             &mock.Former{},
	}),
	constructUserAuthTest(&testInput{
		purpose:            "user auth db error",
		requestMethod:      post,
		responseType:       Core,
		responseStatusCode: http.StatusInternalServerError,
		responseMessage:    friendlyErrorMessage,
		form:               userAuthForm(),
		former:             &mock.Former{},
		db: &mock.DB{
			UserAuthMock: mock.UserAuth{Err: mockDbErr},
		},
	}),
	constructUserAuthTest(&testInput{
		purpose:            "encrypt error",
		requestMethod:      post,
		responseType:       Core,
		responseStatusCode: http.StatusInternalServerError,
		responseMessage:    friendlyErrorMessage,
		form:               userAuthForm(),
		former:             &mock.Former{},
		db: &mock.DB{
			UserAuthMock: mock.UserAuth{ID: mockID},
		},
		tokener: &mock.Tokener{
			EncryptMock: mock.Encrypt{Err: mockDbErr},
		},
	}),
	constructUserAuthTest(&testInput{
		purpose:            "user auth ok",
		requestMethod:      post,
		responseType:       Retrieve,
		responseStatusCode: http.StatusCreated,
		responseMessage:    authorized,
		form:               userAuthForm(),
		former:             &mock.Former{},
		db: &mock.DB{
			UserAuthMock: mock.UserAuth{ID: mockID},
		},
		tokener: &mock.Tokener{
			EncryptMock: mock.Encrypt{Enc: mockEncrypt},
		},
	}),
}

func constructUserAuthTest(testInput *testInput) *test {
	responseResults, _ := json.Marshal(mockEncrypt)

	return constructTest(UserAuth, testInput, responseResults)
}
