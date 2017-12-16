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

func testHandler(env *model.Env, w http.ResponseWriter, r *http.Request) {
	response.Send(w, http.StatusOK, "ok", 1, "OK")
}

func TestLoggerHandler_Log(t *testing.T) {
	req := httptest.NewRequest("GET", "/logger", nil)

	rr := httptest.NewRecorder()

	env := &model.Env{
		Logger: &model.LoggerMock{
			LogMock: model.LogMock{nil},
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

func TestLoggerHandler_Error(t *testing.T) {
	req := httptest.NewRequest("GET", "/logger", nil)

	rr := httptest.NewRecorder()

	env := &model.Env{
		Logger: &model.LoggerMock{
			LogMock: model.LogMock{errors.New("logger error")},
		},
	}

	appHandler := app.Handler{env, testHandler}

	var handler http.Handler = Handler(appHandler)

	handler.ServeHTTP(rr, req)

	// Check the status code is what we expect.
	if status := rr.Code; status != http.StatusInternalServerError {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusInternalServerError)
	}

	// Check the response body is what we expect.
	expected := `{"status":500,"message":"an error occurred, please try again later"}`
	if rr.Body.String() != expected {
		t.Errorf("handler returned unexpected body: got %v want %v",
			rr.Body.String(), expected)
	}
}
