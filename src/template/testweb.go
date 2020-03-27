package main

import (
	"fmt"
	"html/template"
	"net/http"
)

func test(w http.ResponseWriter,r *http.Request) {
	t, _ :=template.ParseFiles("templates/header.html")
	fmt.Println(t.Name())
	t.Execute(w,map[string] string {"Title": "My title","Body": "Hi this is my body"})
	fmt.Println("aaaa")
}

func main() {
	http.HandleFunc("/",test)
	http.ListenAndServe(":8080",nil)
}