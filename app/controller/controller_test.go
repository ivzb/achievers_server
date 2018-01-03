package controller

import (
	"errors"
	"fmt"
	"net/http"
	"net/http/httptest"
	"net/url"
	"strings"
	"testing"

	"github.com/ivzb/achievers_server/app/middleware/app"
	"github.com/ivzb/achievers_server/app/model"
	"github.com/ivzb/achievers_server/app/model/mock"
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
	mockInovlvementID    = "mock involvement_id"
	mockFirstName        = "mock first_name"
	mockLastName         = "mock last_name"
	mockEmail            = "mock email"
	mockPassword         = "mock password"
	mockToken            = "mock token"
	mockPreviewURL       = "mock preview_url"
	mockURL              = "mock url"
	mockMultimediaTypeID = "5"
	mockAchievementID    = "mock achievement_id"
	mockInvolvementID    = "5"
	mockEncrypt          = "mock encrypt"

	mockDbErr     = errors.New("db error")
	mockFormerErr = errors.New("former error")

	mockPage     = "3"
	mockPageSize = 9

	get  = "GET"
	post = "POST"
)

type test struct {
	purpose  string
	handle   app.Handle
	request  *testRequest
	response *testResponse
}

type testRequest struct {
	method string
	form   *url.Values
	env    *model.Env
}

type testResponse struct {
	kind       int
	statusCode int
	message    string
	results    []byte
}

func run(t *testing.T, tests []*test) {
	for _, test := range tests {
		rec := constructRequest(t, test)
		expect(t, rec, test)
	}
}

func constructForm(m map[string]string) *url.Values {
	form := &url.Values{}

	for key, value := range m {
		form.Add(key, value)
	}

	return form
}

func constructEnv(db *mock.DB, logger *mock.Logger, former *mock.Former, tokener *mock.Tokener) *model.Env {
	return &model.Env{
		DB:      db,
		Logger:  logger,
		Former:  former,
		Tokener: tokener,
	}
}

func constructTestRequest(method string, form *url.Values, env *model.Env) *testRequest {
	return &testRequest{
		method,
		form,
		env,
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

	req.Header = http.Header{}
	req.Header.Add("Content-Type", "application/x-www-form-urlencoded")

	testHandler(t, rec, req, test.request.env, test.handle, test.response.statusCode)

	return rec
}

func testMethodNotAllowed(t *testing.T, method string, url string, handle app.Handle) {
	rec := httptest.NewRecorder()
	req, _ := http.NewRequest(method, url, nil)

	appHandler := app.Handler{nil, handle}

	appHandler.ServeHTTP(rec, req)

	// Check the status code is what we expect.
	if status := rec.Code; status != http.StatusMethodNotAllowed {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusMethodNotAllowed)
	}

	statusCode := http.StatusMethodNotAllowed
	message := methodNotAllowed

	expectCore(t, rec, statusCode, message)
}

func checkStatusCode(t *testing.T, rec *httptest.ResponseRecorder, expectedStatusCode int) {
	// Check the status code is what we expect.
	if actualStatusCode := rec.Code; actualStatusCode != expectedStatusCode {
		t.Errorf("handler returned wrong status code: got %v want %v",
			actualStatusCode, expectedStatusCode)
	}
}

func testHandler(
	t *testing.T,
	rec *httptest.ResponseRecorder,
	req *http.Request,
	env *model.Env,
	handle app.Handle,
	statusCode int) {

	appHandler := app.Handler{env, handle}

	appHandler.ServeHTTP(rec, req)

	checkStatusCode(t, rec, statusCode)
}

func testMissingFormValue(t *testing.T, handle app.Handle, form url.Values, expectedMissing string) {
	form.Set(expectedMissing, "")

	rec, req := createRequest(form)

	statusCode := http.StatusBadRequest

	env := &model.Env{
		Former: &mock.Former{
			MapMock: mock.Map{nil},
		},
	}

	testHandler(t, rec, req, env, handle, statusCode)

	// Check the response body is what we expect.
	message := fmt.Sprintf(formatMissing, expectedMissing)
	expectCore(t, rec, statusCode, message)
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

func expect(t *testing.T, rec *httptest.ResponseRecorder, test *test) {
	switch test.response.kind {
	case Core:
		expectCoreTest(t, rec, test)
	case Retrieve:
		expectRetrieveTest(t, rec, test)
	}
}

func expectCore(t *testing.T, rec *httptest.ResponseRecorder, statusCode int, message string) {
	expected := fmt.Sprintf(`{"status":%d,"message":"%s"}`, statusCode, message)

	if rec.Body.String() != expected {
		t.Errorf("handler returned unexpected body: \ngot %v \nwant %v",
			rec.Body.String(), expected)
	}
}

func expectCoreTest(t *testing.T, rec *httptest.ResponseRecorder, test *test) {
	expected := fmt.Sprintf(`{"status":%d,"message":"%s"}`, test.response.statusCode, test.response.message)

	if rec.Body.String() != expected {
		t.Errorf("test %v", test.purpose)

		t.Errorf("handler returned unexpected body: \ngot %v \nwant %v",
			rec.Body.String(), expected)
	}
}

func expectRetrieve(
	t *testing.T,
	rec *httptest.ResponseRecorder,
	statusCode int,
	message string,
	results []byte) {

	expected := fmt.Sprintf(`{"status":%d,"message":"%s","results":%s}`,
		statusCode,
		message,
		results)

	if rec.Body.String() != expected {
		t.Errorf("handler returned unexpected body: \ngot %v \nwant %v",
			rec.Body.String(), expected)
	}
}

func expectRetrieveTest(
	t *testing.T,
	rec *httptest.ResponseRecorder,
	test *test) {

	expected := fmt.Sprintf(`{"status":%d,"message":"%s","results":%s}`,
		test.response.statusCode,
		test.response.message,
		test.response.results)

	if rec.Body.String() != expected {
		t.Errorf("test %v", test.purpose)

		t.Errorf("handler returned unexpected body: \ngot %v \nwant %v",
			rec.Body.String(), expected)
	}
}
