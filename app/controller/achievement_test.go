package controller

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"net/http/httptest"
	"net/url"
	"strconv"
	"testing"

	"github.com/ivzb/achievers_server/app/model"
	"github.com/ivzb/achievers_server/app/model/mock"
)

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

func TestAchievementSingle_InvalidMethod(t *testing.T) {
	testInvalidMethod(t, "POST", "/achievement", AchievementSingle)
}

func TestAchievementSingle_MissingId(t *testing.T) {
	rec := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/achievement", nil)

	req.Form = url.Values{}
	req.Form.Add("id", "")

	handle := AchievementSingle
	statusCode := http.StatusBadRequest

	testHandler(t, rec, req, nil, handle, statusCode)

	// Check the response body is what we expect.
	expectedMessage := fmt.Sprintf(formatMissing, id)
	expected := fmt.Sprintf(`{"status":%d,"message":"%s"}`, statusCode, expectedMessage)
	if rec.Body.String() != expected {
		t.Errorf("handler returned unexpected body: got %v want %v",
			rec.Body.String(), expected)
	}
}

func TestAchievementSingle_AchievementExistDBError(t *testing.T) {
	rec := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/achievement", nil)

	req.Form = url.Values{}
	req.Form.Add("id", "random_achievement_id")

	env := model.Env{
		DB: &mock.DB{
			ExistsMock: mock.Exists{false, errors.New("db error")},
		},
		Logger: &mock.Logger{},
	}

	handle := AchievementSingle
	statusCode := http.StatusInternalServerError

	testHandler(t, rec, req, &env, handle, statusCode)

	// Check the response body is what we expect.
	expected := fmt.Sprintf(`{"status":%d,"message":"%s"}`, statusCode, friendlyErrorMessage)
	if rec.Body.String() != expected {
		t.Errorf("handler returned unexpected body: got %v want %v",
			rec.Body.String(), expected)
	}
}

func TestAchievementSingle_AchievementDoesNotExist(t *testing.T) {
	rec := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/achievement", nil)

	req.Form = url.Values{}
	req.Form.Add("id", "random_achievement_id")

	env := model.Env{
		DB: &mock.DB{
			ExistsMock: mock.Exists{false, nil},
		},
		Logger: &mock.Logger{},
	}

	handle := AchievementSingle
	statusCode := http.StatusNotFound

	testHandler(t, rec, req, &env, handle, statusCode)

	// Check the response body is what we expect.
	expectedMessage := fmt.Sprintf(formatNotFound, achievement)
	expected := fmt.Sprintf(`{"status":%d,"message":"%s"}`, statusCode, expectedMessage)
	if rec.Body.String() != expected {
		t.Errorf("handler returned unexpected body: got %v want %v",
			rec.Body.String(), expected)
	}
}

func TestAchievementSingle_AchievementSingleDBError(t *testing.T) {
	rec := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/achievement", nil)

	req.Form = url.Values{}
	req.Form.Add("id", "random_achievement_id")

	env := model.Env{
		DB: &mock.DB{
			ExistsMock:            mock.Exists{true, nil},
			AchievementSingleMock: mock.AchievementSingle{nil, errors.New("db error")},
		},
		Logger: &mock.Logger{},
	}

	handle := AchievementSingle
	statusCode := http.StatusInternalServerError

	testHandler(t, rec, req, &env, handle, statusCode)

	// Check the response body is what we expect.
	expected := fmt.Sprintf(`{"status":%d,"message":"%s"}`, statusCode, friendlyErrorMessage)
	if rec.Body.String() != expected {
		t.Errorf("handler returned unexpected body: got %v want %v",
			rec.Body.String(), expected)
	}
}

func TestAchievementSingle_AchievementSingleNil(t *testing.T) {
	rec := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/achievement", nil)

	req.Form = url.Values{}
	req.Form.Add("id", "random_achievement_id")

	env := model.Env{
		DB: &mock.DB{
			ExistsMock:            mock.Exists{true, nil},
			AchievementSingleMock: mock.AchievementSingle{nil, nil},
		},
		Logger: &mock.Logger{},
	}

	handle := AchievementSingle
	statusCode := http.StatusNotFound

	testHandler(t, rec, req, &env, handle, statusCode)

	// Check the response body is what we expect.
	expectedMessage := fmt.Sprintf(formatNotFound, achievement)
	expected := fmt.Sprintf(`{"status":%d,"message":"%s"}`, statusCode, expectedMessage)
	if rec.Body.String() != expected {
		t.Errorf("handler returned unexpected body: got %v want %v",
			rec.Body.String(), expected)
	}
}

func TestAchievementSingle_AchievementSingle(t *testing.T) {
	rec := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/achievement", nil)

	req.Form = url.Values{}
	req.Form.Add("id", "random_achievement_id")

	env := model.Env{
		DB: &mock.DB{
			ExistsMock:            mock.Exists{true, nil},
			AchievementSingleMock: mock.AchievementSingle{mock.Achievement(), nil},
		},
		Logger: &mock.Logger{},
	}

	handle := AchievementSingle
	statusCode := http.StatusOK

	testHandler(t, rec, req, &env, handle, statusCode)

	// Check the response body is what we expect.
	expectedMessage := fmt.Sprintf(formatFound, achievement)
	marshalled, _ := json.Marshal(mock.Achievement())
	expectedResults := marshalled
	expected := fmt.Sprintf(`{"status":%d,"message":"%s","length":%d,"results":%s}`,
		statusCode,
		expectedMessage,
		1,
		expectedResults)

	if rec.Body.String() != expected {
		t.Errorf("handler returned unexpected body: \ngot %v \nwant %v",
			rec.Body.String(), expected)
	}
}

func TestAchievementCreate_InvalidMethod(t *testing.T) {
	testInvalidMethod(t, "GET", "/achievement/create", AchievementCreate)
}

func TestAchievementCreate_MissingTitle(t *testing.T) {
	form := url.Values{}
	form.Add("title", "")
	form.Add("description", mockDescription)
	form.Add("picture_url", mockPictureURL)
	form.Add("involvement_id", mockInovlvementID)

	testAchievementCreate_missingFormValue(t, form, title)
}

func TestAchievementCreate_MissingDescription(t *testing.T) {
	form := url.Values{}
	form.Add("title", mockTitle)
	form.Add("description", "")
	form.Add("picture_url", mockPictureURL)
	form.Add("involvement_id", mockInovlvementID)

	testAchievementCreate_missingFormValue(t, form, description)
}

func TestAchievementCreate_MissingPictureUrl(t *testing.T) {
	form := url.Values{}
	form.Add("title", mockTitle)
	form.Add("description", mockDescription)
	form.Add("picture_url", "")
	form.Add("involvement_id", mockInovlvementID)

	testAchievementCreate_missingFormValue(t, form, pictureURL)
}

func TestAchievementCreate_MissingInvolvementId(t *testing.T) {
	form := url.Values{}
	form.Add("title", mockTitle)
	form.Add("description", mockDescription)
	form.Add("picture_url", mockPictureURL)
	form.Add("involvement_id", "")

	testAchievementCreate_missingFormValue(t, form, involvementID)
}

func TestAchievementCreate_InvolvementIdExistDBError(t *testing.T) {
	rec := httptest.NewRecorder()
	req, _ := http.NewRequest("POST", "/achievement/create", nil)

	req.Form = url.Values{}
	req.Form.Add("title", mockTitle)
	req.Form.Add("description", mockDescription)
	req.Form.Add("picture_url", mockPictureURL)
	req.Form.Add("involvement_id", mockInovlvementID)

	env := model.Env{
		DB: &mock.DB{
			ExistsMock: mock.Exists{false, errors.New("db error")},
		},
		Logger: &mock.Logger{},
	}

	handle := AchievementCreate
	statusCode := http.StatusInternalServerError

	testHandler(t, rec, req, &env, handle, statusCode)

	// Check the response body is what we expect.
	expected := fmt.Sprintf(`{"status":%d,"message":"%s"}`, statusCode, friendlyErrorMessage)
	if rec.Body.String() != expected {
		t.Errorf("handler returned unexpected body: got %v want %v",
			rec.Body.String(), expected)
	}
}

func TestAchievementCreate_InvolvementIdDoesNotExist(t *testing.T) {
	rec := httptest.NewRecorder()
	req, _ := http.NewRequest("POST", "/achievement/create", nil)

	req.Form = url.Values{}
	req.Form.Add("title", mockTitle)
	req.Form.Add("description", mockDescription)
	req.Form.Add("picture_url", mockPictureURL)
	req.Form.Add("involvement_id", mockInovlvementID)

	env := model.Env{
		DB: &mock.DB{
			ExistsMock: mock.Exists{false, nil},
		},
		Logger: &mock.Logger{},
	}

	handle := AchievementCreate
	statusCode := http.StatusBadRequest

	testHandler(t, rec, req, &env, handle, statusCode)

	// Check the response body is what we expect.
	expected := fmt.Sprintf(`{"status":%d,"message":"%s"}`, statusCode, fmt.Sprintf(formatNotFound, involvement))
	if rec.Body.String() != expected {
		t.Errorf("handler returned unexpected body: got %v want %v",
			rec.Body.String(), expected)
	}
}

func TestAchievementCreate_AchievementCreateDBError(t *testing.T) {
	rec := httptest.NewRecorder()
	req, _ := http.NewRequest("POST", "/achievement/create", nil)

	req.Form = url.Values{}
	req.Form.Add("title", mockTitle)
	req.Form.Add("description", mockDescription)
	req.Form.Add("picture_url", mockPictureURL)
	req.Form.Add("involvement_id", mockInovlvementID)

	env := model.Env{
		DB: &mock.DB{
			ExistsMock:            mock.Exists{true, nil},
			AchievementCreateMock: mock.AchievementCreate{"", errors.New("db error")},
		},
		Logger: &mock.Logger{},
	}

	handle := AchievementCreate
	statusCode := http.StatusInternalServerError

	testHandler(t, rec, req, &env, handle, statusCode)

	// Check the response body is what we expect.
	expected := fmt.Sprintf(`{"status":%d,"message":"%s"}`, statusCode, friendlyErrorMessage)
	if rec.Body.String() != expected {
		t.Errorf("handler returned unexpected body: got %v want %v",
			rec.Body.String(), expected)
	}
}

func TestAchievementCreate_ValidAchievement(t *testing.T) {
	rec := httptest.NewRecorder()
	req, _ := http.NewRequest("POST", "/achievement/create", nil)

	req.Form = url.Values{}
	req.Form.Add("title", mockTitle)
	req.Form.Add("description", mockDescription)
	req.Form.Add("picture_url", mockPictureURL)
	req.Form.Add("involvement_id", mockInovlvementID)

	env := model.Env{
		DB: &mock.DB{
			ExistsMock:            mock.Exists{true, nil},
			AchievementCreateMock: mock.AchievementCreate{mockID, nil},
		},
		Logger: &mock.Logger{},
	}

	handle := AchievementCreate
	statusCode := http.StatusOK

	testHandler(t, rec, req, &env, handle, statusCode)

	// Check the response body is what we expect.
	expected := `{"status":` + strconv.Itoa(statusCode) + `,"message":"` + fmt.Sprintf(formatFound, achievement) + `","length":1,"results":"` + mockID + `"}`
	if rec.Body.String() != expected {
		t.Errorf("handler returned unexpected body: got %v want %v",
			rec.Body.String(), expected)
	}
}

func testAchievementCreate_missingFormValue(t *testing.T, form url.Values, expectedMissing string) {
	rec := httptest.NewRecorder()
	req, _ := http.NewRequest("POST", "/achievement/create", nil)

	req.Form = form

	handle := AchievementCreate
	statusCode := http.StatusBadRequest

	testHandler(t, rec, req, nil, handle, statusCode)

	// Check the response body is what we expect.
	expectedMessage := fmt.Sprintf(formatMissing, expectedMissing)
	expected := fmt.Sprintf(`{"status":%d,"message":"%s"}`, statusCode, expectedMessage)
	if rec.Body.String() != expected {
		t.Errorf("handler returned unexpected body: got %v want %v",
			rec.Body.String(), expected)
	}
}
