package controller

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/ivzb/achievers_server/app/model/mock"
	"github.com/ivzb/achievers_server/app/shared/consts"
)

func userCreateForm() *map[string]string {
	return &map[string]string{
		consts.FirstName: mockFirstName,
		consts.LastName:  mockLastName,
		consts.Email:     mockEmail,
		consts.Password:  mockPassword,
	}
}

var userCreateTests = []*test{
	constructUserCreateTest(&testInput{
		purpose:            "invalid request method",
		requestMethod:      consts.GET,
		responseType:       Core,
		responseStatusCode: http.StatusMethodNotAllowed,
		responseMessage:    consts.MethodNotAllowed,
		form:               &map[string]string{},
	}),
	constructUserCreateTest(&testInput{
		purpose:            "former error",
		requestMethod:      consts.POST,
		responseType:       Core,
		responseStatusCode: http.StatusBadRequest,
		responseMessage:    "content-type of request is incorrect",
		form:               &map[string]string{},
		removeHeaders:      true,
	}),
	constructUserCreateTest(&testInput{
		purpose:            "missing form first_name",
		requestMethod:      consts.POST,
		responseType:       Core,
		responseStatusCode: http.StatusBadRequest,
		responseMessage:    fmt.Sprintf(consts.FormatMissing, consts.FirstName),
		form:               mapWithout(userCreateForm(), consts.FirstName),
	}),
	constructUserCreateTest(&testInput{
		purpose:            "missing form last_name",
		requestMethod:      consts.POST,
		responseType:       Core,
		responseStatusCode: http.StatusBadRequest,
		responseMessage:    fmt.Sprintf(consts.FormatMissing, consts.LastName),
		form:               mapWithout(userCreateForm(), consts.LastName),
	}),
	constructUserCreateTest(&testInput{
		purpose:            "missing form consts.Email",
		requestMethod:      consts.POST,
		responseType:       Core,
		responseStatusCode: http.StatusBadRequest,
		responseMessage:    fmt.Sprintf(consts.FormatMissing, consts.Email),
		form:               mapWithout(userCreateForm(), consts.Email),
	}),
	constructUserCreateTest(&testInput{
		purpose:            "missing form consts.Password",
		requestMethod:      consts.POST,
		responseType:       Core,
		responseStatusCode: http.StatusBadRequest,
		responseMessage:    fmt.Sprintf(consts.FormatMissing, consts.Password),
		form:               mapWithout(userCreateForm(), consts.Password),
	}),
	constructUserCreateTest(&testInput{
		purpose:            "user consts.Email exists db error",
		requestMethod:      consts.POST,
		responseType:       Core,
		responseStatusCode: http.StatusInternalServerError,
		responseMessage:    consts.FriendlyErrorMessage,
		form:               userCreateForm(),
		db: &mock.DB{
			UserEmailExistsMock: mock.UserEmailExists{Err: mockDbErr},
		},
	}),
	constructUserCreateTest(&testInput{
		purpose:            "user consts.Email exists",
		requestMethod:      consts.POST,
		responseType:       Core,
		responseStatusCode: http.StatusBadRequest,
		responseMessage:    fmt.Sprintf(consts.FormatAlreadyExists, consts.Email),
		form:               userCreateForm(),
		db: &mock.DB{
			UserEmailExistsMock: mock.UserEmailExists{Bool: true},
		},
	}),
	constructUserCreateTest(&testInput{
		purpose:            "user create db error",
		requestMethod:      consts.POST,
		responseType:       Core,
		responseStatusCode: http.StatusInternalServerError,
		responseMessage:    consts.FriendlyErrorMessage,
		form:               userCreateForm(),
		db: &mock.DB{
			UserEmailExistsMock: mock.UserEmailExists{Bool: false},
			UserCreateMock:      mock.UserCreate{Err: mockDbErr},
		},
	}),
	constructUserCreateTest(&testInput{
		purpose:            "user create ok",
		requestMethod:      consts.POST,
		responseType:       Retrieve,
		responseStatusCode: http.StatusCreated,
		responseMessage:    fmt.Sprintf(consts.FormatCreated, consts.User),
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
