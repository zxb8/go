package controllers

import (
	"encoding/json"
	"github.com/gorilla/mux"
	"go-api/models"
	"go-api/utils"
	"net/http"
	"strconv"
)

func PostUser(w http.ResponseWriter,r *http.Request){
	body :=utils.BodyParser(r)
	var user models.User
	err := json.Unmarshal(body,&user)
	if err != nil{
		utils.ToJson(w,err.Error(),http.StatusUnprocessableEntity)
		return
	}
	err = models.NewUser(user)
	if err != nil{
		utils.ToJson(w,err.Error(),http.StatusUnprocessableEntity)
		return
	}
	utils.ToJson(w,"success",http.StatusCreated)
}

func GetUsers(w http.ResponseWriter,r *http.Request){
	users := models.GetAll(models.USERS)
	utils.ToJson(w,users,http.StatusOK)
}

func GetUser(w http.ResponseWriter,r *http.Request){
	vars := mux.Vars(r)
	id,_ := strconv.ParseUint(vars["id"],10,64)
	user := models.GetById(models.USERS,id)
	utils.ToJson(w,user,http.StatusOK)
}
func DeleteUser(w http.ResponseWriter,r *http.Request){
	vars := mux.Vars(r)
	id,_ := strconv.ParseUint(vars["id"],10,64)
	rows,err := models.Delete(models.USERS,id)
	if err != nil{
		utils.ToJson(w,err.Error(),http.StatusUnprocessableEntity)
		return
	}
	utils.ToJson(w,rows,http.StatusOK)
}

