package auth

import (
	"errors"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/ivzb/achievers_server/app/middleware/app"
	"github.com/ivzb/achievers_server/app/model"
	"github.com/ivzb/achievers_server/app/shared/response"
)

func testHandler(env *model.Env, w http.ResponseWriter, r *http.Request) response.Message {
	return response.Created("authorized", "auth token here")
}

func TestAuthHandler_ValidAuthToken(t *testing.T) {
	req := httptest.NewRequest("GET", "/auth", nil)
	req.Header.Add("auth_token", "asdf")

	rr := httptest.NewRecorder()

	env := &model.Env{
		DB: &model.DBMock{
			ExistsMock: model.ExistsMock{true, nil},
		},
		Tokener: &model.TokenMock{
			DecryptedMock: model.DecryptedMock{"decrypted", nil},
		},
	}

	appHandler := app.Handler{env, testHandler}

	var handler http.Handler = Handler(appHandler)

	handler.ServeHTTP(rr, req)

	// Check the status code is what we expect.
	if status := rr.Code; status != http.StatusCreated {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusCreated)
	}

	// Check the response body is what we expect.
	expected := `{"status":201,"message":"authorized","length":1,"results":"auth token here"}`
	if rr.Body.String() != expected {
		t.Errorf("handler returned unexpected body: \ngot %v \nwant %v",
			rr.Body.String(), expected)
	}
}

func TestAuthHandler_MissingAuthToken(t *testing.T) {
	req := httptest.NewRequest("GET", "/auth", nil)

	rr := httptest.NewRecorder()

	env := &model.Env{
		DB:      &model.DBMock{},
		Tokener: &model.TokenMock{},
	}

	appHandler := app.Handler{env, testHandler}

	var handler http.Handler = Handler(appHandler)

	handler.ServeHTTP(rr, req)

	// Check the status code is what we expect.
	if status := rr.Code; status != http.StatusUnauthorized {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusUnauthorized)
	}

	// Check the response body is what we expect.
	expected := `{"status":401,"message":"auth_token is missing"}`
	if rr.Body.String() != expected {
		t.Errorf("handler returned unexpected body: got %v want %v",
			rr.Body.String(), expected)
	}
}

func TestAuthHandler_InvalidAuthToken(t *testing.T) {
	req := httptest.NewRequest("GET", "/auth", nil)
	req.Header.Add("auth_token", "asdf")

	rec := httptest.NewRecorder()

	env := &model.Env{
		DB: &model.DBMock{},
		Tokener: &model.TokenMock{
			DecryptedMock: model.DecryptedMock{"", errors.New("decryption error")},
		},
	}

	appHandler := app.Handler{env, testHandler}

	var handler http.Handler = Handler(appHandler)

	handler.ServeHTTP(rec, req)

	// Check the status code is what we expect.
	if status := rec.Code; status != http.StatusUnauthorized {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusUnauthorized)
	}

	// Check the response body is what we expect.
	expected := `{"status":401,"message":"auth_token is invalid"}`
	if rec.Body.String() != expected {
		t.Errorf("handler returned unexpected body: got %v want %v",
			rec.Body.String(), expected)
	}
}

func TestAuthHandler_DBError(t *testing.T) {
	req := httptest.NewRequest("GET", "/auth", nil)
	req.Header.Add("auth_token", "asdf")

	rec := httptest.NewRecorder()

	env := &model.Env{
		DB: &model.DBMock{
			ExistsMock: model.ExistsMock{false, errors.New("user does not exist")},
		},
		Tokener: &model.TokenMock{
			DecryptedMock: model.DecryptedMock{"decrypted", nil},
		},
	}

	appHandler := app.Handler{env, testHandler}

	var handler http.Handler = Handler(appHandler)

	handler.ServeHTTP(rec, req)

	// Check the status code is what we expect.
	if status := rec.Code; status != http.StatusUnauthorized {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusUnauthorized)
	}

	// Check the response body is what we expect.
	expected := `{"status":401,"message":"auth_token is invalid"}`
	if rec.Body.String() != expected {
		t.Errorf("handler returned unexpected body: got %v want %v",
			rec.Body.String(), expected)
	}
}

func TestAuthHandler_UserDoesNotExist(t *testing.T) {
	req := httptest.NewRequest("GET", "/auth", nil)
	req.Header.Add("auth_token", "asdf")

	rec := httptest.NewRecorder()

	env := &model.Env{
		DB: &model.DBMock{
			ExistsMock: model.ExistsMock{false, nil},
		},
		Tokener: &model.TokenMock{
			DecryptedMock: model.DecryptedMock{"decrypted", nil},
		},
	}

	appHandler := app.Handler{env, testHandler}

	var handler http.Handler = Handler(appHandler)

	handler.ServeHTTP(rec, req)

	// Check the status code is what we expect.
	if status := rec.Code; status != http.StatusUnauthorized {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusUnauthorized)
	}

	// Check the response body is what we expect.
	expected := `{"status":401,"message":"auth_token is invalid"}`
	if rec.Body.String() != expected {
		t.Errorf("handler returned unexpected body: got %v want %v",
			rec.Body.String(), expected)
	}
}
