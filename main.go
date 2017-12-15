package main

import (
	"app/controller"
	"app/middleware"
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

	http.Handle("/achievements", use(middleware.AppHandler{env, controller.AchievementsIndex}, middleware.AuthHandler))
	// lastly used
	// http.Handle("/achievements", use(middleware.Handler{env, controller.AchievementsIndex}, middleware.AuthHandler))

	// http.Handle("/achievements", middleware.Handler{env, middleware.AuthHandler(controller.AchievementsIndex)}))//authChain(env, controller.AchievementsIndex(env)))
	// http.HandleFunc("/achievements/show", showAchievement)
	// http.HandleFunc("/achievements/create", createAchievement)

	http.Handle("/users/create", anonChain(controller.UserCreate(env)))
	http.Handle("/users/auth", anonChain(controller.UserAuth(env)))

	http.ListenAndServe(":8080", nil)
}

// func authChain(env *model.Env, next http.Handler) http.Handler {
// 	return anonChain(auth.Handler(env, next))
// }

func anonChain(next http.Handler) http.Handler {
	return middleware.LoggerHandler(next)
}

func use(appHandler middleware.AppHandler, middlewares ...func(middleware.AppHandler) http.Handler) http.Handler {
	var handler http.Handler = appHandler
	for _, middleware := range middlewares {
		handler = middleware(appHandler)
	}

	return handler
}
