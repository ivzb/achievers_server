package controller

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/ivzb/achievers_server/app/model/mock"
)

type userCreateTest struct {
	purpose            string
	requestMethod      string
	responseType       int
	responseStatusCode int
	responseMessage    string
	formerErr          error
	formFirstName      string
	formLastName       string
	formEmail          string
	formPassword       string
	dbUserEmailExists  mock.UserEmailExists
	dbUserCreate       mock.UserCreate
}

var userCreateTests = []*test{
	constructUserCreateTest(&userCreateTest{
		purpose:            "invalid request method",
		requestMethod:      get,
		responseType:       Core,
		responseStatusCode: http.StatusMethodNotAllowed,
		responseMessage:    methodNotAllowed,
		formerErr:          nil,
		formFirstName:      "",
		formLastName:       "",
		formEmail:          "",
		formPassword:       "",
		dbUserEmailExists:  mock.UserEmailExists{},
		dbUserCreate:       mock.UserCreate{},
	}),
	constructUserCreateTest(&userCreateTest{
		purpose:            "former error",
		requestMethod:      post,
		responseType:       Core,
		responseStatusCode: http.StatusBadRequest,
		responseMessage:    "former error",
		formerErr:          mockFormerErr,
		formFirstName:      "",
		formLastName:       "",
		formEmail:          "",
		formPassword:       "",
		dbUserEmailExists:  mock.UserEmailExists{},
		dbUserCreate:       mock.UserCreate{},
	}),
	constructUserCreateTest(&userCreateTest{
		purpose:            "missing form first_name",
		requestMethod:      post,
		responseType:       Core,
		responseStatusCode: http.StatusBadRequest,
		responseMessage:    fmt.Sprintf(formatMissing, firstName),
		formerErr:          nil,
		formFirstName:      "",
		formLastName:       mockLastName,
		formEmail:          mockEmail,
		formPassword:       mockPassword,
		dbUserEmailExists:  mock.UserEmailExists{},
		dbUserCreate:       mock.UserCreate{},
	}),
	constructUserCreateTest(&userCreateTest{
		purpose:            "missing form last_name",
		requestMethod:      post,
		responseType:       Core,
		responseStatusCode: http.StatusBadRequest,
		responseMessage:    fmt.Sprintf(formatMissing, lastName),
		formerErr:          nil,
		formFirstName:      mockFirstName,
		formLastName:       "",
		formEmail:          mockEmail,
		formPassword:       mockPassword,
		dbUserEmailExists:  mock.UserEmailExists{},
		dbUserCreate:       mock.UserCreate{},
	}),
	constructUserCreateTest(&userCreateTest{
		purpose:            "missing form email",
		requestMethod:      post,
		responseType:       Core,
		responseStatusCode: http.StatusBadRequest,
		responseMessage:    fmt.Sprintf(formatMissing, email),
		formerErr:          nil,
		formFirstName:      mockFirstName,
		formLastName:       mockLastName,
		formEmail:          "",
		formPassword:       mockPassword,
		dbUserEmailExists:  mock.UserEmailExists{},
		dbUserCreate:       mock.UserCreate{},
	}),
	constructUserCreateTest(&userCreateTest{
		purpose:            "missing form password",
		requestMethod:      post,
		responseType:       Core,
		responseStatusCode: http.StatusBadRequest,
		responseMessage:    fmt.Sprintf(formatMissing, password),
		formerErr:          nil,
		formFirstName:      mockFirstName,
		formLastName:       mockLastName,
		formEmail:          mockEmail,
		formPassword:       "",
		dbUserEmailExists:  mock.UserEmailExists{},
		dbUserCreate:       mock.UserCreate{},
	}),
	constructUserCreateTest(&userCreateTest{
		purpose:            "user email exists db error",
		requestMethod:      post,
		responseType:       Core,
		responseStatusCode: http.StatusInternalServerError,
		responseMessage:    friendlyErrorMessage,
		formerErr:          nil,
		formFirstName:      mockFirstName,
		formLastName:       mockLastName,
		formEmail:          mockEmail,
		formPassword:       mockPassword,
		dbUserEmailExists:  mock.UserEmailExists{Err: mockDbErr},
		dbUserCreate:       mock.UserCreate{},
	}),
	constructUserCreateTest(&userCreateTest{
		purpose:            "user email exists",
		requestMethod:      post,
		responseType:       Core,
		responseStatusCode: http.StatusBadRequest,
		responseMessage:    fmt.Sprintf(formatAlreadyExists, email),
		formerErr:          nil,
		formFirstName:      mockFirstName,
		formLastName:       mockLastName,
		formEmail:          mockEmail,
		formPassword:       mockPassword,
		dbUserEmailExists:  mock.UserEmailExists{Bool: true},
		dbUserCreate:       mock.UserCreate{},
	}),
	constructUserCreateTest(&userCreateTest{
		purpose:            "user create db error",
		requestMethod:      post,
		responseType:       Core,
		responseStatusCode: http.StatusInternalServerError,
		responseMessage:    friendlyErrorMessage,
		formerErr:          nil,
		formFirstName:      mockFirstName,
		formLastName:       mockLastName,
		formEmail:          mockEmail,
		formPassword:       mockPassword,
		dbUserEmailExists:  mock.UserEmailExists{Bool: false},
		dbUserCreate:       mock.UserCreate{Err: mockDbErr},
	}),
	constructUserCreateTest(&userCreateTest{
		purpose:            "user create ok",
		requestMethod:      post,
		responseType:       Retrieve,
		responseStatusCode: http.StatusCreated,
		responseMessage:    fmt.Sprintf(formatCreated, user),
		formerErr:          nil,
		formFirstName:      mockFirstName,
		formLastName:       mockLastName,
		formEmail:          mockEmail,
		formPassword:       mockPassword,
		dbUserEmailExists:  mock.UserEmailExists{Bool: false},
		dbUserCreate:       mock.UserCreate{ID: mockID},
	}),
}

func constructUserCreateTest(testInput *userCreateTest) *test {
	responseResults, _ := json.Marshal(mockID)

	db := &mock.DB{
		UserEmailExistsMock: testInput.dbUserEmailExists,
		UserCreateMock:      testInput.dbUserCreate,
	}

	logger := &mock.Logger{}

	former := &mock.Former{
		MapMock: mock.Map{Err: testInput.formerErr},
	}

	return &test{
		purpose: testInput.purpose,
		handle:  UserCreate,
		request: constructTestRequest(
			testInput.requestMethod,
			constructForm(map[string]string{
				firstName: testInput.formFirstName,
				lastName:  testInput.formLastName,
				email:     testInput.formEmail,
				password:  testInput.formPassword,
			}),
			constructEnv(db, logger, former, nil),
		),
		response: constructTestResponse(
			testInput.responseType,
			testInput.responseStatusCode,
			testInput.responseMessage,
			responseResults,
		),
	}
}
