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

func TestRewardsIndex_Fullpage(t *testing.T) {
	size := 9
	statusCode := http.StatusOK
	rec := requestRewards(t, size, statusCode)
	message := fmt.Sprintf(formatFound, rewards)
	results, _ := json.Marshal(mock.Rewards(size))

	expectRetrieve(t, rec, statusCode, message, results)
}

func TestRewardsIndex_Halfpage(t *testing.T) {
	size := 4
	statusCode := http.StatusOK
	rec := requestRewards(t, size, statusCode)
	message := fmt.Sprintf(formatFound, rewards)
	results, _ := json.Marshal(mock.Rewards(size))

	expectRetrieve(t, rec, statusCode, message, results)
}

func TestRewardsIndex_Emptypage(t *testing.T) {
	statusCode := http.StatusNotFound
	size := 0
	rec := requestRewards(t, size, statusCode)

	message := fmt.Sprintf(formatNotFound, page)
	expectCore(t, rec, statusCode, message)
}

func requestRewards(t *testing.T, size int, statusCode int) *httptest.ResponseRecorder {
	rec := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/Rewards", nil)

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

func TestRewardsIndex_InvalidMethod(t *testing.T) {
	testMethodNotAllowed(t, "POST", "/Rewards", RewardsIndex)
}

func TestRewardsIndex_Missingpage(t *testing.T) {
	rec := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/Rewards", nil)

	req.Form = url.Values{}
	req.Form.Add(page, "")

	handle := RewardsIndex
	statusCode := http.StatusBadRequest

	testHandler(t, rec, req, nil, handle, statusCode)

	message := fmt.Sprintf(formatMissing, page)
	expectCore(t, rec, statusCode, message)
}

func TestRewardsIndex_Invalidpage(t *testing.T) {
	rec := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/Rewards", nil)

	req.Form = url.Values{}
	req.Form.Add(page, "-1")

	handle := RewardsIndex
	statusCode := http.StatusBadRequest

	testHandler(t, rec, req, nil, handle, statusCode)

	message := fmt.Sprintf(formatInvalid, page)
	expectCore(t, rec, statusCode, message)
}

func TestRewardsIndex_DBError(t *testing.T) {
	rec := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/Rewards", nil)

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

	message := friendlyErrorMessage
	expectCore(t, rec, statusCode, message)
}

func TestRewardSingle_InvalidMethod(t *testing.T) {
	testMethodNotAllowed(t, "POST", "/Rewards", RewardSingle)
}

func TestRewardSingle_MissingId(t *testing.T) {
	rec := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/Reward", nil)

	req.Form = url.Values{}
	req.Form.Add(id, "")

	handle := RewardSingle
	statusCode := http.StatusBadRequest

	testHandler(t, rec, req, nil, handle, statusCode)

	message := fmt.Sprintf(formatMissing, id)
	expectCore(t, rec, statusCode, message)
}

func TestRewardSingle_RewardExistDBError(t *testing.T) {
	rec := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/Reward", nil)

	req.Form = url.Values{}
	req.Form.Add(id, "random_Reward_id")

	env := model.Env{
		DB: &mock.DB{
			RewardExistsMock: mock.RewardExists{Bool: false, Err: errors.New("db error")},
		},
		Logger: &mock.Logger{},
	}

	handle := RewardSingle
	statusCode := http.StatusInternalServerError

	testHandler(t, rec, req, &env, handle, statusCode)

	message := friendlyErrorMessage
	expectCore(t, rec, statusCode, message)
}

func TestRewardSingle_RewardDoesNotExist(t *testing.T) {
	rec := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/Reward", nil)

	req.Form = url.Values{}
	req.Form.Add(id, mockID)

	env := model.Env{
		DB: &mock.DB{
			RewardExistsMock: mock.RewardExists{Bool: false, Err: nil},
		},
		Logger: &mock.Logger{},
	}

	handle := RewardSingle
	statusCode := http.StatusNotFound

	testHandler(t, rec, req, &env, handle, statusCode)

	message := fmt.Sprintf(formatNotFound, reward)
	expectCore(t, rec, statusCode, message)
}

func TestRewardSingle_RewardSingleDBError(t *testing.T) {
	rec := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/Reward", nil)

	req.Form = url.Values{}
	req.Form.Add(id, "random_Reward_id")

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

	message := friendlyErrorMessage
	expectCore(t, rec, statusCode, message)
}

func TestRewardSingle_RewardSingleNil(t *testing.T) {
	rec := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/Reward", nil)

	req.Form = url.Values{}
	req.Form.Add(id, "random_Reward_id")

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

	message := fmt.Sprintf(formatNotFound, reward)
	expectCore(t, rec, statusCode, message)
}

func TestRewardSingle_RewardSingle(t *testing.T) {
	rec := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/Reward", nil)

	req.Form = url.Values{}
	req.Form.Add(id, "random_Reward_id")

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

	message := fmt.Sprintf(formatFound, reward)
	results, _ := json.Marshal(mock.Reward())
	expectRetrieve(t, rec, statusCode, message, results)
}

func mockRewardForm() url.Values {
	form := url.Values{}
	form.Add(title, mockTitle)
	form.Add(description, mockDescription)
	form.Add(pictureURL, mockPictureURL)
	form.Add(involvementID, mockInovlvementID)

	return form
}
