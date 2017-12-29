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

func TestEvidenceSingle_InvalidMethod(t *testing.T) {
	testInvalidMethod(t, "POST", "/evidence", EvidenceSingle)
}

func TestEvidenceSingle_MissingId(t *testing.T) {
	rec := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/evidence", nil)

	req.Form = url.Values{}
	req.Form.Add(id, "")

	handle := EvidenceSingle
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

func TestEvidenceSingle_EvidenceExistDBError(t *testing.T) {
	rec := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/evidence", nil)

	req.Form = url.Values{}
	req.Form.Add(id, "random_evidence_id")

	env := model.Env{
		DB: &mock.DB{
			EvidenceExistsMock: mock.EvidenceExists{Bool: false, Err: errors.New("db error")},
		},
		Logger: &mock.Logger{},
	}

	handle := EvidenceSingle
	statusCode := http.StatusInternalServerError

	testHandler(t, rec, req, &env, handle, statusCode)

	// Check the response body is what we expect.
	expected := fmt.Sprintf(`{"status":%d,"message":"%s"}`, statusCode, friendlyErrorMessage)
	if rec.Body.String() != expected {
		t.Errorf("handler returned unexpected body: got %v want %v",
			rec.Body.String(), expected)
	}
}

func TestEvidenceSingle_EvidenceDoesNotExist(t *testing.T) {
	rec := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/evidence", nil)

	req.Form = url.Values{}
	req.Form.Add(id, "random_evidence_id")

	env := model.Env{
		DB: &mock.DB{
			EvidenceExistsMock: mock.EvidenceExists{Bool: false, Err: nil},
		},
		Logger: &mock.Logger{},
	}

	handle := EvidenceSingle
	statusCode := http.StatusNotFound

	testHandler(t, rec, req, &env, handle, statusCode)

	// Check the response body is what we expect.
	expectedMessage := fmt.Sprintf(formatNotFound, evidence)
	expected := fmt.Sprintf(`{"status":%d,"message":"%s"}`, statusCode, expectedMessage)
	if rec.Body.String() != expected {
		t.Errorf("handler returned unexpected body: got %v want %v",
			rec.Body.String(), expected)
	}
}

func TestEvidenceSingle_EvidenceSingleDBError(t *testing.T) {
	rec := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/evidence", nil)

	req.Form = url.Values{}
	req.Form.Add(id, "random_evidence_id")

	env := model.Env{
		DB: &mock.DB{
			EvidenceExistsMock: mock.EvidenceExists{Bool: true, Err: nil},
			EvidenceSingleMock: mock.EvidenceSingle{Evd: nil, Err: errors.New("db error")},
		},
		Logger: &mock.Logger{},
	}

	handle := EvidenceSingle
	statusCode := http.StatusInternalServerError

	testHandler(t, rec, req, &env, handle, statusCode)

	// Check the response body is what we expect.
	expected := fmt.Sprintf(`{"status":%d,"message":"%s"}`, statusCode, friendlyErrorMessage)
	if rec.Body.String() != expected {
		t.Errorf("handler returned unexpected body: got %v want %v",
			rec.Body.String(), expected)
	}
}

func TestEvidenceSingle_EvidenceSingleNil(t *testing.T) {
	rec := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/evidence", nil)

	req.Form = url.Values{}
	req.Form.Add(id, "random_evidence_id")

	env := model.Env{
		DB: &mock.DB{
			EvidenceExistsMock: mock.EvidenceExists{Bool: true, Err: nil},
			EvidenceSingleMock: mock.EvidenceSingle{Evd: nil, Err: nil},
		},
		Logger: &mock.Logger{},
	}

	handle := EvidenceSingle
	statusCode := http.StatusNotFound

	testHandler(t, rec, req, &env, handle, statusCode)

	// Check the response body is what we expect.
	expectedMessage := fmt.Sprintf(formatNotFound, evidence)
	expected := fmt.Sprintf(`{"status":%d,"message":"%s"}`, statusCode, expectedMessage)
	if rec.Body.String() != expected {
		t.Errorf("handler returned unexpected body: got %v want %v",
			rec.Body.String(), expected)
	}
}

func TestEvidenceSingle_EvidenceSingle(t *testing.T) {
	rec := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/evidence", nil)

	req.Form = url.Values{}
	req.Form.Add(id, "random_evidence_id")

	env := model.Env{
		DB: &mock.DB{
			EvidenceExistsMock: mock.EvidenceExists{Bool: true, Err: nil},
			EvidenceSingleMock: mock.EvidenceSingle{Evd: mock.Evidence(), Err: nil},
		},
		Logger: &mock.Logger{},
	}

	handle := EvidenceSingle
	statusCode := http.StatusOK

	testHandler(t, rec, req, &env, handle, statusCode)

	// Check the response body is what we expect.
	expectedMessage := fmt.Sprintf(formatFound, evidence)
	marshalled, _ := json.Marshal(mock.Evidence())
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
	testInvalidMethod(t, "GET", "/evidence/create", EvidenceCreate)
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

func TestEvidenceCreate_MissingEvidenceId(t *testing.T) {
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

func TestEvidenceCreate_ValidEvidence(t *testing.T) {
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
