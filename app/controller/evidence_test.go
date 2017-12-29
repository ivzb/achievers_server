package controller

import (
	"errors"
	"fmt"
	"net/http"
	"net/url"
	"strconv"
	"testing"

	"github.com/ivzb/achievers_server/app/model"
	"github.com/ivzb/achievers_server/app/model/mock"
)

func mockEvidenceForm() url.Values {
	form := url.Values{}
	form.Add(description, mockDescription)
	form.Add(previewURL, mockPreviewURL)
	form.Add(_url, mockURL)
	form.Add(multimediaTypeID, mockMultimediaTypeID)
	form.Add(achievementID, mockAchievementID)

	return form
}

func TestEvidenceCreate_InvalidMethod(t *testing.T) {
	testInvalidMethod(t, "GET", "/achievement/create", EvidenceCreate)
}

func TestEvidenceCreate_FormMapError(t *testing.T) {
	rec, req := createRequest(nil)
	statusCode := http.StatusBadRequest
	handle := EvidenceCreate

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

func TestEvidenceCreate_MissingDescription(t *testing.T) {
	testMissingFormValue(t, EvidenceCreate, mockEvidenceForm(), description)
}

func TestEvidenceCreate_MissingPreviewUrl(t *testing.T) {
	testMissingFormValue(t, EvidenceCreate, mockEvidenceForm(), previewURL)
}

func TestEvidenceCreate_MissingUrl(t *testing.T) {
	testMissingFormValue(t, EvidenceCreate, mockEvidenceForm(), _url)
}

func TestEvidenceCreate_MissingMultimediaTypeId(t *testing.T) {
	testMissingFormValue(t, EvidenceCreate, mockEvidenceForm(), multimediaTypeID)
}

func TestEvidenceCreate_MissingAchievementId(t *testing.T) {
	testMissingFormValue(t, EvidenceCreate, mockEvidenceForm(), achievementID)
}

func TestEvidenceCreate_MultimediaTypeIdExistDBError(t *testing.T) {
	rec, req := createRequest(mockEvidenceForm())

	env := model.Env{
		DB: &mock.DB{
			MultimediaTypeExistsMock: mock.MultimediaTypeExists{Bool: false, Err: errors.New("db error")},
		},
		Logger: &mock.Logger{},
		Former: &mock.Former{
			MapMock: mock.Map{Err: nil},
		},
	}

	handle := EvidenceCreate
	statusCode := http.StatusInternalServerError

	testHandler(t, rec, req, &env, handle, statusCode)

	// Check the response body is what we expect.
	expected := fmt.Sprintf(`{"status":%d,"message":"%s"}`, statusCode, friendlyErrorMessage)
	if rec.Body.String() != expected {
		t.Errorf("handler returned unexpected body: got %v want %v",
			rec.Body.String(), expected)
	}
}

func TestEvidenceCreate_MultimediaTypeIdDoesNotExist(t *testing.T) {
	rec, req := createRequest(mockEvidenceForm())

	env := model.Env{
		DB: &mock.DB{
			MultimediaTypeExistsMock: mock.MultimediaTypeExists{Bool: false, Err: nil},
		},
		Logger: &mock.Logger{},
		Former: &mock.Former{
			MapMock: mock.Map{Err: nil},
		},
	}

	handle := EvidenceCreate
	statusCode := http.StatusBadRequest

	testHandler(t, rec, req, &env, handle, statusCode)

	// Check the response body is what we expect.
	expected := fmt.Sprintf(`{"status":%d,"message":"%s"}`, statusCode, fmt.Sprintf(formatNotFound, multimediaType))
	if rec.Body.String() != expected {
		t.Errorf("handler returned unexpected body: got %v want %v",
			rec.Body.String(), expected)
	}
}

func TestEvidenceCreate_AchievementIdExistDBError(t *testing.T) {
	rec, req := createRequest(mockEvidenceForm())

	env := model.Env{
		DB: &mock.DB{
			MultimediaTypeExistsMock: mock.MultimediaTypeExists{Bool: true, Err: nil},
			AchievementExistsMock:    mock.AchievementExists{Bool: false, Err: errors.New("db error")},
		},
		Logger: &mock.Logger{},
		Former: &mock.Former{
			MapMock: mock.Map{Err: nil},
		},
	}

	handle := EvidenceCreate
	statusCode := http.StatusInternalServerError

	testHandler(t, rec, req, &env, handle, statusCode)

	// Check the response body is what we expect.
	expected := fmt.Sprintf(`{"status":%d,"message":"%s"}`, statusCode, friendlyErrorMessage)
	if rec.Body.String() != expected {
		t.Errorf("handler returned unexpected body: got %v want %v",
			rec.Body.String(), expected)
	}
}

func TestEvidenceCreate_AchievementIdDoesNotExist(t *testing.T) {
	rec, req := createRequest(mockEvidenceForm())

	env := model.Env{
		DB: &mock.DB{
			MultimediaTypeExistsMock: mock.MultimediaTypeExists{Bool: true, Err: nil},
			AchievementExistsMock:    mock.AchievementExists{Bool: false, Err: nil},
		},
		Logger: &mock.Logger{},
		Former: &mock.Former{
			MapMock: mock.Map{Err: nil},
		},
	}

	handle := EvidenceCreate
	statusCode := http.StatusBadRequest

	testHandler(t, rec, req, &env, handle, statusCode)

	// Check the response body is what we expect.
	expected := fmt.Sprintf(`{"status":%d,"message":"%s"}`, statusCode, fmt.Sprintf(formatNotFound, achievement))
	if rec.Body.String() != expected {
		t.Errorf("handler returned unexpected body: got %v want %v",
			rec.Body.String(), expected)
	}
}

func TestEvidenceCreate_EvidenceCreateDBError(t *testing.T) {
	rec, req := createRequest(mockEvidenceForm())

	env := model.Env{
		DB: &mock.DB{
			MultimediaTypeExistsMock: mock.MultimediaTypeExists{Bool: true, Err: nil},
			AchievementExistsMock:    mock.AchievementExists{Bool: true, Err: nil},
			EvidenceCreateMock:       mock.EvidenceCreate{ID: "", Err: errors.New("db error")},
		},
		Logger: &mock.Logger{},
		Former: &mock.Former{
			MapMock: mock.Map{Err: nil},
		},
	}

	handle := EvidenceCreate
	statusCode := http.StatusInternalServerError

	testHandler(t, rec, req, &env, handle, statusCode)

	// Check the response body is what we expect.
	expected := fmt.Sprintf(`{"status":%d,"message":"%s"}`, statusCode, friendlyErrorMessage)
	if rec.Body.String() != expected {
		t.Errorf("handler returned unexpected body: got %v want %v",
			rec.Body.String(), expected)
	}
}

func TestEvidenceCreate_ValidAchievement(t *testing.T) {
	rec, req := createRequest(mockEvidenceForm())

	env := model.Env{
		DB: &mock.DB{
			MultimediaTypeExistsMock: mock.MultimediaTypeExists{Bool: true, Err: nil},
			AchievementExistsMock:    mock.AchievementExists{Bool: true, Err: nil},
			EvidenceCreateMock:       mock.EvidenceCreate{ID: mockID, Err: nil},
		},
		Logger: &mock.Logger{},
		Former: &mock.Former{
			MapMock: mock.Map{Err: nil},
		},
	}

	handle := EvidenceCreate
	statusCode := http.StatusOK

	testHandler(t, rec, req, &env, handle, statusCode)

	// Check the response body is what we expect.
	expected := `{"status":` + strconv.Itoa(statusCode) + `,"message":"` + fmt.Sprintf(formatCreated, evidence) + `","length":1,"results":"` + mockID + `"}`
	if rec.Body.String() != expected {
		t.Errorf("handler returned unexpected body: got %v want %v",
			rec.Body.String(), expected)
	}
}
