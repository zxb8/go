package main

import (
	"log"
	"net/http"
	"text/template"
)

type Inventory struct{
	Material string
	Count uint
}

func booksHandler(w http.ResponseWriter,r *http.Request){
	sweaters :=Inventory{"wool",17}
	tmpl,err := template.New("test").Parse("{{.Count}} items are made of {{.Material}}")
	if err != nil{
		panic(err)
	}
	err = tmpl.Execute(w,sweaters) //输入到浏览器
	if err != nil{
		panic(err)
	}
}


func main(){
	http.HandleFunc("/books",booksHandler)
	http.Handle("/",http.FileServer(http.Dir("client")))
	log.Fatal(http.ListenAndServe(":8080",nil))

}
