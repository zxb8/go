package controllers

import (
	"encoding/json"
	"go-api/models"
	"go-api/utils"
	"net/http"
)

func PostFeedbacks(w http.ResponseWriter,r *http.Request){
	body := utils.BodyParser(r)
	var feedback models.Feedback
	err :=json.Unmarshal(body,&feedback)
	if err != nil{
		utils.ToJson(w,err.Error(),http.StatusUnprocessableEntity)
		return
	}
	err = models.NewFeedback(feedback)
	if err != nil{
		utils.ToJson(w,err.Error(),http.StatusUnprocessableEntity)
		return
	}
	utils.ToJson(w,"success",http.StatusCreated)
}

func GetFeedbacks(w http.ResponseWriter,r *http.Request){
	feedbacks := models.GetAll(models.FEEDBACKS)
	utils.ToJson(w,feedbacks,http.StatusOK)
}
