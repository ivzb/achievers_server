package router

import (
	"net/http"

	"github.com/gorilla/context"
	"github.com/julienschmidt/httprouter"

	"github.com/justinas/alice"
	hr "app/route/middleware/httprouterwrapper"
)

var (
	ri RouterInfo
)

const (
	authorize = "authorize"
	params = "params"
)

// RouteInfo is the details
type RouterInfo struct {
	Router *httprouter.Router
}

// *****************************************************************************
// Routes
// *****************************************************************************

// Set up the router
func init() {
	ri.Router = httprouter.New()
}

// ReadConfig returns the information
func ReadConfig() RouterInfo {
	return ri
}

// Instance returns the authorized router
func Instance() *httprouter.Router {
	return ri.Router
}

// Context returns the URL parameters
func Params(ri *http.Request) httprouter.Params {
	return context.Get(ri, params).(httprouter.Params)
}

// Chain returns handle with chaining using Alice
func Chain(fn http.HandlerFunc, c ...alice.Constructor) httprouter.Handle {
	return hr.Handler(alice.New(c...).ThenFunc(fn))
}