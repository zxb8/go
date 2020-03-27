package main

import (
	"io"
	"net/http"
)

func main(){
	http.HandleFunc("/",demo)
	http.Handle("/static/",http.StripPrefix("/static",http.FileServer(http.Dir("./assets"))))
	http.ListenAndServe(":8000",nil)
}

func demo(w http.ResponseWriter,req *http.Request){
	w.Header().Set("Content-Type","text/html;charset=utf-8")
	io.WriteString(w,`<img src="/static/01.png">`)
}
