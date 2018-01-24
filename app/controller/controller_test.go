package controller

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"net/http/httptest"
	"net/url"
	"strings"
	"testing"

	"github.com/ivzb/achievers_server/app/middleware/app"
	"github.com/ivzb/achievers_server/app/shared/env"

	dMock "github.com/ivzb/achievers_server/app/db/mock"
	lMock "github.com/ivzb/achievers_server/app/shared/logger/mock"
	tMock "github.com/ivzb/achievers_server/app/shared/token/mock"
)

const (
	Core     = 0
	Change   = 1
	Retrieve = 2
)

var (
	mockID               = "mock id"
	mockTitle            = "mock title"
	mockDescription      = "mock description"
	mockPictureURL       = "mock picture_url"
	mockName             = "mock name"
	mockEmail            = "mock email"
	mockPassword         = "mock password"
	mockToken            = "mock token"
	mockURL              = "mock url"
	mockEncrypt          = "mock encrypt"
	mockAchievementID    = "mock achievement_id"
	mockInvolvementID    = "3"
	mockRewardTypeID     = "5"
	mockQuestTypeID      = "5"
	mockMultimediaTypeID = "5"

	mockDbErr     = errors.New("db error")
	mockFormerErr = errors.New("former error")

	mockPage     = "3"
	mockPageSize = 9
)

type test struct {
	purpose  string
	handler  app.Handler
	request  *testRequest
	response *testResponse
}

type testRequest struct {
	method        string
	form          *url.Values
	env           *env.Env
	removeHeaders bool
}

type testResponse struct {
	typ        int
	statusCode int
	message    string
	results    []byte
}

type testInput struct {
	purpose            string
	requestMethod      string
	responseType       int
	responseStatusCode int
	responseMessage    string
	form               *map[string]string
	db                 *dMock.DB
	logger             *lMock.Logger
	tokener            *tMock.Tokener
	args               []string
	removeHeaders      bool
}

func constructForm(m *map[string]string) *url.Values {
	form := &url.Values{}

	if m != nil {
		for key, value := range *m {
			form.Add(key, value)
		}
	}

	return form
}

func constructEnv(db *dMock.DB, logger *lMock.Logger, tokener *tMock.Tokener) *env.Env {
	return &env.Env{
		DB:    db,
		Log:   logger,
		Token: tokener,
	}
}

func constructTestRequest(method string, form *url.Values, env *env.Env, removeHeaders bool) *testRequest {
	return &testRequest{
		method,
		form,
		env,
		removeHeaders,
	}
}

func constructTestResponse(typ int, statusCode int, message string, results []byte) *testResponse {
	return &testResponse{
		typ,
		statusCode,
		message,
		results,
	}
}

func constructRequest(t *testing.T, test *test) *httptest.ResponseRecorder {
	rec := httptest.NewRecorder()
	req, _ := http.NewRequest(test.request.method, "/test", encodeForm(*test.request.form))

	req.Form = *test.request.form

	if !test.request.removeHeaders {
		req.Header = http.Header{}
		req.Header.Add("Content-Type", "application/x-www-form-urlencoded")
	}

	test.request.env.Request = req
	app := app.App{test.request.env, test.handler}
	serveHTTP(app, rec, req)

	actualStatusCode := rec.Code
	expectedStatusCode := test.response.statusCode

	if actualStatusCode != expectedStatusCode {
		t.Errorf("handler returned wrong status code: got %v want %v",
			actualStatusCode, expectedStatusCode)
	}

	return rec
}

func serveHTTP(app app.App, w http.ResponseWriter, r *http.Request) {
	response := app.Handler(app.Env)

	js, err := json.Marshal(response.Result)

	if err != nil {
		http.Error(w, "JSON Error: "+err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(response.StatusCode)
	w.Write(js)
}
func constructTest(handler app.Handler, testInput *testInput, responseResults []byte) *test {
	return &test{
		purpose: testInput.purpose,
		handler: handler,
		request: constructTestRequest(
			testInput.requestMethod,
			constructForm(testInput.form),
			constructEnv(testInput.db, testInput.logger, testInput.tokener),
			testInput.removeHeaders,
		),
		response: constructTestResponse(
			testInput.responseType,
			testInput.responseStatusCode,
			testInput.responseMessage,
			responseResults,
		),
	}
}

func createRequest(form url.Values) (*httptest.ResponseRecorder, *http.Request) {
	rec := httptest.NewRecorder()
	req, _ := http.NewRequest("POST", "/mock/path", encodeForm(form))

	req.Header = http.Header{}
	req.Header.Add("Content-Type", "application/x-www-form-urlencoded")

	return rec, req
}

func encodeForm(form url.Values) *strings.Reader {
	return strings.NewReader(form.Encode())
}

func run(t *testing.T, tests []*test) {
	for _, test := range tests {
		rec := constructRequest(t, test)

		switch test.response.typ {
		case Core:
			assertCoreResponse(t, rec, test)
		case Retrieve:
			assertRetrieveResponse(t, rec, test)
		}
	}
}

func assertCoreResponse(
	t *testing.T,
	rec *httptest.ResponseRecorder,
	test *test) {

	expected := fmt.Sprintf(`{"message":"%s"}`, test.response.message)
	actual := rec.Body.String()

	if actual != expected {
		assertFailed(t, test.purpose, expected, actual)
	}
}

func assertRetrieveResponse(
	t *testing.T,
	rec *httptest.ResponseRecorder,
	test *test) {

	expected := fmt.Sprintf(`{"message":"%s","results":%s}`,
		test.response.message,
		test.response.results)

	actual := rec.Body.String()

	if actual != expected {
		assertFailed(t, test.purpose, expected, actual)
	}
}

func assertFailed(t *testing.T, purpose string, expected string, actual string) {
	t.Errorf("test %s", purpose)

	t.Errorf("handler returned unexpected body: \ngot %v \nwant %v",
		actual, expected)
}

func mapWithout(m *map[string]string, key string) *map[string]string {
	(*m)[key] = ""
	return m
}
