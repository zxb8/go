package main

import (
	"apiLearn/routes"
	"encoding/json"
	"net/http"
)

func main(){
	//r := mux.NewRouter().StrictSlash(true)
	//r.HandleFunc("/",handler).Methods("GET")
	r := routes.NewRouter()
	http.ListenAndServe(":3000",r)
}

func ContentTypeJson(w http.ResponseWriter){
	w.Header().Set("Content-Type","application/json")

}

func handler(w http.ResponseWriter,r *http.Request){
	ContentTypeJson(w)
	json.NewEncoder(w).Encode(struct{Message string}{Message:"hello world!",})

}
