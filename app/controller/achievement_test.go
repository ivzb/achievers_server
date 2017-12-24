package controller

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"net/http/httptest"
	"strconv"
	"testing"

	"github.com/ivzb/achievers_server/app/model"
)

func TestAchievementsIndex_ValidAchievements(t *testing.T) {
	rec := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/achievements", nil)

	env := model.Env{
		DB: &model.DBMock{
			AchievementsAllMock: model.AchievementsAllMock{A: model.MockAchievements(), E: nil},
		},
		Logger: &model.LoggerMock{},
	}

	handle := AchievementsIndex
	statusCode := http.StatusOK

	testHandler(t, rec, req, &env, handle, statusCode)

	// Check the response body is what we expect.
	marshalled, _ := json.Marshal(model.MockAchievements())
	expected := `{"status":` + strconv.Itoa(statusCode) + `,"message":"` + fmt.Sprintf(formatFound, achievement) + `","length":2,"results":` + string(marshalled) + `}`
	if rec.Body.String() != expected {
		t.Errorf("handler returned unexpected body: got %v want %v",
			rec.Body.String(), expected)
	}
}

func TestAchievementsIndex_InvalidMethod(t *testing.T) {
	testInvalidMethod(t, "POST", "/achievements", AchievementsIndex)
}

func TestAchievementsIndex_DBError(t *testing.T) {
	rec := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/achievements", nil)

	env := model.Env{
		DB: &model.DBMock{
			AchievementsAllMock: model.AchievementsAllMock{nil, errors.New("db error")},
		},
		Logger: &model.LoggerMock{},
	}

	handle := AchievementsIndex
	statusCode := http.StatusInternalServerError

	testHandler(t, rec, req, &env, handle, statusCode)

	// Check the response body is what we expect.
	expected := `{"status":` + strconv.Itoa(statusCode) + `,"message":"` + friendlyErrorMessage + `"}`
	if rec.Body.String() != expected {
		t.Errorf("handler returned unexpected body: got %v want %v",
			rec.Body.String(), expected)
	}
}
