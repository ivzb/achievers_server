package router

import (
	hr "app/route/middleware/httprouterwrapper"
	"net/http"
)

func Delete(path string, app http.Handler) {
	ri.Router.DELETE(path, hr.Handler(app))
}

func Get(path string, app http.Handler) {
	ri.Router.GET(path, hr.Handler(app))
}

func Head(path string, app http.Handler) {
	ri.Router.HEAD(path, hr.Handler(app))
}

func Options(path string, app http.Handler) {
	ri.Router.OPTIONS(path, hr.Handler(app))
}

func Patch(path string, app http.Handler) {
	ri.Router.PATCH(path, hr.Handler(app))
}

func Post(path string, app http.Handler) {
	ri.Router.POST(path, hr.Handler(app))
}

func Put(path string, app http.Handler) {
	ri.Router.PUT(path, hr.Handler(app))
}
