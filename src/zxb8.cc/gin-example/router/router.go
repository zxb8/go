package router

import (
	"github.com/gin-gonic/gin"
	"zxb8.cc/gin-example/controller"
)

func NewRouter() *gin.Engine{
	router := gin.Default()

	router.GET("/hello",controller.Index)


	return router
}
