package controller

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"net/http/httptest"
	"net/url"
	"strings"
	"testing"

	"github.com/ivzb/achievers_server/app/model"
	"github.com/ivzb/achievers_server/app/model/mock"
)

func TestUserAuth_ValidAuth(t *testing.T) {
	rec := httptest.NewRecorder()

	req, _ := http.NewRequest("POST", "/users/auth", nil)

	req.Form = url.Values{}
	req.Form.Add(email, mockEmail)
	req.Form.Add(password, mockPassword)

	env := model.Env{
		DB: &mock.DB{
			UserAuthMock: mock.UserAuth{ID: "454562", Err: nil},
		},
		Tokener: &mock.Token{
			EncryptedMock: mock.Encrypted{Enc: mockToken, Err: nil},
		},
	}

	handle := UserAuth
	statusCode := http.StatusCreated

	testHandler(t, rec, req, &env, handle, statusCode)

	message := authorized
	results, _ := json.Marshal(mockToken)

	expectRetrieve(t, rec, statusCode, message, results)
}

func TestUserAuth_InvalidMethod(t *testing.T) {
	testMethodNotAllowed(t, "GET", "/users/auth", UserAuth)
}

func TestUserAuth_MissingEmail(t *testing.T) {
	rec := httptest.NewRecorder()

	req, _ := http.NewRequest("POST", "/users/auth", nil)

	req.Form = url.Values{}
	req.Form.Add(password, "P@$$")

	handle := UserAuth
	statusCode := http.StatusBadRequest

	testHandler(t, rec, req, nil, handle, statusCode)

	message := fmt.Sprintf(formatMissing, email)
	expectCore(t, rec, statusCode, message)
}

func TestUserAuth_MissingPassword(t *testing.T) {
	rec := httptest.NewRecorder()

	req, _ := http.NewRequest("POST", "/users/auth", nil)

	req.Form = url.Values{}
	req.Form.Add(email, mockEmail)

	handle := UserAuth
	statusCode := http.StatusBadRequest

	testHandler(t, rec, req, nil, handle, statusCode)

	message := fmt.Sprintf(formatMissing, password)
	expectCore(t, rec, statusCode, message)
}

func TestUserAuth_DBError(t *testing.T) {
	rec := httptest.NewRecorder()

	req, _ := http.NewRequest("POST", "/users/auth", nil)

	req.Form = url.Values{}
	req.Form.Add(email, "Email@gmail.com")
	req.Form.Add(password, "P@$$")

	env := model.Env{
		DB: &mock.DB{
			UserAuthMock: mock.UserAuth{ID: "", Err: errors.New("db error")},
		},
	}

	handle := UserAuth
	statusCode := http.StatusUnauthorized

	testHandler(t, rec, req, &env, handle, statusCode)

	message := unauthorized
	expectCore(t, rec, statusCode, message)
}

func TestUserAuth_EncryptionError(t *testing.T) {
	rec := httptest.NewRecorder()

	req, _ := http.NewRequest("POST", "/users/auth", nil)

	req.Form = url.Values{}
	req.Form.Add(email, "Email@gmail.com")
	req.Form.Add(password, "P@$$")

	env := model.Env{
		DB: &mock.DB{
			UserAuthMock: mock.UserAuth{ID: "454562", Err: nil},
		},
		Tokener: &mock.Token{
			EncryptedMock: mock.Encrypted{Enc: "", Err: errors.New("encryption error")},
		},
	}

	handle := UserAuth
	statusCode := http.StatusInternalServerError

	testHandler(t, rec, req, &env, handle, statusCode)

	message := friendlyErrorMessage
	expectCore(t, rec, statusCode, message)
}

func TestUserCreate_ValidUser(t *testing.T) {
	rec := httptest.NewRecorder()

	form := url.Values{}
	form.Add(firstName, mockFirstName)
	form.Add(lastName, mockLastName)
	form.Add(email, mockEmail)
	form.Add(password, mockPassword)

	req, _ := http.NewRequest("POST", "/users/create", strings.NewReader(form.Encode()))

	req.Header = http.Header{}
	req.Header.Add("Content-Type", "application/x-www-form-urlencoded")

	env := model.Env{
		DB: &mock.DB{
			UserEmailExistsMock: mock.UserEmailExists{Bool: false, Err: nil},
			UserCreateMock:      mock.UserCreate{ID: mockID, Err: nil},
		},
		Former: &mock.Former{
			MapMock: mock.Map{Err: nil},
		},
	}

	handle := UserCreate
	statusCode := http.StatusCreated

	testHandler(t, rec, req, &env, handle, statusCode)

	message := fmt.Sprintf(formatCreated, user)
	results, _ := json.Marshal(mockID)
	expectRetrieve(t, rec, statusCode, message, results)
}

func TestUserCreate_InvalidMethod(t *testing.T) {
	testMethodNotAllowed(t, "GET", "/users/create", UserCreate)
}

func TestUserCreate_FormMapError(t *testing.T) {
	rec := httptest.NewRecorder()
	req, _ := http.NewRequest("POST", "/users/create", nil)

	req.Header = http.Header{}
	req.Header.Add("Content-Type", "application/x-www-form-urlencoded")

	handle := UserCreate
	statusCode := http.StatusBadRequest

	mapError := "map error"

	env := &model.Env{
		Former: &mock.Former{
			MapMock: mock.Map{Err: errors.New(mapError)},
		},
	}

	testHandler(t, rec, req, env, handle, statusCode)

	// Check the response body is what we expect.
	message := mapError
	expectCore(t, rec, statusCode, message)
}

func TestUserCreate_MissingFirstName(t *testing.T) {
	rec := httptest.NewRecorder()

	form := url.Values{}
	form.Add("first_name", "")
	form.Add("last_name", mockLastName)
	form.Add(email, mockEmail)
	form.Add(password, mockPassword)

	req, _ := http.NewRequest("POST", "/users/create", strings.NewReader(form.Encode()))

	req.Header = http.Header{}
	req.Header.Add("Content-Type", "application/x-www-form-urlencoded")

	env := &model.Env{
		Former: &mock.Former{
			MapMock: mock.Map{Err: nil},
		},
	}

	handle := UserCreate
	statusCode := http.StatusBadRequest

	testHandler(t, rec, req, env, handle, statusCode)

	message := fmt.Sprintf(formatMissing, firstName)
	expectCore(t, rec, statusCode, message)
}

func TestUserCreate_MissingLastName(t *testing.T) {
	rec := httptest.NewRecorder()

	form := url.Values{}
	form.Add("first_name", mockFirstName)
	form.Add("last_name", "")
	form.Add(email, mockEmail)
	form.Add(password, mockPassword)

	req, _ := http.NewRequest("POST", "/users/create", strings.NewReader(form.Encode()))

	req.Header = http.Header{}
	req.Header.Add("Content-Type", "application/x-www-form-urlencoded")

	env := &model.Env{
		Former: &mock.Former{
			MapMock: mock.Map{Err: nil},
		},
	}

	handle := UserCreate
	statusCode := http.StatusBadRequest

	testHandler(t, rec, req, env, handle, statusCode)

	message := fmt.Sprintf(formatMissing, lastName)
	expectCore(t, rec, statusCode, message)
}

func TestUserCreate_MissingEmail(t *testing.T) {
	rec := httptest.NewRecorder()

	form := url.Values{}
	form.Add("first_name", mockFirstName)
	form.Add("last_name", mockLastName)
	form.Add(email, "")
	form.Add(password, mockPassword)

	req, _ := http.NewRequest("POST", "/users/create", strings.NewReader(form.Encode()))

	req.Header = http.Header{}
	req.Header.Add("Content-Type", "application/x-www-form-urlencoded")

	env := &model.Env{
		Former: &mock.Former{
			MapMock: mock.Map{Err: nil},
		},
	}

	handle := UserCreate
	statusCode := http.StatusBadRequest

	testHandler(t, rec, req, env, handle, statusCode)

	message := fmt.Sprintf(formatMissing, email)
	expectCore(t, rec, statusCode, message)
}

func TestUserCreate_MissingPassword(t *testing.T) {
	rec := httptest.NewRecorder()

	form := url.Values{}
	form.Add("first_name", mockFirstName)
	form.Add("last_name", mockLastName)
	form.Add(email, mockEmail)
	form.Add(password, "")

	req, _ := http.NewRequest("POST", "/users/create", strings.NewReader(form.Encode()))

	req.Header = http.Header{}
	req.Header.Add("Content-Type", "application/x-www-form-urlencoded")

	env := &model.Env{
		Former: &mock.Former{
			MapMock: mock.Map{Err: nil},
		},
	}

	handle := UserCreate
	statusCode := http.StatusBadRequest

	testHandler(t, rec, req, env, handle, statusCode)

	message := fmt.Sprintf(formatMissing, password)
	expectCore(t, rec, statusCode, message)
}

func TestUserCreate_ExistDBError(t *testing.T) {
	rec := httptest.NewRecorder()

	form := url.Values{}
	form.Add("first_name", mockFirstName)
	form.Add("last_name", mockLastName)
	form.Add(email, mockEmail)
	form.Add(password, mockPassword)

	req, _ := http.NewRequest("POST", "/users/create", strings.NewReader(form.Encode()))

	req.Header = http.Header{}
	req.Header.Add("Content-Type", "application/x-www-form-urlencoded")

	env := model.Env{
		DB: &mock.DB{
			UserEmailExistsMock: mock.UserEmailExists{Bool: false, Err: errors.New("db error")},
		},
		Former: &mock.Former{
			MapMock: mock.Map{Err: nil},
		},
	}

	handle := UserCreate
	statusCode := http.StatusInternalServerError

	testHandler(t, rec, req, &env, handle, statusCode)

	message := friendlyErrorMessage
	expectCore(t, rec, statusCode, message)
}

func TestUserCreate_UserExist(t *testing.T) {
	rec := httptest.NewRecorder()

	form := url.Values{}
	form.Add("first_name", mockFirstName)
	form.Add("last_name", mockLastName)
	form.Add(email, mockEmail)
	form.Add(password, mockPassword)

	req, _ := http.NewRequest("POST", "/users/create", strings.NewReader(form.Encode()))

	req.Header = http.Header{}
	req.Header.Add("Content-Type", "application/x-www-form-urlencoded")

	env := model.Env{
		DB: &mock.DB{
			UserEmailExistsMock: mock.UserEmailExists{Bool: true, Err: nil},
		},
		Former: &mock.Former{
			MapMock: mock.Map{Err: nil},
		},
	}

	handle := UserCreate
	statusCode := http.StatusBadRequest

	testHandler(t, rec, req, &env, handle, statusCode)

	message := fmt.Sprintf(formatAlreadyExists, email)
	expectCore(t, rec, statusCode, message)
}

func TestUserCreate_UserCreateDBError(t *testing.T) {
	rec := httptest.NewRecorder()

	form := url.Values{}
	form.Add("first_name", mockFirstName)
	form.Add("last_name", mockLastName)
	form.Add(email, mockEmail)
	form.Add(password, mockPassword)

	req, _ := http.NewRequest("POST", "/users/create", strings.NewReader(form.Encode()))

	req.Header = http.Header{}
	req.Header.Add("Content-Type", "application/x-www-form-urlencoded")

	env := model.Env{
		DB: &mock.DB{
			UserEmailExistsMock: mock.UserEmailExists{Bool: false, Err: nil},
			UserCreateMock:      mock.UserCreate{ID: "", Err: errors.New("db error")},
		},
		Former: &mock.Former{
			MapMock: mock.Map{Err: nil},
		},
	}

	handle := UserCreate
	statusCode := http.StatusInternalServerError

	testHandler(t, rec, req, &env, handle, statusCode)

	message := friendlyErrorMessage
	expectCore(t, rec, statusCode, message)
}
