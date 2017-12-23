package app

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/ivzb/achievers_server/app/model"
	"github.com/ivzb/achievers_server/app/shared/response"
)

func testHandler(env *model.Env, w http.ResponseWriter, r *http.Request) response.Message {
	return response.Ok("ok", 1, "OK")
}

func jsonErrorHandler(env *model.Env, w http.ResponseWriter, r *http.Request) response.Message {
	return response.Message{http.StatusOK, func() {}}
}

func TestAppHandler_ValidHandler(t *testing.T) {
	req := httptest.NewRequest("GET", "/mock", nil)
	rr := httptest.NewRecorder()

	env := &model.Env{
		DB: &model.DBMock{
			ExistsMock: model.ExistsMock{true, nil},
		},
		Tokener: &model.TokenerMock{
			DecryptedMock: model.DecryptedMock{"decrypted", nil},
		},
	}

	appHandler := Handler{env, testHandler}

	var handler http.Handler = Handler(appHandler)

	handler.ServeHTTP(rr, req)

	// Check the status code is what we expect.
	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusOK)
	}

	// Check the response body is what we expect.
	expected := `{"status":200,"message":"ok","length":1,"results":"OK"}`
	if rr.Body.String() != expected {
		t.Errorf("handler returned unexpected body: got %v want %v",
			rr.Body.String(), expected)
	}
}

func TestAppHandler_InvalidJSON(t *testing.T) {
	req := httptest.NewRequest("GET", "/mock", nil)
	rr := httptest.NewRecorder()

	env := &model.Env{
		DB: &model.DBMock{
			ExistsMock: model.ExistsMock{true, nil},
		},
		Tokener: &model.TokenerMock{
			DecryptedMock: model.DecryptedMock{"decrypted", nil},
		},
	}

	appHandler := Handler{env, jsonErrorHandler}

	var handler http.Handler = Handler(appHandler)

	handler.ServeHTTP(rr, req)

	expectedStatus := http.StatusInternalServerError
	actualStatus := rr.Code

	// Check the status code is what we expect.
	if actualStatus != expectedStatus {
		t.Errorf("handler returned wrong status code: got %v want %v",
			actualStatus, expectedStatus)
	}

	// Check the response body is what we expect.
	expectedBody := "JSON Error: json: unsupported type: func()\n"
	actualBody := rr.Body.String()

	if actualBody != expectedBody {
		t.Errorf("handler returned unexpected body: got %v want %v", actualBody, expectedBody)
	}
}
