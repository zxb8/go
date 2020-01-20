package router

import (
	"blogApp/router/routes"
	"github.com/gorilla/mux"
)

func New() *mux.Router{
	r := mux.NewRouter().StrictSlash(true)
	return routes.SetRoutes(r)
}
