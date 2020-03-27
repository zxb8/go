package main

import "github.com/kataras/iris/v12"

func main(){
	app :=iris.New()
	htmlEngine := iris.HTML("./",".html")
	app.RegisterView(htmlEngine)

	app.Get("/",func(ctx iris.Context){
		ctx.WriteString("Hello world !")
	})

	app.Get("/hello",func(ctx iris.Context){
		ctx.ViewData("Title","测试页面")
		ctx.ViewData("Content","Hello world!")
	})

	app.Run(iris.Addr(":8080"),iris.WithoutServerError(iris.ErrServerClosed))
}
