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

	http.Handle("/v1/user/auth", anonChain(env, controller.UserAuth))
	http.Handle("/v1/user/create", anonChain(env, controller.UserCreate))

	http.Handle("/v1/profile", authChain(env, controller.ProfileSingle))
	http.Handle("/v1/profile/me", authChain(env, controller.ProfileMe))

	http.Handle("/v1/achievement", authChain(env, controller.AchievementSingle))
	http.Handle("/v1/achievements/last", authChain(env, controller.AchievementsLast))
	http.Handle("/v1/achievements/after", authChain(env, controller.AchievementsAfter))
	http.Handle("/v1/achievements/quest/after", authChain(env, controller.AchievementsByQuestIDAfter))
	http.Handle("/v1/achievements/quest/last", authChain(env, controller.AchievementsByQuestIDLast))
	http.Handle("/v1/achievement/create", authChain(env, controller.AchievementCreate))

	http.Handle("/v1/evidence", authChain(env, controller.EvidenceSingle))
	http.Handle("/v1/evidences", authChain(env, controller.EvidencesIndex))
	http.Handle("/v1/evidence/create", authChain(env, controller.EvidenceCreate))

	http.Handle("/v1/reward", authChain(env, controller.RewardSingle))
	http.Handle("/v1/rewards", authChain(env, controller.RewardsIndex))
	http.Handle("/v1/reward/create", authChain(env, controller.RewardCreate))

	http.Handle("/v1/quest", authChain(env, controller.QuestSingle))
	http.Handle("/v1/quests", authChain(env, controller.QuestsIndex))
	http.Handle("/v1/quest/create", authChain(env, controller.QuestCreate))

	http.Handle("/v1/quest_achievement/create", authChain(env, controller.QuestAchievementCreate))

	http.Handle("/v1/file", authChain(env, controller.FileSingle))
	http.Handle("/v1/file/create", authChain(env, controller.FileCreate))

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
