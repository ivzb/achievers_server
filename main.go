package main

import (
	"log"
	"net/http/pprof"
	"strconv"

	"github.com/ivzb/achievers_server/app/controller"
	"github.com/ivzb/achievers_server/app/db"
	"github.com/ivzb/achievers_server/app/middleware/app"
	"github.com/ivzb/achievers_server/app/middleware/auth"
	"github.com/ivzb/achievers_server/app/middleware/logger"
	"github.com/ivzb/achievers_server/app/shared/config"
	"github.com/ivzb/achievers_server/app/shared/env"
	"github.com/ivzb/achievers_server/app/shared/file"
	l "github.com/ivzb/achievers_server/app/shared/logger"
	"github.com/ivzb/achievers_server/app/shared/token"
	"github.com/ivzb/achievers_server/app/shared/uuid"

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

	db, err := db.NewDB(conf.Database)

	if err != nil {
		log.Panic(err)
	}

	token, err := token.NewTokener(conf.Token)

	if err != nil {
		log.Panic(err)
	}

	logger := l.NewLogger()

	uuid := uuid.NewUUID()

	env := &env.Env{
		DB:     db,
		Log:    logger,
		Token:  token,
		Config: conf,
		UUID:   uuid,
	}

	http.Handle("/", anonChain(env, controller.HomeIndex))

	http.HandleFunc("/debug/profile", pprof.Profile)

	http.Handle("/"+conf.Server.Version+"/user/auth", anonChain(env, controller.UserAuth))
	http.Handle("/"+conf.Server.Version+"/user/create", anonChain(env, controller.UserCreate))

	http.Handle("/"+conf.Server.Version+"/profile", authChain(env, controller.ProfileSingle))
	http.Handle("/"+conf.Server.Version+"/profile/me", authChain(env, controller.ProfileMe))

	http.Handle("/"+conf.Server.Version+"/achievement", authChain(env, controller.AchievementSingle))
	http.Handle("/"+conf.Server.Version+"/achievements/last", authChain(env, controller.AchievementsLast))
	http.Handle("/"+conf.Server.Version+"/achievements/after", authChain(env, controller.AchievementsAfter))
	http.Handle("/"+conf.Server.Version+"/achievements/quest/after", authChain(env, controller.AchievementsByQuestIDAfter))
	http.Handle("/"+conf.Server.Version+"/achievements/quest/last", authChain(env, controller.AchievementsByQuestIDLast))
	http.Handle("/"+conf.Server.Version+"/achievement/create", authChain(env, controller.AchievementCreate))

	http.Handle("/"+conf.Server.Version+"/evidence", authChain(env, controller.EvidenceSingle))
	http.Handle("/"+conf.Server.Version+"/evidences/last", authChain(env, controller.EvidencesLast))
	http.Handle("/"+conf.Server.Version+"/evidences/after", authChain(env, controller.EvidencesAfter))
	http.Handle("/"+conf.Server.Version+"/evidence/create", authChain(env, controller.EvidenceCreate))

	http.Handle("/"+conf.Server.Version+"/reward", authChain(env, controller.RewardSingle))
	http.Handle("/"+conf.Server.Version+"/rewards/last", authChain(env, controller.RewardsLast))
	http.Handle("/"+conf.Server.Version+"/rewards/after", authChain(env, controller.RewardsAfter))
	http.Handle("/"+conf.Server.Version+"/reward/create", authChain(env, controller.RewardCreate))

	http.Handle("/"+conf.Server.Version+"/quest", authChain(env, controller.QuestSingle))
	http.Handle("/"+conf.Server.Version+"/quests/last", authChain(env, controller.QuestsLast))
	http.Handle("/"+conf.Server.Version+"/quests/after", authChain(env, controller.QuestsAfter))
	http.Handle("/"+conf.Server.Version+"/quest/create", authChain(env, controller.QuestCreate))

	http.Handle("/"+conf.Server.Version+"/quest_achievement/create", authChain(env, controller.QuestAchievementCreate))

	http.Handle("/"+conf.Server.Version+"/file", authChain(env, controller.FileSingle))
	http.Handle("/"+conf.Server.Version+"/file/create", authChain(env, controller.FileCreate))

	port := strconv.Itoa(conf.Server.HTTPPort)
	logger.Message("started@:" + port)
	http.ListenAndServe(":"+port, nil)
}

func authChain(env *env.Env, handler app.Handler) http.Handler {
	return use(app.App{env, handler}, auth.Handler, logger.Handler)
}

func anonChain(env *env.Env, handler app.Handler) http.Handler {
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
