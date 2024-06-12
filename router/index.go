package router

import (
	"net/http"
	"strings"
)

type Route struct {
	Method  string
	Pattern string
	Handler http.HandlerFunc
}

type Router struct {
	routes []Route
}

func NewRouter() *Router {
	return &Router{}
}

func (r *Router) AddRoute(method, pattern string, handler http.HandlerFunc) {
	r.routes = append(r.routes, Route{Method: method, Pattern: pattern, Handler: handler})
}

func (r *Router) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	for _, route := range r.routes {
		if strings.EqualFold(route.Method, req.Method) && route.Pattern == req.URL.Path {
			route.Handler(w, req)
			return
		}
	}
	http.NotFound(w, req)
}
