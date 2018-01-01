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

func TestAchievementsIndex_FullPage(t *testing.T) {
	size := 9
	statusCode := http.StatusOK
	rec := requestAchievements(t, size, statusCode)
	message := fmt.Sprintf(formatFound, achievements)
	results, _ := json.Marshal(mock.Achievements(size))

	expectRetrieve(t, rec, statusCode, message, results)
}

func TestAchievementsIndex_HalfPage(t *testing.T) {
	size := 4
	statusCode := http.StatusOK
	rec := requestAchievements(t, size, statusCode)
	message := fmt.Sprintf(formatFound, achievements)
	results, _ := json.Marshal(mock.Achievements(size))

	expectRetrieve(t, rec, statusCode, message, results)
}

func TestAchievementsIndex_EmptyPage(t *testing.T) {
	statusCode := http.StatusNotFound
	size := 0
	rec := requestAchievements(t, size, statusCode)

	message := fmt.Sprintf(formatNotFound, page)
	expectCore(t, rec, statusCode, message)
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

func TestAchievementsIndex_InvalidMethod(t *testing.T) {
	testMethodNotAllowed(t, "POST", "/achievements", AchievementsIndex)
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
	testMethodNotAllowed(t, "POST", "/achievement", AchievementSingle)
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

	message := fmt.Sprintf(formatFound, achievement)
	results, _ := json.Marshal(mock.Achievement())
	expectRetrieve(t, rec, statusCode, message, results)
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
	testMethodNotAllowed(t, "GET", "/achievement/create", AchievementCreate)
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

	message := fmt.Sprintf(formatCreated, achievement)
	results, _ := json.Marshal(mockID)
	expectRetrieve(t, rec, statusCode, message, results)
}
