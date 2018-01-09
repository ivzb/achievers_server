package controller

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/ivzb/achievers_server/app/model/mock"
)

func userCreateForm() *map[string]string {
	return &map[string]string{
		firstName: mockFirstName,
		lastName:  mockLastName,
		email:     mockEmail,
		password:  mockPassword,
	}
}

var userCreateTests = []*test{
	constructUserCreateTest(&testInput{
		purpose:            "invalid request method",
		requestMethod:      GET,
		responseType:       Core,
		responseStatusCode: http.StatusMethodNotAllowed,
		responseMessage:    methodNotAllowed,
		form:               &map[string]string{},
	}),
	constructUserCreateTest(&testInput{
		purpose:            "former error",
		requestMethod:      POST,
		responseType:       Core,
		responseStatusCode: http.StatusBadRequest,
		responseMessage:    "content-type of request is incorrect",
		form:               &map[string]string{},
		removeHeaders:      true,
	}),
	constructUserCreateTest(&testInput{
		purpose:            "missing form first_name",
		requestMethod:      POST,
		responseType:       Core,
		responseStatusCode: http.StatusBadRequest,
		responseMessage:    fmt.Sprintf(formatMissing, firstName),
		form:               mapWithout(userCreateForm(), firstName),
	}),
	constructUserCreateTest(&testInput{
		purpose:            "missing form last_name",
		requestMethod:      POST,
		responseType:       Core,
		responseStatusCode: http.StatusBadRequest,
		responseMessage:    fmt.Sprintf(formatMissing, lastName),
		form:               mapWithout(userCreateForm(), lastName),
	}),
	constructUserCreateTest(&testInput{
		purpose:            "missing form email",
		requestMethod:      POST,
		responseType:       Core,
		responseStatusCode: http.StatusBadRequest,
		responseMessage:    fmt.Sprintf(formatMissing, email),
		form:               mapWithout(userCreateForm(), email),
	}),
	constructUserCreateTest(&testInput{
		purpose:            "missing form password",
		requestMethod:      POST,
		responseType:       Core,
		responseStatusCode: http.StatusBadRequest,
		responseMessage:    fmt.Sprintf(formatMissing, password),
		form:               mapWithout(userCreateForm(), password),
	}),
	constructUserCreateTest(&testInput{
		purpose:            "user email exists db error",
		requestMethod:      POST,
		responseType:       Core,
		responseStatusCode: http.StatusInternalServerError,
		responseMessage:    friendlyErrorMessage,
		form:               userCreateForm(),
		db: &mock.DB{
			UserEmailExistsMock: mock.UserEmailExists{Err: mockDbErr},
		},
	}),
	constructUserCreateTest(&testInput{
		purpose:            "user email exists",
		requestMethod:      POST,
		responseType:       Core,
		responseStatusCode: http.StatusBadRequest,
		responseMessage:    fmt.Sprintf(formatAlreadyExists, email),
		form:               userCreateForm(),
		db: &mock.DB{
			UserEmailExistsMock: mock.UserEmailExists{Bool: true},
		},
	}),
	constructUserCreateTest(&testInput{
		purpose:            "user create db error",
		requestMethod:      POST,
		responseType:       Core,
		responseStatusCode: http.StatusInternalServerError,
		responseMessage:    friendlyErrorMessage,
		form:               userCreateForm(),
		db: &mock.DB{
			UserEmailExistsMock: mock.UserEmailExists{Bool: false},
			UserCreateMock:      mock.UserCreate{Err: mockDbErr},
		},
	}),
	constructUserCreateTest(&testInput{
		purpose:            "user create ok",
		requestMethod:      POST,
		responseType:       Retrieve,
		responseStatusCode: http.StatusCreated,
		responseMessage:    fmt.Sprintf(formatCreated, user),
		form:               userCreateForm(),
		db: &mock.DB{
			UserEmailExistsMock: mock.UserEmailExists{Bool: false},
			UserCreateMock:      mock.UserCreate{ID: mockID},
		},
	}),
}

func constructUserCreateTest(testInput *testInput) *test {
	responseResults, _ := json.Marshal(mockID)

	return constructTest(UserCreate, testInput, responseResults)
}
