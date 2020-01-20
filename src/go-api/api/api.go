package api

import (
	"fmt"
	"go-api/routes"
	"log"
	"net/http"
)

func Run(){
	//models.AutoMigrations()
	//models.NewUser(models.User{Nickname:"jack",Email:"jack@163.com",Password:"123"})
	listen(9000)
}

func listen(p int){
	fmt.Printf("Listening port %d...",p)
	port := fmt.Sprintf(":%d",p)
	r := routes.NewRouter()
	log.Fatal(http.ListenAndServe(port,r))
}