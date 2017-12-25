package controller

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/ivzb/achievers_server/app/model"
	"github.com/ivzb/achievers_server/app/shared/response"
)

const (
	id           = "id"
	achievement  = "achievement"
	achievements = "achievements"
	page         = "page"
)

func AchievementsIndex(
	env *model.Env,
	w http.ResponseWriter,
	r *http.Request) response.Message {

	if r.Method != "GET" {
		return response.MethodNotAllowed(methodNotAllowed)
	}

	pg, err := strconv.Atoi(r.FormValue("page"))

	if err != nil {
		return response.BadRequest(fmt.Sprintf(formatMissing, page))
	}

	if pg < 0 {
		return response.BadRequest(fmt.Sprintf(formatInvalid, page))
	}

	uID := env.UserId
	env.Logger.Log(uID)

	achs, err := env.DB.AchievementsAll(pg)

	if err != nil {
		return response.InternalServerError(friendlyErrorMessage)
	}

	if len(achs) == 0 {
		return response.NotFound(fmt.Sprintf(formatNotFound, page))
	}

	return response.Ok(
		fmt.Sprintf(formatFound, achievements),
		len(achs),
		achs)
}

func AchievementSingle(
	env *model.Env,
	w http.ResponseWriter,
	r *http.Request) response.Message {

	if r.Method != "GET" {
		return response.MethodNotAllowed(methodNotAllowed)
	}

	achID := r.FormValue(id)

	if achID == "" {
		return response.BadRequest(fmt.Sprintf(formatMissing, id))
	}

	exists, err := env.DB.Exists(achievement, id, achID)

	if err != nil {
		return response.InternalServerError(friendlyErrorMessage)
	}

	if !exists {
		return response.NotFound(fmt.Sprintf(formatNotFound, achievement))
	}

	ach, err := env.DB.AchievementSingle(achID)

	if err != nil {
		return response.InternalServerError(friendlyErrorMessage)
	}

	if ach == nil {
		return response.NotFound(fmt.Sprintf(formatNotFound, achievement))
	}

	return response.Ok(
		fmt.Sprintf(formatFound, achievement),
		1,
		ach)
}

// func createAchievement(w http.ResponseWriter, r *http.Request) {
//   if r.Method != "POST" {
//     http.Error(w, http.StatusText(405), 405)
//     return
//   }

//   title := r.FormValue("title")
//   description := r.FormValue("description")
//   picture_url := r.FormValue("picture_url")
//   involvement_id := r.FormValue("involvement_id")
//   author_id := r.FormValue("author_id")

//   if title == "" || description == "" || picture_url == "" || involvement_id == "" || author_id == "" {
//     http.Error(w, http.StatusText(400), 400)
//     return
//   }

//   id, err := uuid()
// 	// If error on UUID generation
// 	if err != nil {
// 		http.Error(w, http.StatusText(500), 500)
//     return
// 	}

//   result, err := db.Exec(`INSERT INTO achievement (id, title, description, picture_url, involvement_id, author_id)
//      VALUES(?, ?, ?, ?, ?, ?)`,
//      id, title, description, picture_url, involvement_id, author_id)

//   if err != nil {
//     log.Println(err)
//     http.Error(w, http.StatusText(500), 500)
//     return
//   }

//   rowsAffected, err := result.RowsAffected()
//   if err != nil {
//     http.Error(w, http.StatusText(500), 500)
//     return
//   }

//   fmt.Fprintf(w, "Achievement %s created successfully (%d row affected)\n", id, rowsAffected)
// }

// func uuid() (string, error) {
// 	b := make([]byte, 16)
// 	_, err := rand.Read(b)
// 	if err != nil {
// 		return "", err
// 	}

// 	return fmt.Sprintf("%x-%x-%x-%x-%x", b[0:4], b[4:6], b[6:8], b[8:10], b[10:]), nil
// }
