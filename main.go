package main

import (
	"log"

	"github.com/ivzb/achievers_server/app/controller"
	"github.com/ivzb/achievers_server/app/middleware/app"
	"github.com/ivzb/achievers_server/app/middleware/auth"
	"github.com/ivzb/achievers_server/app/middleware/logger"
	"github.com/ivzb/achievers_server/app/model"
	"github.com/ivzb/achievers_server/app/shared/config"
	"github.com/ivzb/achievers_server/app/shared/file"

	"net/http"
	"os"
	"runtime"
)

func init() {
	runtime.GOMAXPROCS(runtime.NumCPU())
}

func main() {
	confBytes, err := file.Read("config" + string(os.PathSeparator) + "config.json")

	if err != nil {
		log.Panic(err)
	}

	conf, err := config.New(confBytes)

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

	http.Handle("/achievement", authChain(env, controller.AchievementSingle))
	http.Handle("/achievements", authChain(env, controller.AchievementsIndex))

	http.Handle("/users/auth", anonChain(env, controller.UserAuth))
	http.Handle("/users/create", anonChain(env, controller.UserCreate))

	http.ListenAndServe(":8080", nil)
}

func authChain(env *model.Env, handler app.Handle) http.Handler {
	return use(app.Handler{env, handler}, auth.Handler, logger.Handler)
}

func anonChain(env *model.Env, handler app.Handle) http.Handler {
	return use(app.Handler{env, handler}, logger.Handler)
}

// specify middlewares in reverse order since it is chaining them recursively
func use(appHandler app.Handler, middlewares ...func(app.Handler) app.Handler) http.Handler {
	for _, middleware := range middlewares {
		appHandler = middleware(appHandler)
	}
	var handler http.Handler = appHandler

	return handler
}
