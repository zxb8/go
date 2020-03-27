package main

import (
	"github.com/kataras/iris/v12"
	"github.com/kataras/iris/v12/mvc"
	"github.com/kataras/iris/v12/sessions"
	"os"
	"time"
	"zxb8/src/config"
	"zxb8/src/modules/profile/controller"
	"zxb8/src/modules/profile/repository"
	"zxb8/src/modules/profile/usecase"
)

func main(){
	app :=iris.New()

	app.Logger().SetLevel("debug")

	//app.RegisterView(iris.HTML("./templates", ".html").Layout("layout.html"))
	//	// TIP: append .Reload(true) to reload the templates on each request.
	app.RegisterView(iris.HTML("./web/views",".html").Layout("layout.html").Reload(true))
	//app.RegisterView(iris.HTML("./web/views",".html").Reload(true))

	app.HandleDir("/public", "./web/public")


	//app.Get("/",func(ctx iris.Context){
	//	ctx.ViewData("Message","Golang Socmed")
	//	ctx.View("index.html")
	//})

	db,err :=config.GetMongoDB()
	if err!= nil{
		os.Exit(2)
	}

	profileRepository := repository.NewProfileRepositoryMongo(db,"profile")
	profileUsecase := usecase.NewProfileUsecase(profileRepository)

	sessionManager := sessions.New(sessions.Config{
		Cookie :"cookiename",
		Expires:time.Minute * 10,
	})

	//5.注册控制路由
	profile := mvc.New(app.Party("/profile"))

	profile.Register(profileUsecase,sessionManager.Start)

	profile.Handle(new(controller.ProfileController))


	app.Run(iris.Addr(":3000"))

}
