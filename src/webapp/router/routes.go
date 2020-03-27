package router

import (
	"net/http"
	"webapp/controllers/home"
	"webapp/routes"
)

func Middleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		next.ServeHTTP(w, r)
	})
}

func GetRoutes() routes.Routes {
	return routes.Routes{
		routes.Route{"Home", "GET", "/index", home.Index},
	}
}
