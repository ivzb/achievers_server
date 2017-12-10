package main

import (
	"app/controller"
	"app/model"
	"github.com/justinas/alice"
	"log"
	"net/http"
)

func main() {
	db, err := model.NewDB("root:@/achievers?parseTime=true")
	if err != nil {
		log.Panic(err)
	}

	env := &model.Env{db}

	stdChain := alice.New( /*myLoggingHandler, authHandler, enforceJSONHandler*/ )

	log.Println("started@:8080")

	http.Handle("/achievements", stdChain.Then(controller.AchievementsIndex(env)))
	// http.HandleFunc("/achievements/show", showAchievement)
	// http.HandleFunc("/achievements/create", createAchievement)
	http.ListenAndServe(":8080", nil)
}
