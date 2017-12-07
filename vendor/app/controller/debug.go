package controller

import (
	mauth "app/route/middleware/auth"
	"app/route/middleware/pprofhandler"
	"app/shared/router"

	"github.com/justinas/alice"
)

func init() {
	// Enable Pprof
	router.Get("/debug/pprof/*pprof", alice.
		New(mauth.Handler).
		ThenFunc(pprofhandler.Handler))
}
