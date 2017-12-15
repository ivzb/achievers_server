package main

import (
	"app/controller"
	"app/middleware/app"
	"app/middleware/auth"
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

	token, err := model.NewToken(conf.Token)

	if err != nil {
		log.Panic(err)
	}

	env := &model.Env{
		DB:    db,
		Token: token,
	}

	log.Println("started@:8080")

	http.Handle("/achievements", use(app.Handler{env, controller.AchievementsIndex}, auth.Handler))

	http.Handle("/users/create", use(app.Handler{env, controller.UserCreate}))
	http.Handle("/users/auth", use(app.Handler{env, controller.UserAuth}))

	http.ListenAndServe(":8080", nil)
}

func use(appHandler app.Handler, middlewares ...func(app.Handler) http.Handler) http.Handler {
	var handler http.Handler = appHandler
	for _, middleware := range middlewares {
		handler = middleware(appHandler)
	}

	return handler
}
