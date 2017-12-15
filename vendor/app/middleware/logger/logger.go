package logger

import (
	"app/middleware/app"
	"app/model"
	"fmt"
	"net/http"
	"time"
)

// Handler will log the HTTP requests
func Handler(handler app.Handler) http.Handler {
	return app.Handler{handler.Env, func(env *model.Env, w http.ResponseWriter, r *http.Request) {
		fmt.Println(time.Now().Format("2006-01-02 03:04:05 PM"), r.RemoteAddr, r.Method, r.URL)
		handler.ServeHTTP(w, r)
	}}
}
