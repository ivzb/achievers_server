package controller

import (
	"errors"
	"net/http"
	"net/http/httptest"
	"net/url"
	"strconv"
	"testing"

	"github.com/ivzb/achievers_server/app/model"
)

var (
	mockId        = "42"
	mockFirstName = "John"
	mockLastName  = "Doe"
	mockEmail     = "email@gmail.com"
	mockPassword  = "P@$$"
	mockToken     = "34567899876543"
)

func TestUserAuth_ValidAuth(t *testing.T) {
	rec := httptest.NewRecorder()

	req, _ := http.NewRequest("POST", "/users/auth", nil)

	req.Form = url.Values{}
	req.Form.Add("email", mockEmail)
	req.Form.Add("password", mockPassword)

	env := model.Env{
		DB: &model.DBMock{
			UserAuthMock: model.UserAuthMock{"454562", nil},
		},
		Tokener: &model.TokenerMock{
			EncryptedMock: model.EncryptedMock{mockToken, nil},
		},
	}

	handle := UserAuth
	statusCode := http.StatusOK

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
	expected := `{"status":` + strconv.Itoa(statusCode) + `,"message":"` + MissingEmailErrorMessage + `"}`
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
	expected := `{"status":` + strconv.Itoa(statusCode) + `,"message":"` + MissingPasswordErrorMessage + `"}`
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
		DB: &model.DBMock{
			UserAuthMock: model.UserAuthMock{"", errors.New("db error")},
		},
	}

	handle := UserAuth
	statusCode := http.StatusUnauthorized

	testHandler(t, rec, req, &env, handle, statusCode)

	// Check the response body is what we expect.
	expected := `{"status":` + strconv.Itoa(statusCode) + `,"message":"` + Unauthorized + `"}`
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
		DB: &model.DBMock{
			UserAuthMock: model.UserAuthMock{"454562", nil},
		},
		Tokener: &model.TokenerMock{
			EncryptedMock: model.EncryptedMock{"", errors.New("encryption error")},
		},
	}

	handle := UserAuth
	statusCode := http.StatusInternalServerError

	testHandler(t, rec, req, &env, handle, statusCode)

	// Check the response body is what we expect.
	expected := `{"status":` + strconv.Itoa(statusCode) + `,"message":"` + FriendlyErrorMessage + `"}`
	if rec.Body.String() != expected {
		t.Errorf("handler returned unexpected body: got %v want %v",
			rec.Body.String(), expected)
	}
}

func TestUserCreate_ValidUser(t *testing.T) {
	rec := httptest.NewRecorder()

	req, _ := http.NewRequest("POST", "/users/create", nil)

	req.Form = url.Values{}
	req.Form.Add("first_name", mockFirstName)
	req.Form.Add("last_name", mockLastName)
	req.Form.Add("email", mockEmail)
	req.Form.Add("password", mockPassword)

	env := model.Env{
		DB: &model.DBMock{
			ExistsMock:     model.ExistsMock{false, nil},
			UserCreateMock: model.UserCreateMock{mockId, nil},
		},
	}

	handle := UserCreate
	statusCode := http.StatusCreated

	testHandler(t, rec, req, &env, handle, statusCode)

	// Check the response body is what we expect.
	expected := `{"status":` + strconv.Itoa(statusCode) + `,"message":"` + UserCreated + `","length":1,"results":"` + mockId + `"}`
	if rec.Body.String() != expected {
		t.Errorf("handler returned unexpected body: got %v want %v",
			rec.Body.String(), expected)
	}
}

func TestUserCreate_InvalidMethod(t *testing.T) {
	testInvalidMethod(t, "GET", "/users/create", UserCreate)
}

func TestUserCreate_MissingFirstName(t *testing.T) {
	rec := httptest.NewRecorder()

	req, _ := http.NewRequest("POST", "/users/create", nil)

	req.Form = url.Values{}
	req.Form.Add("last_name", mockLastName)
	req.Form.Add("email", mockEmail)
	req.Form.Add("password", mockPassword)

	handle := UserCreate
	statusCode := http.StatusBadRequest

	testHandler(t, rec, req, nil, handle, statusCode)

	// Check the response body is what we expect.
	expected := `{"status":` + strconv.Itoa(statusCode) + `,"message":"` + MissingFirstNameErrorMessage + `"}`
	if rec.Body.String() != expected {
		t.Errorf("handler returned unexpected body: got %v want %v",
			rec.Body.String(), expected)
	}
}

func TestUserCreate_MissingLastName(t *testing.T) {
	rec := httptest.NewRecorder()

	req, _ := http.NewRequest("POST", "/users/create", nil)

	req.Form = url.Values{}
	req.Form.Add("first_name", mockFirstName)
	req.Form.Add("email", mockEmail)
	req.Form.Add("password", mockPassword)

	handle := UserCreate
	statusCode := http.StatusBadRequest

	testHandler(t, rec, req, nil, handle, statusCode)

	// Check the response body is what we expect.
	expected := `{"status":` + strconv.Itoa(statusCode) + `,"message":"` + MissingLastNameErrorMessage + `"}`
	if rec.Body.String() != expected {
		t.Errorf("handler returned unexpected body: got %v want %v",
			rec.Body.String(), expected)
	}
}

func TestUserCreate_MissingEmail(t *testing.T) {
	rec := httptest.NewRecorder()

	req, _ := http.NewRequest("POST", "/users/create", nil)

	req.Form = url.Values{}
	req.Form.Add("first_name", mockFirstName)
	req.Form.Add("last_name", mockLastName)
	req.Form.Add("password", mockPassword)

	handle := UserCreate
	statusCode := http.StatusBadRequest

	testHandler(t, rec, req, nil, handle, statusCode)

	// Check the response body is what we expect.
	expected := `{"status":` + strconv.Itoa(statusCode) + `,"message":"` + MissingEmailErrorMessage + `"}`
	if rec.Body.String() != expected {
		t.Errorf("handler returned unexpected body: got %v want %v",
			rec.Body.String(), expected)
	}
}

func TestUserCreate_MissingPassword(t *testing.T) {
	rec := httptest.NewRecorder()

	req, _ := http.NewRequest("POST", "/users/create", nil)

	req.Form = url.Values{}
	req.Form.Add("first_name", mockFirstName)
	req.Form.Add("last_name", mockLastName)
	req.Form.Add("email", mockEmail)

	handle := UserCreate
	statusCode := http.StatusBadRequest

	testHandler(t, rec, req, nil, handle, statusCode)

	// Check the response body is what we expect.
	expected := `{"status":` + strconv.Itoa(statusCode) + `,"message":"` + MissingPasswordErrorMessage + `"}`
	if rec.Body.String() != expected {
		t.Errorf("handler returned unexpected body: got %v want %v",
			rec.Body.String(), expected)
	}
}

func TestUserCreate_ExistDBError(t *testing.T) {
	rec := httptest.NewRecorder()

	req, _ := http.NewRequest("POST", "/users/create", nil)

	req.Form = url.Values{}
	req.Form.Add("first_name", mockFirstName)
	req.Form.Add("last_name", mockLastName)
	req.Form.Add("email", mockEmail)
	req.Form.Add("password", mockPassword)

	env := model.Env{
		DB: &model.DBMock{
			ExistsMock: model.ExistsMock{false, errors.New("db error")},
		},
	}

	handle := UserCreate
	statusCode := http.StatusInternalServerError

	testHandler(t, rec, req, &env, handle, statusCode)

	// Check the response body is what we expect.
	expected := `{"status":` + strconv.Itoa(statusCode) + `,"message":"` + FriendlyErrorMessage + `"}`
	if rec.Body.String() != expected {
		t.Errorf("handler returned unexpected body: got %v want %v",
			rec.Body.String(), expected)
	}
}

func TestUserCreate_UserExist(t *testing.T) {
	rec := httptest.NewRecorder()

	req, _ := http.NewRequest("POST", "/users/create", nil)

	req.Form = url.Values{}
	req.Form.Add("first_name", mockFirstName)
	req.Form.Add("last_name", mockLastName)
	req.Form.Add("email", mockEmail)
	req.Form.Add("password", mockPassword)

	env := model.Env{
		DB: &model.DBMock{
			ExistsMock: model.ExistsMock{true, nil},
		},
	}

	handle := UserCreate
	statusCode := http.StatusBadRequest

	testHandler(t, rec, req, &env, handle, statusCode)

	// Check the response body is what we expect.
	expected := `{"status":` + strconv.Itoa(statusCode) + `,"message":"` + EmailAlreadyExistsErrorMessage + `"}`
	if rec.Body.String() != expected {
		t.Errorf("handler returned unexpected body: got %v want %v",
			rec.Body.String(), expected)
	}
}

func TestUserCreate_UserCreateDBError(t *testing.T) {
	rec := httptest.NewRecorder()

	req, _ := http.NewRequest("POST", "/users/create", nil)

	req.Form = url.Values{}
	req.Form.Add("first_name", mockFirstName)
	req.Form.Add("last_name", mockLastName)
	req.Form.Add("email", mockEmail)
	req.Form.Add("password", mockPassword)

	env := model.Env{
		DB: &model.DBMock{
			ExistsMock:     model.ExistsMock{false, nil},
			UserCreateMock: model.UserCreateMock{"", errors.New("db error")},
		},
	}

	handle := UserCreate
	statusCode := http.StatusInternalServerError

	testHandler(t, rec, req, &env, handle, statusCode)

	// Check the response body is what we expect.
	expected := `{"status":` + strconv.Itoa(statusCode) + `,"message":"` + FriendlyErrorMessage + `"}`
	if rec.Body.String() != expected {
		t.Errorf("handler returned unexpected body: got %v want %v",
			rec.Body.String(), expected)
	}
}
