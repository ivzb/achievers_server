package controller

import (
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
)

func testInvalidMethod(t *testing.T, method string, url string, handle app.Handle) {
	rec := httptest.NewRecorder()
	req, _ := http.NewRequest(method, url, nil)

	appHandler := app.Handler{nil, handle}

	appHandler.ServeHTTP(rec, req)

	// Check the status code is what we expect.
	if status := rec.Code; status != http.StatusMethodNotAllowed {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusMethodNotAllowed)
	}

	// Check the response body is what we expect.
	expected := `{"status":405,"message":"` + methodNotAllowed + `"}`
	if rec.Body.String() != expected {
		t.Errorf("handler returned unexpected body: got %v want %v",
			rec.Body.String(), expected)
	}
}

func testStatusCode(t *testing.T, rec *httptest.ResponseRecorder, expectedStatusCode int) {
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

	testStatusCode(t, rec, statusCode)
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
	expectedMessage := fmt.Sprintf(formatMissing, expectedMissing)
	expected := fmt.Sprintf(`{"status":%d,"message":"%s"}`, statusCode, expectedMessage)
	if rec.Body.String() != expected {
		t.Errorf("handler returned unexpected body: got %v want %v",
			rec.Body.String(), expected)
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
