package main

import (
	"flag"
	"log"
	"net/http"
)

func main(){
	port := flag.String("p","8000","port")
	dir := flag.String("d",".","dir")

	http.Handle("/",http.FileServer(http.Dir(*dir)))
	log.Printf("serving %s on Http port:%s\n",*dir,*port)
	log.Fatal(http.ListenAndServe(":"+ *port,nil))
}