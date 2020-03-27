package main

import (
	"fmt"
	"net/http"
)

func handlerFunc(w http.ResponseWriter, r *http.Request){
	fmt.Fprint(w,"<h1>Welcome to my wasome site!</h1>")
	fmt.Fprint(w,r.URL.Path)
}

func main(){
	http.HandleFunc("/",handlerFunc)
	http.ListenAndServe(":8000",nil)
}
