package routes

import (
	"apiLearn/models"
	"apiLearn/utils"
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"io/ioutil"
	"net/http"
	"strconv"
)

func NewRouter() *mux.Router{
	r := mux.NewRouter().StrictSlash(true)
	r.HandleFunc("/users",usersGetHandler).Methods("GET")
	r.HandleFunc("/users",usersPostHandler).Methods("POST")
	r.HandleFunc("/users/{id}",userGetHandler).Methods("GET")
	r.HandleFunc("/users/{id}",userPutHandler).Methods("PUT")
	r.HandleFunc("/users/{id}",userDeleteHandler).Methods("DELETE")
	r.HandleFunc("/login",loginPostHandler).Methods("POST")
	return r
}

func ContentTypeJson(w http.ResponseWriter){
	w.Header().Set("Content-Type","application/json")
}

func HttpInfo(r *http.Request){
	fmt.Printf("%s\t%s\t%s%s\t%s",r.Method,r.Proto,r.Host,r.URL,utils.GetDateTime())
}

func usersGetHandler(w http.ResponseWriter,r *http.Request){
	ContentTypeJson(w)
	HttpInfo(r)
	users,err := models.GetUsers()
	if err != nil{
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(struct{
			Error string
			Status int
		}{
			Error:"BAD REQUEST",
			Status:400,
		})
	}
	json.NewEncoder(w).Encode(users)
}

func usersPostHandler(w http.ResponseWriter,r *http.Request){
	ContentTypeJson(w)
	body,_:=ioutil.ReadAll(r.Body)
	var user models.User
	err := json.Unmarshal(body,&user)

	_,err  = models.NewUser(user)

	if err != nil{
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(struct{
			Error string
			Status int
		}{
			Error:"UNPROCESSBALE ENTITY",
			Status:422,
		})
	}

	json.NewEncoder(w).Encode(struct{
		Message string
		Status int
	}{
		Message:"Success",
		Status:200,
	})

}


func userGetHandler(w http.ResponseWriter,r *http.Request){
	params := mux.Vars(r)
	id,_ := strconv.Atoi(params["id"])

	user,err := models.GetUser(id)

	if err != nil{
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(struct{
			Error string
			Status int
		}{
			Error:"BAD REQUEST",
			Status:400,
		})
	}

	json.NewEncoder(w).Encode(user)

}


func userPutHandler(w http.ResponseWriter,r * http.Request){
	ContentTypeJson(w)
	HttpInfo(r)
	params := mux.Vars(r)
	id,_ := strconv.Atoi(params["id"])

	var user models.User
	user.Id = id

	body,_ :=ioutil.ReadAll(r.Body)
	err := json.Unmarshal(body,&user)

	if err != nil{
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(struct{
			Error string
			Status int
		}{
			Error:"UNPROCESSBALE ENTITY",
			Status:422,
		})
	}
	rows,err := models.UpdateUser(user)

	if err != nil{
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(struct{
			Error string
			Status int
		}{
			Error:"UNPROCESSBALE ENTITY",
			Status:422,
		})
	}

	json.NewEncoder(w).Encode(struct{RowsAffected int64}{RowsAffected:rows})

}

func userDeleteHandler(w http.ResponseWriter,r *http.Request){
	ContentTypeJson(w)
	HttpInfo(r)
	params := mux.Vars(r)
	id,_ := strconv.Atoi(params["id"])
	rows,err :=models.DeleteUser(id)
	if err != nil{
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(struct{
			Error string
			Status int
		}{
			Error:"BAD REQUEST",
			Status:400,
		})
		return
	}
	json.NewEncoder(w).Encode(struct{
		RowsAffected int64
	}{
		RowsAffected:rows,
	})
}


func loginPostHandler(w http.ResponseWriter,r *http.Request){
	ContentTypeJson(w)
	HttpInfo(r)
	body,_ :=ioutil.ReadAll(r.Body)
	var user models.User

	err := json.Unmarshal(body,&user)
	if err != nil{
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(struct{
			Error string
			Status int
		}{
			Error:"UNPROCESSBALE ENTITY",
			Status:422,
		})
		return
	}

	if user.Email =="" || user.Password ==""{
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(struct{
			Error string
			Status int
		}{
			Error:"UNAUTHORIZED",
			Status:401,
		})
	}

	json.NewEncoder(w).Encode(user)
}



