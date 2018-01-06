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
	form := model.NewFormer()

	env := &model.Env{
		DB:    db,
		Token: token,
		Log:   log,
		Form:  form,
	}

	log.Message("started@:8080")

	http.Handle("/", anonChain(env, controller.HomeIndex))

	http.Handle("/user/auth", anonChain(env, controller.UserAuth))
	http.Handle("/user/create", anonChain(env, controller.UserCreate))

	http.Handle("/achievement", authChain(env, controller.AchievementSingle))
	http.Handle("/achievements", authChain(env, controller.AchievementsIndex))
	http.Handle("/achievements/quest", authChain(env, controller.AchievementsByQuestID))
	http.Handle("/achievement/create", authChain(env, controller.AchievementCreate))

	http.Handle("/evidence", authChain(env, controller.EvidenceSingle))
	http.Handle("/evidences", authChain(env, controller.EvidencesIndex))
	http.Handle("/evidence/create", authChain(env, controller.EvidenceCreate))

	http.Handle("/reward", authChain(env, controller.RewardSingle))
	http.Handle("/rewards", authChain(env, controller.RewardsIndex))
	http.Handle("/reward/create", authChain(env, controller.RewardCreate))

	http.Handle("/quest", authChain(env, controller.QuestSingle))
	http.Handle("/quests", authChain(env, controller.QuestsIndex))
	http.Handle("/quest/create", authChain(env, controller.QuestCreate))

	http.Handle("/quest_achievement/create", authChain(env, controller.QuestAchievementCreate))

	http.ListenAndServe(":8080", nil)
}

func authChain(env *model.Env, handler app.Handler) http.Handler {
	return use(app.App{env, handler}, auth.Handler, logger.Handler)
}

func anonChain(env *model.Env, handler app.Handler) http.Handler {
	return use(app.App{env, handler}, logger.Handler)
}

// specify middlewares in reverse order since it is chaining them recursively
func use(app app.App, middlewares ...func(app.App) app.App) http.Handler {
	for _, middleware := range middlewares {
		app = middleware(app)
	}

	var handler http.Handler = app

	return handler
}
