package logger

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/ivzb/achievers_server/app/middleware/app"
	"github.com/ivzb/achievers_server/app/shared/config"
	"github.com/ivzb/achievers_server/app/shared/env"
	"github.com/ivzb/achievers_server/app/shared/logger/mock"
	"github.com/ivzb/achievers_server/app/shared/response"
)

func testHandler(env *env.Env) *response.Message {
	return response.Ok("ok", 1, "OK")
}

func TestLoggerHandler_Log(t *testing.T) {
	req := httptest.NewRequest("GET", "/logger", nil)

	rec := httptest.NewRecorder()

	env := &env.Env{
		Log:    &mock.Logger{},
		Config: &config.Config{},
	}

	app := app.App{Env: env, Handler: testHandler}

	var handler http.Handler = Handler(app)

	handler.ServeHTTP(rec, req)

	// Check the status code is what we expect.
	if status := rec.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusOK)
	}

	// Check the response body is what we expect.
	expected := `{"message":"ok found","results":"OK"}`
	if rec.Body.String() != expected {
		t.Errorf("handler returned unexpected body: got %v want %v",
			rec.Body.String(), expected)
	}
}
