package auth

import (
	"app/middleware/app"
	"app/model"
	"app/shared/response"
	"net/http"
	"net/http/httptest"
	"testing"
)

func testHandler(env *model.Env, w http.ResponseWriter, r *http.Request) {
	response.Send(w, http.StatusOK, "ok", 1, "OK")
}

type mockDB struct{}

func (mdb *mockDB) Exists(table string, column string, value string) (bool, error) {
	return true, nil
}

func (mdb *mockDB) AchievementsAll() ([]*model.Achievement, error) {
	return make([]*model.Achievement, 0), nil
}

func (mdb *mockDB) UserCreate(firstName string, lastName string, email string, password string) (string, error) {
	return "user_id", nil
}

func (mdb *mockDB) UserAuth(email string, password string) (string, error) {
	return "user_id", nil
}

type mockToken struct{}

func (mtk *mockToken) Encrypt(token string) (string, error) {
	return "encrypted", nil
}

func (mtk *mockToken) Decrypt(encoded string) (string, error) {
	return "decrypted", nil
}

func TestAuthHandler(t *testing.T) {
	req := httptest.NewRequest("GET", "/auth", nil)
	req.Header.Add("auth_token", "asdf")

	rr := httptest.NewRecorder()

	env := &model.Env{
		DB:    &mockDB{},
		Token: &mockToken{},
	}

	appHandler := app.Handler{env, testHandler}

	var handler http.Handler = Handler(appHandler)

	handler.ServeHTTP(rr, req)

	// Check the status code is what we expect.
	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusOK)
	}

	// Check the response body is what we expect.
	expected := `{"status":200,"message":"ok","count":1,"results":"OK"}`
	if rr.Body.String() != expected {
		t.Errorf("handler returned unexpected body: got %v want %v",
			rr.Body.String(), expected)
	}
}
