package routes

import (
	"blogApp/controller"
	"net/http"
)

var usersRoutes = []Route{
	Route{
		Uri:"/users",
		Method:http.MethodGet,
		Handler:controller.GetUsers,
	},
	Route{
		Uri:"/users",
		Method:http.MethodPost,
		Handler:controller.CreateUser,
	},
	Route{
		Uri:"/users/{id}",
		Method:http.MethodGet,
		Handler:controller.GetUser,
	},
	Route{
		Uri:"/users/{id}",
		Method:http.MethodPut,
		Handler:controller.UpdateUser,
	},
	Route{
		Uri:"/users/{id}",
		Method:http.MethodGet,
		Handler:controller.DeleteUser,
	},
}
