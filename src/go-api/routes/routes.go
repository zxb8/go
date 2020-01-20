package routes

import (
	"github.com/gorilla/mux"
	"go-api/controllers"
)

func NewRouter() *mux.Router{
	r :=mux.NewRouter().StrictSlash(true)
	/**
	  users router
	 */
	r.HandleFunc("/users",controllers.PostUser).Methods("POST")
	r.HandleFunc("/users",controllers.GetUsers).Methods("GET")
	r.HandleFunc("/users/{id}",controllers.GetUser).Methods("GET")
	r.HandleFunc("/users/{id}",controllers.DeleteUser).Methods("DELETE")
	/**
	  post router
	 */
	r.HandleFunc("/posts",controllers.PostPost).Methods("POST")
	r.HandleFunc("/posts",controllers.GetPosts).Methods("GET")
	/**
	 feedbacks router
	 */
	r.HandleFunc("/feedbacks",controllers.PostFeedbacks).Methods("POST")
	r.HandleFunc("/feedbacks",controllers.GetFeedbacks).Methods("GET")
	return r
}

/*func home(w http.ResponseWriter,r *http.Request){
	w.Header().Set("Content-Type","application/json")
	utils.ToJson(w,struct{
		Message string `json:"message"`
	}{
		Message:"Go Restful Api !",
	})
}*/

