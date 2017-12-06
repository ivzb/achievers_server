package router

import (
	"net/http"
	"github.com/justinas/alice"
	hr "app/route/middleware/httprouterwrapper"
	"app/route/middleware/auth"
)

// Delete is a shortcut for router.Handle("DELETE", path, handle)
func DeleteAuth(path string, fn http.HandlerFunc) {
	ri.Router.DELETE(path, hr.Handler(alice.
		New(auth.Handler).
		ThenFunc(fn)))
}

func DeleteAnon(path string, fn http.HandlerFunc) {
	ri.Router.DELETE(path, hr.Handler(alice.
		New().
		ThenFunc(fn)))
}

// Get is a shortcut for router.Handle("GET", path, handle)
func GetAuth(path string, fn http.HandlerFunc) {
	ri.Router.GET(path, hr.Handler(alice.
		New(auth.Handler).
		ThenFunc(fn)))
}

// Get is a shortcut for router.Handle("GET", path, handle)
func GetAnon(path string, fn http.HandlerFunc) {
	ri.Router.GET(path, hr.Handler(alice.
		New().
		ThenFunc(fn)))
}

// Head is a shortcut for router.Handle("HEAD", path, handle)
func HeadAuth(path string, fn http.HandlerFunc) {
	ri.Router.HEAD(path, hr.Handler(alice.
		New(auth.Handler).
		ThenFunc(fn)))
}

// Head is a shortcut for router.Handle("HEAD", path, handle)
func HeadAnon(path string, fn http.HandlerFunc) {
	ri.Router.HEAD(path, hr.Handler(alice.
		New().
		ThenFunc(fn)))
}

// Options is a shortcut for router.Handle("OPTIONS", path, handle)
func OptionsAuth(path string, fn http.HandlerFunc) {
	ri.Router.OPTIONS(path, hr.Handler(alice.
		New(auth.Handler).
		ThenFunc(fn)))
}

// Options is a shortcut for router.Handle("OPTIONS", path, handle)
func OptionsAnon(path string, fn http.HandlerFunc) {
	ri.Router.OPTIONS(path, hr.Handler(alice.
		New(auth.Handler).
		ThenFunc(fn)))
}

// Patch is a shortcut for router.Handle("PATCH", path, handle)
func PatchAuth(path string, fn http.HandlerFunc) {
	ri.Router.PATCH(path, hr.Handler(alice.
		New(auth.Handler).
		ThenFunc(fn)))
}

// Patch is a shortcut for router.Handle("PATCH", path, handle)
func PatchAnon(path string, fn http.HandlerFunc) {
	ri.Router.PATCH(path, hr.Handler(alice.
		New().
		ThenFunc(fn)))
}

// Post is a shortcut for router.Handle("POST", path, handle)
func PostAuth(path string, fn http.HandlerFunc) {
	ri.Router.POST(path, hr.Handler(alice.
		New(auth.Handler).
		ThenFunc(fn)))
}

// Post is a shortcut for router.Handle("POST", path, handle)
func PostAnon(path string, fn http.HandlerFunc) {
	ri.Router.POST(path, hr.Handler(alice.
		New().
		ThenFunc(fn)))
}

// Put is a shortcut for router.Handle("PUT", path, handle)
func PutAuth(path string, fn http.HandlerFunc) {
	ri.Router.PUT(path, hr.Handler(alice.
		New(auth.Handler).
		ThenFunc(fn)))
}

// Put is a shortcut for router.Handle("PUT", path, handle)
func PutAnon(path string, fn http.HandlerFunc) {
	ri.Router.PUT(path, hr.Handler(alice.
		New().
		ThenFunc(fn)))
}
