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
	verifyCorrectAchievementsResult(t, rec, 9)
}

func TestAchievementsIndex_HalfPage(t *testing.T) {
	statusCode := http.StatusOK
	rec := requestAchievements(t, 4, statusCode)
	verifyCorrectAchievementsResult(t, rec, 4)
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
	req.Form.Add(page, "3")

	env := model.Env{
		DB: &mock.DB{
			AchievementsAllMock: mock.AchievementsAll{Achs: mock.Achievements(size), Err: nil},
		},
		Logger: &mock.Logger{},
	}

	handle := AchievementsIndex

	testHandler(t, rec, req, &env, handle, statusCode)

	return rec
}

func verifyCorrectAchievementsResult(t *testing.T, rec *httptest.ResponseRecorder, size int) {
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
	req.Form.Add(page, "")

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
	req.Form.Add(page, "-1")

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
	req.Form.Add(page, "3")

	env := model.Env{
		DB: &mock.DB{
			AchievementsAllMock: mock.AchievementsAll{Achs: nil, Err: errors.New("db error")},
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
	req.Form.Add(id, "")

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
	req.Form.Add(id, "random_achievement_id")

	env := model.Env{
		DB: &mock.DB{
			AchievementExistsMock: mock.AchievementExists{Bool: false, Err: errors.New("db error")},
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
	req.Form.Add(id, "random_achievement_id")

	env := model.Env{
		DB: &mock.DB{
			AchievementExistsMock: mock.AchievementExists{Bool: false, Err: nil},
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
	req.Form.Add(id, "random_achievement_id")

	env := model.Env{
		DB: &mock.DB{
			AchievementExistsMock: mock.AchievementExists{Bool: true, Err: nil},
			AchievementSingleMock: mock.AchievementSingle{Ach: nil, Err: errors.New("db error")},
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
	req.Form.Add(id, "random_achievement_id")

	env := model.Env{
		DB: &mock.DB{
			AchievementExistsMock: mock.AchievementExists{Bool: true, Err: nil},
			AchievementSingleMock: mock.AchievementSingle{Ach: nil, Err: nil},
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
	req.Form.Add(id, "random_achievement_id")

	env := model.Env{
		DB: &mock.DB{
			AchievementExistsMock: mock.AchievementExists{Bool: true, Err: nil},
			AchievementSingleMock: mock.AchievementSingle{Ach: mock.Achievement(), Err: nil},
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

func mockAchievementForm() url.Values {
	form := url.Values{}
	form.Add(title, mockTitle)
	form.Add(description, mockDescription)
	form.Add(pictureURL, mockPictureURL)
	form.Add(involvementID, mockInovlvementID)

	return form
}

func TestAchievementCreate_InvalidMethod(t *testing.T) {
	testInvalidMethod(t, "GET", "/achievement/create", AchievementCreate)
}

func TestAchievementCreate_FormMapError(t *testing.T) {
	rec, req := createRequest(nil)
	statusCode := http.StatusBadRequest
	handle := AchievementCreate

	mapError := "map error"

	env := &model.Env{
		Former: &mock.Former{
			MapMock: mock.Map{Err: errors.New(mapError)},
		},
	}

	testHandler(t, rec, req, env, handle, statusCode)

	// Check the response body is what we expect.
	expectedMessage := mapError
	expected := fmt.Sprintf(`{"status":%d,"message":"%s"}`, statusCode, expectedMessage)
	if rec.Body.String() != expected {
		t.Errorf("handler returned unexpected body: got %v want %v",
			rec.Body.String(), expected)
	}
}

func TestAchievementCreate_MissingTitle(t *testing.T) {
	testMissingFormValue(t, AchievementCreate, mockAchievementForm(), title)
}

func TestAchievementCreate_MissingDescription(t *testing.T) {
	testMissingFormValue(t, AchievementCreate, mockAchievementForm(), description)
}

func TestAchievementCreate_MissingPictureUrl(t *testing.T) {
	testMissingFormValue(t, AchievementCreate, mockAchievementForm(), pictureURL)
}

func TestAchievementCreate_MissingInvolvementId(t *testing.T) {
	testMissingFormValue(t, AchievementCreate, mockAchievementForm(), involvementID)
}

func TestAchievementCreate_InvolvementIdExistDBError(t *testing.T) {
	rec, req := createRequest(mockAchievementForm())

	env := model.Env{
		DB: &mock.DB{
			InvolvementExistsMock: mock.InvolvementExists{Bool: false, Err: errors.New("db error")},
		},
		Logger: &mock.Logger{},
		Former: &mock.Former{
			MapMock: mock.Map{Err: nil},
		},
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
	rec, req := createRequest(mockAchievementForm())

	env := model.Env{
		DB: &mock.DB{
			InvolvementExistsMock: mock.InvolvementExists{Bool: false, Err: nil},
		},
		Logger: &mock.Logger{},
		Former: &mock.Former{
			MapMock: mock.Map{Err: nil},
		},
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
	rec, req := createRequest(mockAchievementForm())

	env := model.Env{
		DB: &mock.DB{
			InvolvementExistsMock: mock.InvolvementExists{Bool: true, Err: nil},
			AchievementCreateMock: mock.AchievementCreate{ID: "", Err: errors.New("db error")},
		},
		Logger: &mock.Logger{},
		Former: &mock.Former{
			MapMock: mock.Map{Err: nil},
		},
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
	rec, req := createRequest(mockAchievementForm())

	env := model.Env{
		DB: &mock.DB{
			InvolvementExistsMock: mock.InvolvementExists{Bool: true, Err: nil},
			AchievementCreateMock: mock.AchievementCreate{ID: mockID, Err: nil},
		},
		Logger: &mock.Logger{},
		Former: &mock.Former{
			MapMock: mock.Map{Err: nil},
		},
	}

	handle := AchievementCreate
	statusCode := http.StatusOK

	testHandler(t, rec, req, &env, handle, statusCode)

	// Check the response body is what we expect.
	expected := `{"status":` + strconv.Itoa(statusCode) + `,"message":"` + fmt.Sprintf(formatCreated, achievement) + `","length":1,"results":"` + mockID + `"}`
	if rec.Body.String() != expected {
		t.Errorf("handler returned unexpected body: got %v want %v",
			rec.Body.String(), expected)
	}
}
