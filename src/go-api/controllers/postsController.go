package controllers

import (
	"encoding/json"
	"go-api/models"
	"go-api/utils"
	"net/http"
)

func PostPost(w http.ResponseWriter,r *http.Request){
	body := utils.BodyParser(r)
	var post models.Post
	err :=json.Unmarshal(body,&post)
	if err != nil{
		utils.ToJson(w,err.Error(),http.StatusUnprocessableEntity)
		return
	}
	err = models.NewPost(post)
	if err != nil{
		utils.ToJson(w,err.Error(),http.StatusUnprocessableEntity)
		return
	}
	utils.ToJson(w,"success",http.StatusCreated)
}

func GetPosts(w http.ResponseWriter,r *http.Request){
	posts := models.GetAll(models.POSTS)
	utils.ToJson(w,posts,http.StatusOK)
}
