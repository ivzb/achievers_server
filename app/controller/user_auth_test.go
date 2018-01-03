package controller

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/ivzb/achievers_server/app/model/mock"
)

type userAuthTest struct {
	purpose            string
	requestMethod      string
	responseType       int
	responseStatusCode int
	responseMessage    string
	formerErr          error
	formEmail          string
	formPassword       string
	dbUserAuth         mock.UserAuth
	tokenerEncrypt     mock.Encrypt
}

var userAuthTests = []*test{
	constructUserAuthTest(&userAuthTest{
		purpose:            "invalid request method",
		requestMethod:      get,
		responseType:       Core,
		responseStatusCode: http.StatusMethodNotAllowed,
		responseMessage:    methodNotAllowed,
		formerErr:          nil,
		formEmail:          "",
		formPassword:       "",
		dbUserAuth:         mock.UserAuth{},
		tokenerEncrypt:     mock.Encrypt{},
	}),
	constructUserAuthTest(&userAuthTest{
		purpose:            "former error",
		requestMethod:      post,
		responseType:       Core,
		responseStatusCode: http.StatusBadRequest,
		responseMessage:    "former error",
		formerErr:          mockFormerErr,
		formEmail:          "",
		formPassword:       "",
		dbUserAuth:         mock.UserAuth{},
		tokenerEncrypt:     mock.Encrypt{},
	}),
	constructUserAuthTest(&userAuthTest{
		purpose:            "missing form email",
		requestMethod:      post,
		responseType:       Core,
		responseStatusCode: http.StatusBadRequest,
		responseMessage:    fmt.Sprintf(formatMissing, email),
		formerErr:          nil,
		formEmail:          "",
		formPassword:       mockPassword,
		dbUserAuth:         mock.UserAuth{},
		tokenerEncrypt:     mock.Encrypt{},
	}),
	constructUserAuthTest(&userAuthTest{
		purpose:            "missing form password",
		requestMethod:      post,
		responseType:       Core,
		responseStatusCode: http.StatusBadRequest,
		responseMessage:    fmt.Sprintf(formatMissing, password),
		formerErr:          nil,
		formEmail:          mockEmail,
		formPassword:       "",
		dbUserAuth:         mock.UserAuth{},
		tokenerEncrypt:     mock.Encrypt{},
	}),
	constructUserAuthTest(&userAuthTest{
		purpose:            "user auth db error",
		requestMethod:      post,
		responseType:       Core,
		responseStatusCode: http.StatusInternalServerError,
		responseMessage:    friendlyErrorMessage,
		formerErr:          nil,
		formEmail:          mockEmail,
		formPassword:       mockPassword,
		dbUserAuth:         mock.UserAuth{Err: mockDbErr},
		tokenerEncrypt:     mock.Encrypt{},
	}),
	constructUserAuthTest(&userAuthTest{
		purpose:            "encrypt error",
		requestMethod:      post,
		responseType:       Core,
		responseStatusCode: http.StatusInternalServerError,
		responseMessage:    friendlyErrorMessage,
		formerErr:          nil,
		formEmail:          mockEmail,
		formPassword:       mockPassword,
		dbUserAuth:         mock.UserAuth{ID: mockID},
		tokenerEncrypt:     mock.Encrypt{Err: mockDbErr},
	}),
	constructUserAuthTest(&userAuthTest{
		purpose:            "user auth ok",
		requestMethod:      post,
		responseType:       Retrieve,
		responseStatusCode: http.StatusCreated,
		responseMessage:    authorized,
		formerErr:          nil,
		formEmail:          mockEmail,
		formPassword:       mockPassword,
		dbUserAuth:         mock.UserAuth{ID: mockID},
		tokenerEncrypt:     mock.Encrypt{Enc: mockEncrypt},
	}),
}

func constructUserAuthTest(testInput *userAuthTest) *test {
	responseResults, _ := json.Marshal(mockEncrypt)

	db := &mock.DB{
		UserAuthMock: testInput.dbUserAuth,
	}

	logger := &mock.Logger{}

	former := &mock.Former{
		MapMock: mock.Map{Err: testInput.formerErr},
	}

	tokener := &mock.Tokener{
		EncryptMock: testInput.tokenerEncrypt,
	}

	return &test{
		purpose: testInput.purpose,
		handle:  UserAuth,
		request: constructTestRequest(
			testInput.requestMethod,
			constructForm(map[string]string{
				email:    testInput.formEmail,
				password: testInput.formPassword,
			}),
			constructEnv(db, logger, former, tokener),
		),
		response: constructTestResponse(
			testInput.responseType,
			testInput.responseStatusCode,
			testInput.responseMessage,
			responseResults,
		),
	}
}
