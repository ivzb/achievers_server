package main

import (
	"app/controller"
	"app/middleware/app"
	"app/middleware/auth"
	"app/middleware/logger"
	"app/model"
	"app/shared/config"
	"log"

	"net/http"
	"os"
	"runtime"
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

	token, err := model.NewTokener(conf.Token)

	if err != nil {
		log.Panic(err)
	}

	log := model.NewLogger()

	env := &model.Env{
		DB:      db,
		Tokener: token,
		Logger:  log,
	}

	log.Log("started@:8080")

	http.Handle("/achievements", use(app.Handler{env, controller.AchievementsIndex}, auth.Handler, logger.Handler))

	http.Handle("/users/create", use(app.Handler{env, controller.UserCreate}, logger.Handler))
	http.Handle("/users/auth", use(app.Handler{env, controller.UserAuth}, logger.Handler))

	http.ListenAndServe(":8080", nil)
}

func use(appHandler app.Handler, middlewares ...func(app.Handler) app.Handler) http.Handler {
	for _, middleware := range middlewares {
		appHandler = middleware(appHandler)
	}
	var handler http.Handler = appHandler

	return handler
}
