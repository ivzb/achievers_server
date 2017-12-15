package auth

import (
	"app/middleware/app"
	"app/model"
	"app/shared/response"
	"errors"
	"net/http"
	"net/http/httptest"
	"testing"
)

func testHandler(env *model.Env, w http.ResponseWriter, r *http.Request) {
	response.Send(w, http.StatusOK, "ok", 1, "OK")
}

func TestAuthHandler_ValidAuthToken(t *testing.T) {
	req := httptest.NewRequest("GET", "/auth", nil)
	req.Header.Add("auth_token", "asdf")

	rr := httptest.NewRecorder()

	env := &model.Env{
		DB: &model.DBMock{
			ExistsMock: model.ExistsMock{true, nil},
		},
		Token: &model.TokenMock{
			DecryptedMock: model.DecryptedMock{"decrypted", nil},
		},
	}

	appHandler := app.Handler{env, testHandler}

	var handler http.Handler = Handler(appHandler)

	handler.ServeHTTP(rr, req)

	// Check the status code is what we expect.
	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusOK)
	}

	// Check the response body is what we expect.
	expected := `{"status":200,"message":"ok","count":1,"results":"OK"}`
	if rr.Body.String() != expected {
		t.Errorf("handler returned unexpected body: got %v want %v",
			rr.Body.String(), expected)
	}
}

func TestAuthHandler_MissingAuthToken(t *testing.T) {
	req := httptest.NewRequest("GET", "/auth", nil)

	rr := httptest.NewRecorder()

	env := &model.Env{
		DB:    &model.DBMock{},
		Token: &model.TokenMock{},
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

	rr := httptest.NewRecorder()

	env := &model.Env{
		DB: &model.DBMock{},
		Token: &model.TokenMock{
			DecryptedMock: model.DecryptedMock{"", errors.New("decryption error")},
		},
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
	expected := `{"status":401,"message":"auth_token is invalid"}`
	if rr.Body.String() != expected {
		t.Errorf("handler returned unexpected body: got %v want %v",
			rr.Body.String(), expected)
	}
}

func TestAuthHandler_DBError(t *testing.T) {
	req := httptest.NewRequest("GET", "/auth", nil)
	req.Header.Add("auth_token", "asdf")

	rr := httptest.NewRecorder()

	env := &model.Env{
		DB: &model.DBMock{
			ExistsMock: model.ExistsMock{false, errors.New("user does not exist")},
		},
		Token: &model.TokenMock{
			DecryptedMock: model.DecryptedMock{"decrypted", nil},
		},
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
	expected := `{"status":401,"message":"auth_token is invalid"}`
	if rr.Body.String() != expected {
		t.Errorf("handler returned unexpected body: got %v want %v",
			rr.Body.String(), expected)
	}
}

func TestAuthHandler_UserDoesNotExist(t *testing.T) {
	req := httptest.NewRequest("GET", "/auth", nil)
	req.Header.Add("auth_token", "asdf")

	rr := httptest.NewRecorder()

	env := &model.Env{
		DB: &model.DBMock{
			ExistsMock: model.ExistsMock{false, nil},
		},
		Token: &model.TokenMock{
			DecryptedMock: model.DecryptedMock{"decrypted", nil},
		},
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
	expected := `{"status":401,"message":"auth_token is invalid"}`
	if rr.Body.String() != expected {
		t.Errorf("handler returned unexpected body: got %v want %v",
			rr.Body.String(), expected)
	}
}
