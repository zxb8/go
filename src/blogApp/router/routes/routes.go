package routes

import (
	"blogApp/middlewares"
	"github.com/gorilla/mux"
	"net/http"
)

type Route struct{
	Uri string
	Method string
	Handler func(http.ResponseWriter,*http.Request)
}
func Load() []Route{
	routes := usersRoutes
	return routes
}

func SetRoutes(r *mux.Router) * mux.Router{
	for _,route :=range Load(){
		r.HandleFunc(route.Uri,route.Handler).Methods(route.Method)
	}
	return r
}

func SetupRoutesWithMiddlewares(r *mux.Router) *mux.Router{
	for _,route := range Load(){
		r.HandleFunc(
			route.Uri,
			middlewares.SetMiddlewareLogger(
			middlewares.SetMiddlewareJSON(route.Handler))).Methods(route.Method)
	}
	return r
}