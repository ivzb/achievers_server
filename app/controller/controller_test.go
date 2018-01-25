package controller

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"net/url"
	"strings"
	"testing"

	"github.com/ivzb/achievers_server/app/middleware/app"
	"github.com/ivzb/achievers_server/app/shared/config"
	"github.com/ivzb/achievers_server/app/shared/env"

	dMock "github.com/ivzb/achievers_server/app/db/mock"
	lMock "github.com/ivzb/achievers_server/app/shared/logger/mock"
	tMock "github.com/ivzb/achievers_server/app/shared/token/mock"
	uMock "github.com/ivzb/achievers_server/app/shared/uuid/mock"
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
	mockFileID           = "."
	mockAchievementID    = "mock achievement_id"
	mockInvolvementID    = "3"
	mockRewardTypeID     = "5"
	mockQuestTypeID      = "5"
	mockMultimediaTypeID = "5"
	mockMultipartBoundry = "MockBoundry"
	mockUUID             = "mock_uuid"

	mockDbErr     = errors.New("db error")
	mockFormerErr = errors.New("former error")
	mockUUIDerErr = errors.New("UUIDer error")

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
	multipartForm string
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
	multipartForm      string
	db                 *dMock.DB
	logger             *lMock.Logger
	tokener            *tMock.Tokener
	uuider             *uMock.UUIDer
	config             *config.Config
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

func constructMultipartForm(data string) io.Reader {
	return ioutil.NopCloser(strings.NewReader(data))
}

func constructEnv(
	db *dMock.DB,
	logger *lMock.Logger,
	tokener *tMock.Tokener,
	uuider *uMock.UUIDer,
	config *config.Config) *env.Env {

	return &env.Env{
		DB:     db,
		Log:    logger,
		Token:  tokener,
		UUID:   uuider,
		Config: config,
	}
}

func constructTestRequest(
	method string,
	form *url.Values,
	multipartForm string,
	env *env.Env,
	removeHeaders bool) *testRequest {

	return &testRequest{
		method,
		form,
		multipartForm,
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

	var body io.Reader

	if test.request.form != nil {
		body = encodeForm(*test.request.form)
	}

	if test.request.multipartForm != "" {
		body = strings.NewReader(test.request.multipartForm)
	}

	req, err := http.NewRequest(test.request.method, "/test", body)

	if err != nil {
		t.Fatal(err)
	}

	req.Form = *test.request.form

	if !test.request.removeHeaders {
		req.Header = http.Header{}
		var contentType string

		if test.request.form != nil {
			contentType = "application/x-www-form-urlencoded"
		}

		if test.request.multipartForm != "" {
			contentType = `multipart/form-data; boundary="` + mockMultipartBoundry + `"`
		}

		req.Header.Add("Content-Type", contentType)
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
			//constructMultipartForm(testInput.multipartForm),
			testInput.multipartForm,
			constructEnv(testInput.db, testInput.logger, testInput.tokener, testInput.uuider, testInput.config),
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

func createRequest(form io.Reader) (*httptest.ResponseRecorder, *http.Request) {
	rec := httptest.NewRecorder()

	req, _ := http.NewRequest("POST", "/mock/path", form) //encodeForm(form))

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
