package api

import (
	"blogApp/auto"
	"blogApp/config"
	"blogApp/router"
	"fmt"
	"log"
	"net/http"
)

func Run(){
	config.Load()
	auto.Load() //自动加载数据
	fmt.Printf("\n\tListening [::]:%d\n",config.PORT)
	listen(config.PORT)
}

func listen(port int){
	r := router.New()
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%d",port),r))
}