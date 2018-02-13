package test

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

	a "github.com/ivzb/achievers_server/app"
	"github.com/ivzb/achievers_server/app/middleware/app"
	"github.com/ivzb/achievers_server/app/shared/config"

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
	MockID               = "c85a9afc-8e44-4ac8-9b4e-f56608066d3f"
	MockTitle            = "mock title"
	MockDescription      = "mock description"
	MockPictureURL       = "mock picture_url"
	MockName             = "mock name"
	MockEmail            = "mock email"
	MockPassword         = "mock password"
	MockToken            = "mock token"
	MockURL              = "mock url"
	MockEncrypt          = "mock encrypt"
	MockFileID           = "."
	MockAchievementID    = "c85a9afc-8e44-4ac8-9b4e-f56608066d3f"
	MockInvolvementID    = "3"
	MockRewardTypeID     = "5"
	MockQuestTypeID      = "5"
	MockMultimediaTypeID = "5"
	MockMultipartBoundry = "MockBoundry"
	MockUUID             = "mock_uuid"

	MockDbErr     = errors.New("db error")
	MockFormerErr = errors.New("former error")
	MockUUIDerErr = errors.New("UUIDer error")

	MockPage     = "3"
	MockPageSize = 9
)

type Test struct {
	Purpose  string
	Handler  app.Handler
	Request  *TestRequest
	Response *TestResponse
}

type TestRequest struct {
	Method        string
	Form          *url.Values
	MultipartForm string
	Env           *a.Env
	RemoveHeaders bool
}

type TestResponse struct {
	Typ        int
	StatusCode int
	Message    string
	Results    []byte
}

type TestInput struct {
	Purpose            string
	RequestMethod      string
	ResponseType       int
	ResponseStatusCode int
	ResponseMessage    string
	Form               *map[string]string
	MultipartForm      string
	DB                 *dMock.DB
	Logger             *lMock.Logger
	Tokener            *tMock.Tokener
	UUIDer             *uMock.UUIDer
	Config             *config.Config
	Args               []string
	RemoveHeaders      bool
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
	config *config.Config) *a.Env {

	return &a.Env{
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
	env *a.Env,
	removeHeaders bool) *TestRequest {

	return &TestRequest{
		method,
		form,
		multipartForm,
		env,
		removeHeaders,
	}
}

func constructTestResponse(typ int, statusCode int, message string, results []byte) *TestResponse {
	return &TestResponse{
		typ,
		statusCode,
		message,
		results,
	}
}

func constructRequest(t *testing.T, test *Test) *httptest.ResponseRecorder {
	rec := httptest.NewRecorder()

	var body io.Reader

	if test.Request.Form != nil {
		body = encodeForm(*test.Request.Form)
	}

	if test.Request.MultipartForm != "" {
		body = strings.NewReader(test.Request.MultipartForm)
	}

	req, err := http.NewRequest(test.Request.Method, "/test", body)

	if err != nil {
		t.Fatal(err)
	}

	req.Form = *test.Request.Form

	if !test.Request.RemoveHeaders {
		req.Header = http.Header{}
		var contentType string

		if test.Request.Form != nil {
			contentType = "application/x-www-form-urlencoded"
		}

		if test.Request.MultipartForm != "" {
			contentType = `multipart/form-data; boundary="` + MockMultipartBoundry + `"`
		}

		req.Header.Add("Content-Type", contentType)
	}

	test.Request.Env.Request = req
	app := app.App{test.Request.Env, test.Handler}
	serveHTTP(app, rec, req)

	actualStatusCode := rec.Code
	expectedStatusCode := test.Response.StatusCode

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

func ConstructTest(handler app.Handler, testInput *TestInput, responseResults []byte) *Test {
	return &Test{
		Purpose: testInput.Purpose,
		Handler: handler,
		Request: constructTestRequest(
			testInput.RequestMethod,
			constructForm(testInput.Form),
			testInput.MultipartForm,
			constructEnv(
				testInput.DB,
				testInput.Logger,
				testInput.Tokener,
				testInput.UUIDer,
				testInput.Config),
			testInput.RemoveHeaders,
		),
		Response: constructTestResponse(
			testInput.ResponseType,
			testInput.ResponseStatusCode,
			testInput.ResponseMessage,
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

func Run(t *testing.T, tests []*Test) {
	for _, test := range tests {
		rec := constructRequest(t, test)

		switch test.Response.Typ {
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
	test *Test) {

	expected := fmt.Sprintf(`{"message":"%s"}`, test.Response.Message)
	actual := rec.Body.String()

	if actual != expected {
		assertFailed(t, test.Purpose, expected, actual)
	}
}

func assertRetrieveResponse(
	t *testing.T,
	rec *httptest.ResponseRecorder,
	test *Test) {

	expected := fmt.Sprintf(`{"message":"%s","results":%s}`,
		test.Response.Message,
		test.Response.Results)

	actual := rec.Body.String()

	if actual != expected {
		assertFailed(t, test.Purpose, expected, actual)
	}
}

func assertFailed(t *testing.T, purpose string, expected string, actual string) {
	t.Errorf("test %s", purpose)

	t.Errorf("handler returned unexpected body: \ngot %v \nwant %v",
		actual, expected)
}

func MapWithout(m *map[string]string, key string) *map[string]string {
	(*m)[key] = ""
	return m
}
