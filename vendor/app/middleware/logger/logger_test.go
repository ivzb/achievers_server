package logger

import (
	"app/middleware/app"
	"app/model"
	"app/shared/response"
	"errors"
	"net/http"
	"net/http/httptest"
	"testing"
)

func testHandler(env *model.Env, w http.ResponseWriter, r *http.Request) response.Message {
	return response.Send(http.StatusOK, "ok", 1, "OK")
}

func TestLoggerHandler_Log(t *testing.T) {
	req := httptest.NewRequest("GET", "/logger", nil)

	rec := httptest.NewRecorder()

	env := &model.Env{
		Logger: &model.LoggerMock{
			LogMock: model.LogMock{nil},
		},
	}

	appHandler := app.Handler{env, testHandler}

	var handler http.Handler = Handler(appHandler)

	handler.ServeHTTP(rec, req)

	// Check the status code is what we expect.
	if status := rec.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusOK)
	}

	// Check the response body is what we expect.
	expected := `{"status":200,"message":"ok","count":1,"results":"OK"}`
	if rec.Body.String() != expected {
		t.Errorf("handler returned unexpected body: got %v want %v",
			rec.Body.String(), expected)
	}
}

func TestLoggerHandler_Error(t *testing.T) {
	req := httptest.NewRequest("GET", "/logger", nil)

	rec := httptest.NewRecorder()

	env := &model.Env{
		Logger: &model.LoggerMock{
			LogMock: model.LogMock{errors.New("logger error")},
		},
	}

	appHandler := app.Handler{env, testHandler}

	var handler http.Handler = Handler(appHandler)

	handler.ServeHTTP(rec, req)

	// Check the status code is what we expect.
	if status := rec.Code; status != http.StatusInternalServerError {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusInternalServerError)
	}

	// Check the response body is what we expect.
	expected := `{"status":500,"message":"an error occurred, please try again later"}`
	if rec.Body.String() != expected {
		t.Errorf("handler returned unexpected body: got %v want %v",
			rec.Body.String(), expected)
	}
}
