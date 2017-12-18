package controller

import (
	"app/model"
	"app/shared/response"
	"net/http"
)

func AchievementsIndex(env *model.Env, w http.ResponseWriter, r *http.Request) response.Message {
	if r.Method != "GET" {
		return response.SendError(w, http.StatusMethodNotAllowed, MethodNotAllowedErrorMessage)
	}

	uID := env.UserId
	env.Logger.Log(uID)

	achs, err := env.DB.AchievementsAll()
	if err != nil {
		return response.SendError(w, http.StatusInternalServerError, FriendlyErrorMessage)
	}

	return response.Send(w, http.StatusOK, ItemFound, len(achs), achs)
}

// func showAchievement(w http.ResponseWriter, r *http.Request) {
//   if r.Method != "GET" {
//     http.Error(w, http.StatusText(405), 405)
//     return
//   }

//   id := r.FormValue("id")
//   if id == "" {
//     http.Error(w, http.StatusText(400), 400)
//     return
//   }

//   row := db.QueryRow("SELECT * FROM achievement WHERE id = ?", id)

//   ach := new(Achievement)
//   err := row.Scan(
// 			&ach.id,
// 			&ach.title,
// 			&ach.description,
// 			&ach.picture_url,
// 			&ach.involvement_id,
// 			&ach.author_id,
// 			&ach.created_at,
// 			&ach.updated_at,
// 			&ach.deleted_at)

//   if err == sql.ErrNoRows {
//     http.NotFound(w, r)
//     return
//   }

//   if err != nil {
//     http.Error(w, http.StatusText(500), 500)
//     return
//   }

//   fmt.Fprintf(w, "%s, %s, %s, %s, %s, %s, %s, %s, %s\n",
//       ach.id,
// 			ach.title,
// 			ach.description,
// 			ach.picture_url,
// 			ach.involvement_id,
// 			ach.author_id,
// 			ach.created_at,
// 			ach.updated_at,
// 			ach.deleted_at)
// }

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
