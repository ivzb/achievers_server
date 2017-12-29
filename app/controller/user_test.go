package controller

import (
	"errors"
	"fmt"
	"net/http"
	"net/http/httptest"
	"net/url"
	"strconv"
	"strings"
	"testing"

	"github.com/ivzb/achievers_server/app/model"
	"github.com/ivzb/achievers_server/app/model/mock"
)

func TestUserAuth_ValidAuth(t *testing.T) {
	rec := httptest.NewRecorder()

	req, _ := http.NewRequest("POST", "/users/auth", nil)

	req.Form = url.Values{}
	req.Form.Add("email", mockEmail)
	req.Form.Add("password", mockPassword)

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

	// Check the response body is what we expect.
	expected := `{"status":` + strconv.Itoa(statusCode) + `,"message":"authorized","length":1,"results":"` + mockToken + `"}`
	if rec.Body.String() != expected {
		t.Errorf("handler returned unexpected body: got %v want %v",
			rec.Body.String(), expected)
	}
}

func TestUserAuth_InvalidMethod(t *testing.T) {
	testInvalidMethod(t, "GET", "/users/auth", UserAuth)
}

func TestUserAuth_MissingEmail(t *testing.T) {
	rec := httptest.NewRecorder()

	req, _ := http.NewRequest("POST", "/users/auth", nil)

	req.Form = url.Values{}
	req.Form.Add("password", "P@$$")

	handle := UserAuth
	statusCode := http.StatusBadRequest

	testHandler(t, rec, req, nil, handle, statusCode)

	// Check the response body is what we expect.
	expected := `{"status":` + strconv.Itoa(statusCode) + `,"message":"` + fmt.Sprintf(formatMissing, email) + `"}`
	if rec.Body.String() != expected {
		t.Errorf("handler returned unexpected body: got %v want %v",
			rec.Body.String(), expected)
	}
}

func TestUserAuth_MissingPassword(t *testing.T) {
	rec := httptest.NewRecorder()

	req, _ := http.NewRequest("POST", "/users/auth", nil)

	req.Form = url.Values{}
	req.Form.Add("email", "email@gmail.com")

	handle := UserAuth
	statusCode := http.StatusBadRequest

	testHandler(t, rec, req, nil, handle, statusCode)

	// Check the response body is what we expect.
	expected := `{"status":` + strconv.Itoa(statusCode) + `,"message":"` + fmt.Sprintf(formatMissing, password) + `"}`
	if rec.Body.String() != expected {
		t.Errorf("handler returned unexpected body: got %v want %v",
			rec.Body.String(), expected)
	}
}

func TestUserAuth_DBError(t *testing.T) {
	rec := httptest.NewRecorder()

	req, _ := http.NewRequest("POST", "/users/auth", nil)

	req.Form = url.Values{}
	req.Form.Add("email", "email@gmail.com")
	req.Form.Add("password", "P@$$")

	env := model.Env{
		DB: &mock.DB{
			UserAuthMock: mock.UserAuth{ID: "", Err: errors.New("db error")},
		},
	}

	handle := UserAuth
	statusCode := http.StatusUnauthorized

	testHandler(t, rec, req, &env, handle, statusCode)

	// Check the response body is what we expect.
	expected := `{"status":` + strconv.Itoa(statusCode) + `,"message":"` + unauthorized + `"}`
	if rec.Body.String() != expected {
		t.Errorf("handler returned unexpected body: got %v want %v",
			rec.Body.String(), expected)
	}
}

func TestUserAuth_EncryptionError(t *testing.T) {
	rec := httptest.NewRecorder()

	req, _ := http.NewRequest("POST", "/users/auth", nil)

	req.Form = url.Values{}
	req.Form.Add("email", "email@gmail.com")
	req.Form.Add("password", "P@$$")

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

	// Check the response body is what we expect.
	expected := `{"status":` + strconv.Itoa(statusCode) + `,"message":"` + friendlyErrorMessage + `"}`
	if rec.Body.String() != expected {
		t.Errorf("handler returned unexpected body: got %v want %v",
			rec.Body.String(), expected)
	}
}

func TestUserCreate_ValidUser(t *testing.T) {
	rec := httptest.NewRecorder()

	form := url.Values{}
	form.Add("first_name", mockFirstName)
	form.Add("last_name", mockLastName)
	form.Add("email", mockEmail)
	form.Add("password", mockPassword)

	req, _ := http.NewRequest("POST", "/users/create", strings.NewReader(form.Encode()))

	req.Header = http.Header{}
	req.Header.Add("Content-Type", "application/x-www-form-urlencoded")

	env := model.Env{
		DB: &mock.DB{
			UserExistsMock: mock.UserExists{Bool: false, Err: nil},
			UserCreateMock: mock.UserCreate{ID: mockID, Err: nil},
		},
		Former: &mock.Former{
			MapMock: mock.Map{Err: nil},
		},
	}

	handle := UserCreate
	statusCode := http.StatusCreated

	testHandler(t, rec, req, &env, handle, statusCode)

	// Check the response body is what we expect.
	expected := `{"status":` + strconv.Itoa(statusCode) + `,"message":"` + fmt.Sprintf(formatCreated, user) + `","length":1,"results":"` + mockID + `"}`
	if rec.Body.String() != expected {
		t.Errorf("handler returned unexpected body: got %v want %v",
			rec.Body.String(), expected)
	}
}

func TestUserCreate_InvalidMethod(t *testing.T) {
	testInvalidMethod(t, "GET", "/users/create", UserCreate)
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
	expectedMessage := mapError
	expected := fmt.Sprintf(`{"status":%d,"message":"%s"}`, statusCode, expectedMessage)
	if rec.Body.String() != expected {
		t.Errorf("handler returned unexpected body: got %v want %v",
			rec.Body.String(), expected)
	}

}

func TestUserCreate_MissingFirstName(t *testing.T) {
	rec := httptest.NewRecorder()

	form := url.Values{}
	form.Add("first_name", "")
	form.Add("last_name", mockLastName)
	form.Add("email", mockEmail)
	form.Add("password", mockPassword)

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

	// Check the response body is what we expect.
	expected := `{"status":` + strconv.Itoa(statusCode) + `,"message":"` + fmt.Sprintf(formatMissing, firstName) + `"}`
	if rec.Body.String() != expected {
		t.Errorf("handler returned unexpected body: got %v want %v",
			rec.Body.String(), expected)
	}
}

func TestUserCreate_MissingLastName(t *testing.T) {
	rec := httptest.NewRecorder()

	form := url.Values{}
	form.Add("first_name", mockFirstName)
	form.Add("last_name", "")
	form.Add("email", mockEmail)
	form.Add("password", mockPassword)

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

	// Check the response body is what we expect.
	expected := `{"status":` + strconv.Itoa(statusCode) + `,"message":"` + fmt.Sprintf(formatMissing, lastName) + `"}`
	if rec.Body.String() != expected {
		t.Errorf("handler returned unexpected body: got %v want %v",
			rec.Body.String(), expected)
	}
}

func TestUserCreate_MissingEmail(t *testing.T) {
	rec := httptest.NewRecorder()

	form := url.Values{}
	form.Add("first_name", mockFirstName)
	form.Add("last_name", mockLastName)
	form.Add("email", "")
	form.Add("password", mockPassword)

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

	// Check the response body is what we expect.
	expected := `{"status":` + strconv.Itoa(statusCode) + `,"message":"` + fmt.Sprintf(formatMissing, email) + `"}`
	if rec.Body.String() != expected {
		t.Errorf("handler returned unexpected body: got %v want %v",
			rec.Body.String(), expected)
	}
}

func TestUserCreate_MissingPassword(t *testing.T) {
	rec := httptest.NewRecorder()

	form := url.Values{}
	form.Add("first_name", mockFirstName)
	form.Add("last_name", mockLastName)
	form.Add("email", mockEmail)
	form.Add("password", "")

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

	// Check the response body is what we expect.
	expected := `{"status":` + strconv.Itoa(statusCode) + `,"message":"` + fmt.Sprintf(formatMissing, password) + `"}`
	if rec.Body.String() != expected {
		t.Errorf("handler returned unexpected body: got %v want %v",
			rec.Body.String(), expected)
	}
}

func TestUserCreate_ExistDBError(t *testing.T) {
	rec := httptest.NewRecorder()

	form := url.Values{}
	form.Add("first_name", mockFirstName)
	form.Add("last_name", mockLastName)
	form.Add("email", mockEmail)
	form.Add("password", mockPassword)

	req, _ := http.NewRequest("POST", "/users/create", strings.NewReader(form.Encode()))

	req.Header = http.Header{}
	req.Header.Add("Content-Type", "application/x-www-form-urlencoded")

	env := model.Env{
		DB: &mock.DB{
			UserExistsMock: mock.UserExists{Bool: false, Err: errors.New("db error")},
		},
		Former: &mock.Former{
			MapMock: mock.Map{Err: nil},
		},
	}

	handle := UserCreate
	statusCode := http.StatusInternalServerError

	testHandler(t, rec, req, &env, handle, statusCode)

	// Check the response body is what we expect.
	expected := `{"status":` + strconv.Itoa(statusCode) + `,"message":"` + friendlyErrorMessage + `"}`
	if rec.Body.String() != expected {
		t.Errorf("handler returned unexpected body: got %v want %v",
			rec.Body.String(), expected)
	}
}

func TestUserCreate_UserExist(t *testing.T) {
	rec := httptest.NewRecorder()

	form := url.Values{}
	form.Add("first_name", mockFirstName)
	form.Add("last_name", mockLastName)
	form.Add("email", mockEmail)
	form.Add("password", mockPassword)

	req, _ := http.NewRequest("POST", "/users/create", strings.NewReader(form.Encode()))

	req.Header = http.Header{}
	req.Header.Add("Content-Type", "application/x-www-form-urlencoded")

	env := model.Env{
		DB: &mock.DB{
			UserExistsMock: mock.UserExists{Bool: true, Err: nil},
		},
		Former: &mock.Former{
			MapMock: mock.Map{Err: nil},
		},
	}

	handle := UserCreate
	statusCode := http.StatusBadRequest

	testHandler(t, rec, req, &env, handle, statusCode)

	// Check the response body is what we expect.
	expected := `{"status":` + strconv.Itoa(statusCode) + `,"message":"` + fmt.Sprintf(formatAlreadyExists, email) + `"}`
	if rec.Body.String() != expected {
		t.Errorf("handler returned unexpected body: got %v want %v",
			rec.Body.String(), expected)
	}
}

func TestUserCreate_UserCreateDBError(t *testing.T) {
	rec := httptest.NewRecorder()

	form := url.Values{}
	form.Add("first_name", mockFirstName)
	form.Add("last_name", mockLastName)
	form.Add("email", mockEmail)
	form.Add("password", mockPassword)

	req, _ := http.NewRequest("POST", "/users/create", strings.NewReader(form.Encode()))

	req.Header = http.Header{}
	req.Header.Add("Content-Type", "application/x-www-form-urlencoded")

	env := model.Env{
		DB: &mock.DB{
			UserExistsMock: mock.UserExists{Bool: false, Err: nil},
			UserCreateMock: mock.UserCreate{ID: "", Err: errors.New("db error")},
		},
		Former: &mock.Former{
			MapMock: mock.Map{Err: nil},
		},
	}

	handle := UserCreate
	statusCode := http.StatusInternalServerError

	testHandler(t, rec, req, &env, handle, statusCode)

	// Check the response body is what we expect.
	expected := `{"status":` + strconv.Itoa(statusCode) + `,"message":"` + friendlyErrorMessage + `"}`
	if rec.Body.String() != expected {
		t.Errorf("handler returned unexpected body: got %v want %v",
			rec.Body.String(), expected)
	}
}
