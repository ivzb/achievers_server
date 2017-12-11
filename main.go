package main

import (
	"app/controller"
	"app/middleware/auth"
	"app/middleware/logger"
	"app/model"
	"app/shared/config"
	"app/shared/token"

	"log"
	"net/http"
	"os"
	"runtime"

	"github.com/justinas/alice"
)

func init() {
	runtime.GOMAXPROCS(runtime.NumCPU())
}

func main() {
	conf, err := config.Load("config" + string(os.PathSeparator) + "config.json")

	if err != nil {
		log.Panic(err)
	}

	db, err := model.NewDB(conf.Database)

	if err != nil {
		log.Panic(err)
	}

	ctx := model.NewContext()

	env := &model.Env{db, ctx}

	// Load token key
	// todo: move token to model and put it inside env.Context
	token.Configure(conf.Token)

	// authChain := alice.New(logger.Handler, auth.Handler)
	anonChain := alice.New(logger.Handler)

	log.Println("started@:8080")

	http.Handle("/achievements", logger.Handler(auth.Handler(env, controller.AchievementsIndex(env)))) //authChain.Then(controller.AchievementsIndex(env)))
	// http.HandleFunc("/achievements/show", showAchievement)
	// http.HandleFunc("/achievements/create", createAchievement)

	http.Handle("/users/create", anonChain.Then(controller.UserCreate(env)))
	http.Handle("/users/auth", anonChain.Then(controller.UserAuth(env)))

	http.ListenAndServe(":8080", nil)
}
