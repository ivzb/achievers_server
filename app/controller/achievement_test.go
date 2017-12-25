package controller

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"net/http/httptest"
	"net/url"
	"testing"

	"github.com/ivzb/achievers_server/app/model"
	"github.com/ivzb/achievers_server/app/model/mock"
)

// todo: test 3 pages:
// 0 -> 9 results - full
// 1 -> 3 results - end
// 2 -> 0 results - possibly 404

func TestAchievementsIndex_FullPage(t *testing.T) {
	statusCode := http.StatusOK
	rec := requestAchievements(t, 9, statusCode)
	verifyCorrectResult(t, rec, 9)
}

func TestAchievementsIndex_HalfPage(t *testing.T) {
	statusCode := http.StatusOK
	rec := requestAchievements(t, 4, statusCode)
	verifyCorrectResult(t, rec, 4)
}

func TestAchievementsIndex_EmptyPage(t *testing.T) {
	statusCode := http.StatusNotFound
	rec := requestAchievements(t, 0, statusCode)

	expectedMessage := fmt.Sprintf(formatNotFound, page)
	expected := fmt.Sprintf(`{"status":%d,"message":"%s"}`, statusCode, expectedMessage)

	if rec.Body.String() != expected {
		t.Errorf("handler returned unexpected body: got %v want %v",
			rec.Body.String(), expected)
	}
}

func requestAchievements(t *testing.T, size int, statusCode int) *httptest.ResponseRecorder {
	rec := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/achievements", nil)

	req.Form = url.Values{}
	req.Form.Add("page", "3")

	env := model.Env{
		DB: &mock.DB{
			AchievementsAllMock: mock.AchievementsAll{A: mock.Achievements(size), E: nil},
		},
		Logger: &mock.Logger{},
	}

	handle := AchievementsIndex

	testHandler(t, rec, req, &env, handle, statusCode)

	return rec
}

func verifyCorrectResult(t *testing.T, rec *httptest.ResponseRecorder, size int) {
	expectedStatusCode := http.StatusOK
	expectedMessage := fmt.Sprintf(formatFound, achievements)
	mocks := mock.Achievements(size)
	expectedLength := len(mocks)
	marshalled, _ := json.Marshal(mocks)
	expectedResults := marshalled
	expected := fmt.Sprintf(`{"status":%d,"message":"%s","length":%d,"results":%s}`,
		expectedStatusCode,
		expectedMessage,
		expectedLength,
		expectedResults)

	if rec.Body.String() != expected {
		t.Errorf("handler returned unexpected body: \ngot %v \nwant %v",
			rec.Body.String(), expected)
	}
}

func TestAchievementsIndex_InvalidMethod(t *testing.T) {
	testInvalidMethod(t, "POST", "/achievements", AchievementsIndex)
}

func TestAchievementsIndex_MissingPage(t *testing.T) {
	rec := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/achievements", nil)

	req.Form = url.Values{}
	req.Form.Add("page", "")

	handle := AchievementsIndex
	statusCode := http.StatusBadRequest

	testHandler(t, rec, req, nil, handle, statusCode)

	// Check the response body is what we expect.
	expectedMessage := fmt.Sprintf(formatMissing, page)
	expected := fmt.Sprintf(`{"status":%d,"message":"%s"}`, statusCode, expectedMessage)
	if rec.Body.String() != expected {
		t.Errorf("handler returned unexpected body: got %v want %v",
			rec.Body.String(), expected)
	}
}

func TestAchievementsIndex_InvalidPage(t *testing.T) {
	rec := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/achievements", nil)

	req.Form = url.Values{}
	req.Form.Add("page", "-1")

	handle := AchievementsIndex
	statusCode := http.StatusBadRequest

	testHandler(t, rec, req, nil, handle, statusCode)

	// Check the response body is what we expect.
	expectedMessage := fmt.Sprintf(formatInvalid, page)
	expected := fmt.Sprintf(`{"status":%d,"message":"%s"}`, statusCode, expectedMessage)
	if rec.Body.String() != expected {
		t.Errorf("handler returned unexpected body: got %v want %v",
			rec.Body.String(), expected)
	}
}

func TestAchievementsIndex_DBError(t *testing.T) {
	rec := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/achievements", nil)

	req.Form = url.Values{}
	req.Form.Add("page", "3")

	env := model.Env{
		DB: &mock.DB{
			AchievementsAllMock: mock.AchievementsAll{nil, errors.New("db error")},
		},
		Logger: &mock.Logger{},
	}

	handle := AchievementsIndex
	statusCode := http.StatusInternalServerError

	testHandler(t, rec, req, &env, handle, statusCode)

	// Check the response body is what we expect.
	expected := fmt.Sprintf(`{"status":%d,"message":"%s"}`, statusCode, friendlyErrorMessage)
	if rec.Body.String() != expected {
		t.Errorf("handler returned unexpected body: got %v want %v",
			rec.Body.String(), expected)
	}
}
