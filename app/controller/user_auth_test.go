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
		requestMethod:      GET,
		responseType:       Core,
		responseStatusCode: http.StatusMethodNotAllowed,
		responseMessage:    methodNotAllowed,
		form:               &map[string]string{},
	}),
	constructUserAuthTest(&testInput{
		purpose:            "former error",
		requestMethod:      POST,
		responseType:       Core,
		responseStatusCode: http.StatusBadRequest,
		responseMessage:    "content-type of request is incorrect",
		form:               &map[string]string{},
		removeHeaders:      true,
	}),
	constructUserAuthTest(&testInput{
		purpose:            "missing form email",
		requestMethod:      POST,
		responseType:       Core,
		responseStatusCode: http.StatusBadRequest,
		responseMessage:    fmt.Sprintf(formatMissing, email),
		form:               mapWithout(userAuthForm(), email),
	}),
	constructUserAuthTest(&testInput{
		purpose:            "missing form password",
		requestMethod:      POST,
		responseType:       Core,
		responseStatusCode: http.StatusBadRequest,
		responseMessage:    fmt.Sprintf(formatMissing, password),
		form:               mapWithout(userAuthForm(), password),
	}),
	constructUserAuthTest(&testInput{
		purpose:            "user email exists db error",
		requestMethod:      POST,
		responseType:       Core,
		responseStatusCode: http.StatusInternalServerError,
		responseMessage:    friendlyErrorMessage,
		form:               userAuthForm(),
		db: &mock.DB{
			UserEmailExistsMock: mock.UserEmailExists{Err: mockDbErr},
		},
	}),
	constructUserAuthTest(&testInput{
		purpose:            "user email does not exist",
		requestMethod:      POST,
		responseType:       Core,
		responseStatusCode: http.StatusNotFound,
		responseMessage:    fmt.Sprintf(formatNotFound, email),
		form:               userAuthForm(),
		db: &mock.DB{
			UserEmailExistsMock: mock.UserEmailExists{Bool: false},
		},
	}),
	constructUserAuthTest(&testInput{
		purpose:            "user auth db error",
		requestMethod:      POST,
		responseType:       Core,
		responseStatusCode: http.StatusInternalServerError,
		responseMessage:    friendlyErrorMessage,
		form:               userAuthForm(),
		db: &mock.DB{
			UserEmailExistsMock: mock.UserEmailExists{Bool: true},
			UserAuthMock:        mock.UserAuth{Err: mockDbErr},
		},
	}),
	constructUserAuthTest(&testInput{
		purpose:            "encrypt error",
		requestMethod:      POST,
		responseType:       Core,
		responseStatusCode: http.StatusInternalServerError,
		responseMessage:    friendlyErrorMessage,
		form:               userAuthForm(),
		db: &mock.DB{
			UserEmailExistsMock: mock.UserEmailExists{Bool: true},
			UserAuthMock:        mock.UserAuth{ID: mockID},
		},
		tokener: &mock.Tokener{
			EncryptMock: mock.Encrypt{Err: mockDbErr},
		},
	}),
	constructUserAuthTest(&testInput{
		purpose:            "user auth ok",
		requestMethod:      POST,
		responseType:       Retrieve,
		responseStatusCode: http.StatusCreated,
		responseMessage:    fmt.Sprintf(formatCreated, authToken),
		form:               userAuthForm(),
		db: &mock.DB{
			UserEmailExistsMock: mock.UserEmailExists{Bool: true},
			UserAuthMock:        mock.UserAuth{ID: mockID},
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
