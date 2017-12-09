package main

// import (
// 	"encoding/json"
// 	"log"
// 	"os"
// 	"runtime"

// 	"app/controller"
// 	"app/route"
// 	"app/shared/database"
// 	"app/shared/email"
// 	"app/shared/jsonconfig"
// 	"app/shared/server"
// 	"app/shared/token"
// )

// // *****************************************************************************
// // Application Logic
// // *****************************************************************************

// func init() {
// 	// Verbose logging with file name and line number
// 	log.SetFlags(log.Lshortfile)

// 	// Use all CPU cores
// 	runtime.GOMAXPROCS(runtime.NumCPU())
// }

// func main() {
// 	// Load the configuration file
// 	jsonconfig.Load("config"+string(os.PathSeparator)+"config.json", config)

// 	// Configure the email settings
// 	email.Configure(config.Email)

// 	// Connect to database
// 	database.Connect(config.Database)

// 	// Load token key
// 	token.Configure(config.Token)

// 	// Load the controller routes
// 	controller.Load()

// 	// Start the listener
// 	server.Run(route.LoadHTTP(), route.LoadHTTPS(), config.Server)
// }

// // *****************************************************************************
// // Application Settings
// // *****************************************************************************

// // config the settings variable
// var config = &configuration{}

// // configuration contains the application settings
// type configuration struct {
// 	Database database.Info   `json:"Database"`
// 	Email    email.SMTPInfo  `json:"Email"`
// 	Server   server.Server   `json:"Server"`
// 	Token    token.TokenInfo `json:"Token"`
// }

// // ParseJSON unmarshals bytes to structs
// func (c *configuration) ParseJSON(b []byte) error {
// 	return json.Unmarshal(b, &c)
// }

import (
	"database/sql"
	"fmt"
	"crypto/rand"
	_ "github.com/go-sql-driver/mysql"
	"log"
	"net/http"
)

type Achievement struct {
	id             string
	title          string
	description    string
	picture_url    string
	involvement_id string
	author_id      string
	created_at     []uint8
	updated_at     []uint8
	deleted_at     []uint8
}

var db *sql.DB

func init() {
	var err error
	db, err = sql.Open("mysql", "root:@/achievers")
	if err != nil {
		log.Fatal(err)
	}

	if err = db.Ping(); err != nil {
		log.Fatal(err)
	}
}

func main() {
  log.Println("started@:8080")
	http.HandleFunc("/achievements", indexAchievements)
  http.HandleFunc("/achievements/show", showAchievement)
  http.HandleFunc("/achievements/create", createAchievement)
	http.ListenAndServe(":8080", nil)
}

func indexAchievements(w http.ResponseWriter, r *http.Request) {
  log.Println("/achievements")

	if r.Method != "GET" {
		http.Error(w, http.StatusText(405), 405)
		return
	}

	rows, err := db.Query("SELECT * FROM achievement")
	if err != nil {
		http.Error(w, http.StatusText(500), 500)
		return
	}
	defer rows.Close()

	achs := make([]*Achievement, 0)
	for rows.Next() {
		ach := new(Achievement)
		err := rows.Scan(
			&ach.id,
			&ach.title,
			&ach.description,
			&ach.picture_url,
			&ach.involvement_id,
			&ach.author_id,
			&ach.created_at,
			&ach.updated_at,
			&ach.deleted_at)

		if err != nil {
			http.Error(w, http.StatusText(500), 500)
			return
		}
		achs = append(achs, ach)
	}
	if err = rows.Err(); err != nil {
		http.Error(w, http.StatusText(500), 500)
		return
	}

	for _, ach := range achs {
		fmt.Fprintf(w, "%s, %s, %s, %s, %s, %s, %s, %s, %s\n", 
      ach.id,
			ach.title,
			ach.description,
			ach.picture_url,
			ach.involvement_id,
			ach.author_id,
			ach.created_at,
			ach.updated_at,
			ach.deleted_at)
	}
}

func showAchievement(w http.ResponseWriter, r *http.Request) {
  if r.Method != "GET" {
    http.Error(w, http.StatusText(405), 405)
    return
  }

  id := r.FormValue("id")
  if id == "" {
    http.Error(w, http.StatusText(400), 400)
    return
  }

  row := db.QueryRow("SELECT * FROM achievement WHERE id = ?", id)

  ach := new(Achievement)
  err := row.Scan(
			&ach.id,
			&ach.title,
			&ach.description,
			&ach.picture_url,
			&ach.involvement_id,
			&ach.author_id,
			&ach.created_at,
			&ach.updated_at,
			&ach.deleted_at)
    
  if err == sql.ErrNoRows {
    http.NotFound(w, r)
    return
  }
  
  if err != nil {
    http.Error(w, http.StatusText(500), 500)
    return
  }

  fmt.Fprintf(w, "%s, %s, %s, %s, %s, %s, %s, %s, %s\n", 
      ach.id,
			ach.title,
			ach.description,
			ach.picture_url,
			ach.involvement_id,
			ach.author_id,
			ach.created_at,
			ach.updated_at,
			ach.deleted_at)
}

func createAchievement(w http.ResponseWriter, r *http.Request) {
  if r.Method != "POST" {
    http.Error(w, http.StatusText(405), 405)
    return
  }

  title := r.FormValue("title")
  description := r.FormValue("description")
  picture_url := r.FormValue("picture_url")
  involvement_id := r.FormValue("involvement_id")
  author_id := r.FormValue("author_id")

  if title == "" || description == "" || picture_url == "" || involvement_id == "" || author_id == "" {
    http.Error(w, http.StatusText(400), 400)
    return
  }

  id, err := uuid()
	// If error on UUID generation
	if err != nil {
		http.Error(w, http.StatusText(500), 500)
    return
	}

  result, err := db.Exec(`INSERT INTO achievement (id, title, description, picture_url, involvement_id, author_id)
     VALUES(?, ?, ?, ?, ?, ?)`, 
     id, title, description, picture_url, involvement_id, author_id)
  
  if err != nil {
    log.Println(err)
    http.Error(w, http.StatusText(500), 500)
    return
  }

  rowsAffected, err := result.RowsAffected()
  if err != nil {
    http.Error(w, http.StatusText(500), 500)
    return
  }

  fmt.Fprintf(w, "Achievement %s created successfully (%d row affected)\n", id, rowsAffected)
}

func uuid() (string, error) {
	b := make([]byte, 16)
	_, err := rand.Read(b)
	if err != nil {
		return "", err
	}

	return fmt.Sprintf("%x-%x-%x-%x-%x", b[0:4], b[4:6], b[6:8], b[8:10], b[10:]), nil
}