package main

import (
	"log"
	"net/http"
)

func main(){
	mux := http.NewServeMux()
	mux.handle("/get",handlers.GetKey(db))
	mux.handle("/set",handlers.PutKey(db))

	log.Printf("server on port 8080")

	err :=http.ListenAndServe(":8080",mux)

	log.Fatal(err)
}
