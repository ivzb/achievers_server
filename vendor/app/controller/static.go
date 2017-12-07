package controller

import (
	"net/http"
	"strings"

	mauth "app/route/middleware/auth"
	"app/shared/router"

	"github.com/justinas/alice"
)

func init() {
	// Required so the trailing slash is not redirected
	router.Instance().RedirectTrailingSlash = false

	// Serve static files, no directory browsing
	router.Get("/static/*filepath", alice.
		New(mauth.Handler).
		ThenFunc(Static))
}

// Static maps static files
func Static(w http.ResponseWriter, r *http.Request) {
	// Disable listing directories
	if strings.HasSuffix(r.URL.Path, "/") {
		Error404(w, r)
		return
	}
	http.ServeFile(w, r, r.URL.Path[1:])
}
