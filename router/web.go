package router

import (
	"fmt"
	"net/http"
)

func RegisterRoute(r *Router) {

	r.AddRoute("GET", "/", func(w http.ResponseWriter, r *http.Request) { fmt.Fprintf(w, "dunno man") })
}
