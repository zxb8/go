package main

import (
	"log"
	"net/http"
	"os"
	"text/template"
)

type Inventory struct{
	Material string
	Count uint
}


func main(){
	sweaters :=Inventory{"wool",17}
	tmpl,err := template.New("test").Parse("{{.Count}} items are made of {{.Material}}")
	if err != nil{
		panic(err)
	}

	err = tmpl.Execute(os.Stdout,sweaters) //输入到控制台
	if err != nil{
		panic(err)
	}

	http.Handle("/",http.FileServer(http.Dir("client")))
	log.Fatal(http.ListenAndServe(":8080",nil))

}
