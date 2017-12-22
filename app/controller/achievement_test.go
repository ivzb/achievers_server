package controller

import (
	"encoding/json"
	"errors"
	"net/http"
	"net/http/httptest"
	"strconv"
	"testing"
	"time"

	"github.com/ivzb/achievers_server/app/model"
)

func MockAchievements() []*model.Achievement {
	achs := make([]*model.Achievement, 0)

	achs = append(achs, &model.Achievement{
		"fb7691eb-ea1d-b20f-edee-9cadcf23181f",
		"title",
		"desc",
		"http://picture.jpg",
		"3",
		"4e69c9ba-719c-ca7c-fb66-80ab235c8e39",
		time.Date(2017, 12, 9, 15, 4, 23, 0, time.UTC),
		time.Date(2017, 12, 9, 15, 4, 23, 0, time.UTC),
		time.Date(0, 1, 1, 0, 0, 0, 0, time.UTC),
	})

	achs = append(achs, &model.Achievement{
		"93821a67-9c82-96e4-dc3c-423e5581d036",
		"another title",
		"another desc",
		"http://another-picture.jpg",
		"1",
		"4e69c9ba-719c-ca7c-fb66-80ab235c8e39",
		time.Date(2017, 12, 9, 15, 4, 23, 0, time.UTC),
		time.Date(2017, 12, 9, 15, 4, 23, 0, time.UTC),
		time.Date(0, 1, 1, 0, 0, 0, 0, time.UTC),
	})

	return achs
}

func TestAchievementsIndex_ValidAchievements(t *testing.T) {
	rec := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/achievements", nil)

	env := model.Env{
		DB: &model.DBMock{
			AchievementsAllMock: model.AchievementsAllMock{MockAchievements(), nil},
		},
		Logger: &model.LoggerMock{},
	}

	handle := AchievementsIndex
	statusCode := http.StatusOK

	testHandler(t, rec, req, &env, handle, statusCode)

	// Check the response body is what we expect.
	marshalled, _ := json.Marshal(MockAchievements())
	expected := `{"status":` + strconv.Itoa(statusCode) + `,"message":"item found","length":2,"results":` + string(marshalled) + `}`
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
	expected := `{"status":` + strconv.Itoa(statusCode) + `,"message":"` + FriendlyErrorMessage + `"}`
	if rec.Body.String() != expected {
		t.Errorf("handler returned unexpected body: got %v want %v",
			rec.Body.String(), expected)
	}
}
