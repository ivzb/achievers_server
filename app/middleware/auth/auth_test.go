package auth

import (
	"errors"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/ivzb/achievers_server/app/middleware/app"
	"github.com/ivzb/achievers_server/app/model"
	"github.com/ivzb/achievers_server/app/model/mock"
	"github.com/ivzb/achievers_server/app/shared/config"
	"github.com/ivzb/achievers_server/app/shared/response"
)

func testHandler(env *model.Env) *response.Message {
	return response.Created("auth_token", "auth token here")
}

func TestAuthHandler_ValidAuthToken(t *testing.T) {
	req := httptest.NewRequest("GET", "/auth", nil)
	req.Header.Add("auth_token", "asdf")

	rr := httptest.NewRecorder()

	env := &model.Env{
		DB: &mock.DB{
			UserExistsMock: mock.UserExists{Bool: true, Err: nil},
		},
		Token: &mock.Tokener{
			DecryptMock: mock.Decrypt{Dec: "decrypted", Err: nil},
		},
		Config: &config.Config{},
	}

	app := app.App{Env: env, Handler: testHandler}

	var handler http.Handler = Handler(app)

	handler.ServeHTTP(rr, req)

	// Check the status code is what we expect.
	if status := rr.Code; status != http.StatusCreated {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusCreated)
	}

	// Check the response body is what we expect.
	expected := `{"status":201,"message":"auth_token created","results":"auth token here"}`
	if rr.Body.String() != expected {
		t.Errorf("handler returned unexpected body: \ngot %v \nwant %v",
			rr.Body.String(), expected)
	}
}

func TestAuthHandler_MissingAuthToken(t *testing.T) {
	req := httptest.NewRequest("GET", "/auth", nil)

	rr := httptest.NewRecorder()

	env := &model.Env{
		DB:     &mock.DB{},
		Token:  &mock.Tokener{},
		Config: &config.Config{},
	}

	app := app.App{Env: env, Handler: testHandler}

	var handler http.Handler = Handler(app)

	handler.ServeHTTP(rr, req)

	// Check the status code is what we expect.
	if status := rr.Code; status != http.StatusUnauthorized {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusUnauthorized)
	}

	// Check the response body is what we expect.
	expected := `{"status":401,"message":"missing auth_token"}`
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
		DB: &mock.DB{},
		Token: &mock.Tokener{
			DecryptMock: mock.Decrypt{Dec: "", Err: errors.New("decryption error")},
		},
		Config: &config.Config{},
	}

	app := app.App{Env: env, Handler: testHandler}

	var handler http.Handler = Handler(app)

	handler.ServeHTTP(rec, req)

	// Check the status code is what we expect.
	if status := rec.Code; status != http.StatusUnauthorized {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusUnauthorized)
	}

	// Check the response body is what we expect.
	expected := `{"status":401,"message":"invalid auth_token"}`
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
		DB: &mock.DB{
			UserExistsMock: mock.UserExists{Bool: false, Err: errors.New("user does not exist")},
		},
		Token: &mock.Tokener{
			DecryptMock: mock.Decrypt{Dec: "decrypted", Err: nil},
		},
		Config: &config.Config{},
	}

	app := app.App{Env: env, Handler: testHandler}

	var handler http.Handler = Handler(app)

	handler.ServeHTTP(rec, req)

	// Check the status code is what we expect.
	if status := rec.Code; status != http.StatusInternalServerError {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusUnauthorized)
	}

	// Check the response body is what we expect.
	expected := `{"status":500,"message":"an error occurred, please try again later"}`
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
		DB: &mock.DB{
			UserExistsMock: mock.UserExists{Bool: false, Err: nil},
		},
		Token: &mock.Tokener{
			DecryptMock: mock.Decrypt{Dec: "decrypted", Err: nil},
		},
		Config: &config.Config{},
	}

	app := app.App{Env: env, Handler: testHandler}

	var handler http.Handler = Handler(app)

	handler.ServeHTTP(rec, req)

	// Check the status code is what we expect.
	if status := rec.Code; status != http.StatusUnauthorized {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusUnauthorized)
	}

	// Check the response body is what we expect.
	expected := `{"status":401,"message":"invalid auth_token"}`
	if rec.Body.String() != expected {
		t.Errorf("handler returned unexpected body: got %v want %v",
			rec.Body.String(), expected)
	}
}
