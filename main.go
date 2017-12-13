package main

import (
	"app/controller"
	"app/middleware/auth"
	"app/middleware/logger"
	"app/model"
	"app/shared/config"
	// "app/shared/token"

	"log"
	"net/http"
	"os"
	"runtime"
)

func init() {
	runtime.GOMAXPROCS(runtime.NumCPU())
}

type Handler struct {
	*model.Env
	H func(e *model.Env, w http.ResponseWriter, r *http.Request)
}

// ServeHTTP allows our Handler type to satisfy http.Handler.
func (h Handler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	h.H(h.Env, w, r)
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

	// handl := authChain(env, controller.AchievementsIndex(env))

	http.Handle("/achievements", authChain(env, controller.AchievementsIndex(env)))
	// http.HandleFunc("/achievements/show", showAchievement)
	// http.HandleFunc("/achievements/create", createAchievement)

	http.Handle("/users/create", anonChain(controller.UserCreate(env)))
	http.Handle("/users/auth", anonChain(controller.UserAuth(env)))

	http.ListenAndServe(":8080", nil)
}

func authChain(env *model.Env, next http.Handler) http.Handler {
	return anonChain(auth.Handler(env, next))
}

func anonChain(next http.Handler) http.Handler {
	return logger.Handler(next)
}
