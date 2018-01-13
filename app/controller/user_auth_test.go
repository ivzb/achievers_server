package controller

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/ivzb/achievers_server/app/model/mock"
	"github.com/ivzb/achievers_server/app/shared/consts"
)

func userAuthForm() *map[string]string {
	return &map[string]string{
		consts.Email:    mockEmail,
		consts.Password: mockPassword,
	}
}

var userAuthTests = []*test{
	constructUserAuthTest(&testInput{
		purpose:            "invalid request method",
		requestMethod:      consts.GET,
		responseType:       Core,
		responseStatusCode: http.StatusMethodNotAllowed,
		responseMessage:    consts.MethodNotAllowed,
		form:               &map[string]string{},
	}),
	constructUserAuthTest(&testInput{
		purpose:            "former error",
		requestMethod:      consts.POST,
		responseType:       Core,
		responseStatusCode: http.StatusBadRequest,
		responseMessage:    "content-type of request is incorrect",
		form:               &map[string]string{},
		removeHeaders:      true,
	}),
	constructUserAuthTest(&testInput{
		purpose:            "missing form consts.Email",
		requestMethod:      consts.POST,
		responseType:       Core,
		responseStatusCode: http.StatusBadRequest,
		responseMessage:    fmt.Sprintf(consts.FormatMissing, consts.Email),
		form:               mapWithout(userAuthForm(), consts.Email),
	}),
	constructUserAuthTest(&testInput{
		purpose:            "missing form consts.Password",
		requestMethod:      consts.POST,
		responseType:       Core,
		responseStatusCode: http.StatusBadRequest,
		responseMessage:    fmt.Sprintf(consts.FormatMissing, consts.Password),
		form:               mapWithout(userAuthForm(), consts.Password),
	}),
	constructUserAuthTest(&testInput{
		purpose:            "user consts.Email exists db error",
		requestMethod:      consts.POST,
		responseType:       Core,
		responseStatusCode: http.StatusInternalServerError,
		responseMessage:    consts.FriendlyErrorMessage,
		form:               userAuthForm(),
		db: &mock.DB{
			UserEmailExistsMock: mock.UserEmailExists{Err: mockDbErr},
		},
	}),
	constructUserAuthTest(&testInput{
		purpose:            "user consts.Email does not exist",
		requestMethod:      consts.POST,
		responseType:       Core,
		responseStatusCode: http.StatusNotFound,
		responseMessage:    fmt.Sprintf(consts.FormatNotFound, consts.Email),
		form:               userAuthForm(),
		db: &mock.DB{
			UserEmailExistsMock: mock.UserEmailExists{Bool: false},
		},
	}),
	constructUserAuthTest(&testInput{
		purpose:            "user auth db error",
		requestMethod:      consts.POST,
		responseType:       Core,
		responseStatusCode: http.StatusInternalServerError,
		responseMessage:    consts.FriendlyErrorMessage,
		form:               userAuthForm(),
		db: &mock.DB{
			UserEmailExistsMock: mock.UserEmailExists{Bool: true},
			UserAuthMock:        mock.UserAuth{Err: mockDbErr},
		},
	}),
	constructUserAuthTest(&testInput{
		purpose:            "encrypt error",
		requestMethod:      consts.POST,
		responseType:       Core,
		responseStatusCode: http.StatusInternalServerError,
		responseMessage:    consts.FriendlyErrorMessage,
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
		requestMethod:      consts.POST,
		responseType:       Retrieve,
		responseStatusCode: http.StatusCreated,
		responseMessage:    fmt.Sprintf(consts.FormatCreated, consts.AuthToken),
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
