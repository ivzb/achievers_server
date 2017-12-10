package controller

import (
	"app/model"
	"encoding/json"
	"log"
	"net/http"
)

func UserCreate(env *model.Env) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		log.Println("/user/create")

		if r.Method != "POST" {
			http.Error(w, http.StatusText(405), 405)
			return
		}

		first_name := r.FormValue("first_name")
		last_name := r.FormValue("last_name")
		email := r.FormValue("email")
		password := r.FormValue("password")

		if first_name == "" {
			http.Error(w, "missing first_name", 400)
			return
		}

		if last_name == "" {
			http.Error(w, "missing last_name", 400)
			return
		}

		if email == "" {
			http.Error(w, "missing email", 400)
			return
		}

		if password == "" {
			http.Error(w, "missing password", 400)
			return
		}

		id, err := env.DB.UserCreate(first_name, last_name, email, password)

		if err != nil {
			if err == model.ErrExists {
				http.Error(w, "email already exists", 400)
				return
			}

			http.Error(w, http.StatusText(500), 500)
			return
		}

		js, err := json.Marshal(id)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		w.Header().Set("Content-Type", "application/json")
		w.Write(js)
	})
}
