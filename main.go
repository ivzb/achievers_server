package main

import (
	"app/controller"
	"app/model"
	"log"
	"net/http"
	"runtime"
)

func init() {
	runtime.GOMAXPROCS(runtime.NumCPU())
}

func main() {
	db, err := model.NewDB("root:@/achievers?parseTime=true")
	if err != nil {
		log.Panic(err)
	}

	env := &model.Env{db}

	log.Println("started@:8080")

	http.Handle("/achievements", controller.AchievementsIndex(env))
	// http.HandleFunc("/achievements/show", showAchievement)
	// http.HandleFunc("/achievements/create", createAchievement)

	http.Handle("/users/create", controller.UserCreate(env))

	http.ListenAndServe(":8080", nil)
}
