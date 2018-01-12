package app

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/ivzb/achievers_server/app/model"
	"github.com/ivzb/achievers_server/app/model/mock"
	"github.com/ivzb/achievers_server/app/shared/config"
	"github.com/ivzb/achievers_server/app/shared/response"
)

func testOkHandler(env *model.Env) *response.Message {
	return response.Ok("ok", 1, "OK")
}

func jsonErrorHandler(env *model.Env) *response.Message {
	return &response.Message{http.StatusOK, func() {}, response.TypeJSON}
}

func testFileHandler(env *model.Env) *response.Message {
	return &response.Message{
		StatusCode: 200,
		Result:     ".",
		Type:       response.TypeFile,
	}
}

func TestAppHandler_ValidJSONHandler(t *testing.T) {
	req := httptest.NewRequest("GET", "/mock", nil)
	rr := httptest.NewRecorder()

	env := &model.Env{
		Token: &mock.Tokener{
			DecryptMock: mock.Decrypt{"decrypted", nil},
		},
		Config: &config.Config{},
	}

	app := App{env, testOkHandler}

	var handler http.Handler = App(app)

	handler.ServeHTTP(rr, req)

	// Check the status code is what we expect.
	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusOK)
	}

	// Check the response body is what we expect.
	expected := `{"status":200,"message":"ok found","results":"OK"}`
	if rr.Body.String() != expected {
		t.Errorf("handler returned unexpected body: got %v want %v",
			rr.Body.String(), expected)
	}
}

func TestAppHandler_InvalidJSONHandler(t *testing.T) {
	req := httptest.NewRequest("GET", "/mock", nil)
	rr := httptest.NewRecorder()

	env := &model.Env{
		Token: &mock.Tokener{
			DecryptMock: mock.Decrypt{"decrypted", nil},
		},
		Config: &config.Config{},
	}

	app := App{env, jsonErrorHandler}

	var handler http.Handler = App(app)

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

func TestAppHandler_ValidFileHandler(t *testing.T) {
	req := httptest.NewRequest("GET", "/mock", nil)
	rr := httptest.NewRecorder()

	env := &model.Env{
		Token: &mock.Tokener{
			DecryptMock: mock.Decrypt{"decrypted", nil},
		},
		Config: &config.Config{},
	}

	app := App{env, testFileHandler}

	var handler http.Handler = App(app)

	handler.ServeHTTP(rr, req)

	// Check the status code is what we expect.
	if status := rr.Code; status != http.StatusMovedPermanently {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusOK)
	}
}
