package main

import (
	"zxb8.cc/gin-example/router"
)

func main(){
	router := router.NewRouter()

	router.Run(":8000")
}
