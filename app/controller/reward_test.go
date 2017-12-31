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

func TestRewardsIndex_FullPage(t *testing.T) {
	statusCode := http.StatusOK
	rec := requestRewards(t, 9, statusCode)
	verifyCorrectRewardsResult(t, rec, 9)
}

func TestRewardsIndex_HalfPage(t *testing.T) {
	statusCode := http.StatusOK
	rec := requestRewards(t, 4, statusCode)
	verifyCorrectRewardsResult(t, rec, 4)
}

func TestRewardsIndex_EmptyPage(t *testing.T) {
	statusCode := http.StatusNotFound
	rec := requestRewards(t, 0, statusCode)

	expectedMessage := fmt.Sprintf(formatNotFound, page)
	expected := fmt.Sprintf(`{"status":%d,"message":"%s"}`, statusCode, expectedMessage)

	if rec.Body.String() != expected {
		t.Errorf("handler returned unexpected body: got %v want %v",
			rec.Body.String(), expected)
	}
}

func requestRewards(t *testing.T, size int, statusCode int) *httptest.ResponseRecorder {
	rec := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/rewards", nil)

	req.Form = url.Values{}
	req.Form.Add(page, "3")

	env := model.Env{
		DB: &mock.DB{
			RewardsAllMock: mock.RewardsAll{Rwds: mock.Rewards(size), Err: nil},
		},
		Logger: &mock.Logger{},
	}

	handle := RewardsIndex

	testHandler(t, rec, req, &env, handle, statusCode)

	return rec
}

func verifyCorrectRewardsResult(t *testing.T, rec *httptest.ResponseRecorder, size int) {
	expectedStatusCode := http.StatusOK
	expectedMessage := fmt.Sprintf(formatFound, rewards)
	mocks := mock.Rewards(size)
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

func TestRewardsIndex_InvalidMethod(t *testing.T) {
	testInvalidMethod(t, "POST", "/rewards", RewardsIndex)
}

func TestRewardsIndex_MissingPage(t *testing.T) {
	rec := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/rewards", nil)

	req.Form = url.Values{}
	req.Form.Add(page, "")

	handle := RewardsIndex
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

func TestRewardsIndex_InvalidPage(t *testing.T) {
	rec := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/rewards", nil)

	req.Form = url.Values{}
	req.Form.Add(page, "-1")

	handle := RewardsIndex
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

func TestRewardsIndex_DBError(t *testing.T) {
	rec := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/rewards", nil)

	req.Form = url.Values{}
	req.Form.Add(page, "3")

	env := model.Env{
		DB: &mock.DB{
			RewardsAllMock: mock.RewardsAll{Rwds: nil, Err: errors.New("db error")},
		},
		Logger: &mock.Logger{},
	}

	handle := RewardsIndex
	statusCode := http.StatusInternalServerError

	testHandler(t, rec, req, &env, handle, statusCode)

	// Check the response body is what we expect.
	expected := fmt.Sprintf(`{"status":%d,"message":"%s"}`, statusCode, friendlyErrorMessage)
	if rec.Body.String() != expected {
		t.Errorf("handler returned unexpected body: got %v want %v",
			rec.Body.String(), expected)
	}
}

func TestRewardSingle_InvalidMethod(t *testing.T) {
	testInvalidMethod(t, "POST", "/rewards", RewardSingle)
}

func TestRewardSingle_MissingId(t *testing.T) {
	rec := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/reward", nil)

	req.Form = url.Values{}
	req.Form.Add(id, "")

	handle := RewardSingle
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

func TestRewardSingle_RewardExistDBError(t *testing.T) {
	rec := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/reward", nil)

	req.Form = url.Values{}
	req.Form.Add(id, "random_reward_id")

	env := model.Env{
		DB: &mock.DB{
			RewardExistsMock: mock.RewardExists{Bool: false, Err: errors.New("db error")},
		},
		Logger: &mock.Logger{},
	}

	handle := RewardSingle
	statusCode := http.StatusInternalServerError

	testHandler(t, rec, req, &env, handle, statusCode)

	// Check the response body is what we expect.
	expected := fmt.Sprintf(`{"status":%d,"message":"%s"}`, statusCode, friendlyErrorMessage)
	if rec.Body.String() != expected {
		t.Errorf("handler returned unexpected body: got %v want %v",
			rec.Body.String(), expected)
	}
}

func TestRewardSingle_RewardDoesNotExist(t *testing.T) {
	rec := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/reward", nil)

	req.Form = url.Values{}
	req.Form.Add(id, "random_reward_id")

	env := model.Env{
		DB: &mock.DB{
			RewardExistsMock: mock.RewardExists{Bool: false, Err: nil},
		},
		Logger: &mock.Logger{},
	}

	handle := RewardSingle
	statusCode := http.StatusNotFound

	testHandler(t, rec, req, &env, handle, statusCode)

	// Check the response body is what we expect.
	expectedMessage := fmt.Sprintf(formatNotFound, reward)
	expected := fmt.Sprintf(`{"status":%d,"message":"%s"}`, statusCode, expectedMessage)
	if rec.Body.String() != expected {
		t.Errorf("handler returned unexpected body: got %v want %v",
			rec.Body.String(), expected)
	}
}

func TestRewardSingle_RewardSingleDBError(t *testing.T) {
	rec := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/reward", nil)

	req.Form = url.Values{}
	req.Form.Add(id, "random_reward_id")

	env := model.Env{
		DB: &mock.DB{
			RewardExistsMock: mock.RewardExists{Bool: true, Err: nil},
			RewardSingleMock: mock.RewardSingle{Rwd: nil, Err: errors.New("db error")},
		},
		Logger: &mock.Logger{},
	}

	handle := RewardSingle
	statusCode := http.StatusInternalServerError

	testHandler(t, rec, req, &env, handle, statusCode)

	// Check the response body is what we expect.
	expected := fmt.Sprintf(`{"status":%d,"message":"%s"}`, statusCode, friendlyErrorMessage)
	if rec.Body.String() != expected {
		t.Errorf("handler returned unexpected body: got %v want %v",
			rec.Body.String(), expected)
	}
}

func TestRewardSingle_RewardSingleNil(t *testing.T) {
	rec := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/reward", nil)

	req.Form = url.Values{}
	req.Form.Add(id, "random_reward_id")

	env := model.Env{
		DB: &mock.DB{
			RewardExistsMock: mock.RewardExists{Bool: true, Err: nil},
			RewardSingleMock: mock.RewardSingle{Rwd: nil, Err: nil},
		},
		Logger: &mock.Logger{},
	}

	handle := RewardSingle
	statusCode := http.StatusNotFound

	testHandler(t, rec, req, &env, handle, statusCode)

	// Check the response body is what we expect.
	expectedMessage := fmt.Sprintf(formatNotFound, reward)
	expected := fmt.Sprintf(`{"status":%d,"message":"%s"}`, statusCode, expectedMessage)
	if rec.Body.String() != expected {
		t.Errorf("handler returned unexpected body: got %v want %v",
			rec.Body.String(), expected)
	}
}

func TestRewardSingle_RewardSingle(t *testing.T) {
	rec := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/reward", nil)

	req.Form = url.Values{}
	req.Form.Add(id, "random_reward_id")

	env := model.Env{
		DB: &mock.DB{
			RewardExistsMock: mock.RewardExists{Bool: true, Err: nil},
			RewardSingleMock: mock.RewardSingle{Rwd: mock.Reward(), Err: nil},
		},
		Logger: &mock.Logger{},
	}

	handle := RewardSingle
	statusCode := http.StatusOK

	testHandler(t, rec, req, &env, handle, statusCode)

	// Check the response body is what we expect.
	expectedMessage := fmt.Sprintf(formatFound, reward)
	marshalled, _ := json.Marshal(mock.Reward())
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

func mockRewardForm() url.Values {
	form := url.Values{}
	form.Add(title, mockTitle)
	form.Add(description, mockDescription)
	form.Add(pictureURL, mockPictureURL)
	form.Add(involvementID, mockInovlvementID)

	return form
}
