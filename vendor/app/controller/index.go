package controller

import (
	"net/http"

	mauth "app/route/middleware/auth"
	"app/shared/response"
	"app/shared/router"

	"github.com/justinas/alice"
)

func init() {
	// Main page
	router.Get("/", alice.
		New(mauth.Handler).
		ThenFunc(Index))
}

func Index(w http.ResponseWriter, r *http.Request) {
	response.Send(w, http.StatusOK, "welcome", 0, nil)
}
