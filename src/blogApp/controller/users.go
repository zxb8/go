package controller

import (
	"blogApp/database"
	"blogApp/models"
	"blogApp/repository/crud"
	"blogApp/responses"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

func GetUsers(w http.ResponseWriter,r *http.Request){
	w.Write([]byte("list Users"))
}

func CreateUser(w http.ResponseWriter,r *http.Request){
	body,err := ioutil.ReadAll(r.Body)
	if err != nil{
		responses.ERROR(w,http.StatusUnprocessableEntity,err)
		return
	}
	user := models.User{}
	err = json.Unmarshal(body,&user)

	if err != nil{
		responses.ERROR(w,http.StatusUnprocessableEntity,err)
		return
	}

	db,err :=database.Connect()

	if err != nil{
		responses.ERROR(w,http.StatusUnprocessableEntity,err)
		return
	}

	repo := crud.NewRepositoryUsersCRUD(db)

	user,err = repo.Save(user)
	if err != nil{
		responses.ERROR(w,http.StatusUnprocessableEntity,err)
		return
	}
	w.Header().Set("Location",fmt.Sprintf("%s%s%d",r.Host,r.RequestURI,user.ID))
	responses.JSON(w,http.StatusCreated,user)
}

func GetUser(w http.ResponseWriter,r *http.Request){
	w.Write([]byte("An Users"))
}

func UpdateUser(w http.ResponseWriter,r *http.Request){
	w.Write([]byte("update User"))
}
func DeleteUser(w http.ResponseWriter,r *http.Request){
	w.Write([]byte("delete User"))
}
