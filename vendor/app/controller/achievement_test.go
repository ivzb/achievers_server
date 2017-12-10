package controller

import (
	"app/model"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"
)

type mockDB struct{}

func (mdb *mockDB) AllAchievements() ([]*model.Achievement, error) {
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

	return achs, nil
}

func TestAchievementsIndex(t *testing.T) {
	rec := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/achievements", nil)

	env := model.Env{DB: &mockDB{}}
	http.Handler(AchievementsIndex(&env)).ServeHTTP(rec, req)

	expected := `[{"id":"fb7691eb-ea1d-b20f-edee-9cadcf23181f","title":"title","description":"desc","picture_url":"http://picture.jpg","involvement_id":"3","author_id":"4e69c9ba-719c-ca7c-fb66-80ab235c8e39","created_at":"2017-12-09T15:04:23Z","updated_at":"2017-12-09T15:04:23Z","deleted_at":"0000-01-01T00:00:00Z"},{"id":"93821a67-9c82-96e4-dc3c-423e5581d036","title":"another title","description":"another desc","picture_url":"http://another-picture.jpg","involvement_id":"1","author_id":"4e69c9ba-719c-ca7c-fb66-80ab235c8e39","created_at":"2017-12-09T15:04:23Z","updated_at":"2017-12-09T15:04:23Z","deleted_at":"0000-01-01T00:00:00Z"}]`

	if expected != rec.Body.String() {
		t.Errorf("\n...expected = %v\n...obtained = %v", expected, rec.Body.String())
	}
}
