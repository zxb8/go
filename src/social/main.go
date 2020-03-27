package main

import (
	"fmt"
	"github.com/hako/branca"
	"log"
	"net/http"
	"social/database"
	"social/handler"
	"social/service"
)

var (
	port = 8000
)

func main(){
	codec := branca.NewBranca("supersecretkeyyoushouldnotcommit") // This key must be exactly 32 bytes long.
	codec.SetTTL(uint32(service.TokenLifespan.Seconds()))
	db := database.Connect()
	s :=  service.New(db,codec)
	h := handler.New(s)
	addr := fmt.Sprintf(":%d",port)
	log.Printf("accepting connections on port %d\n",port)

	if err := http.ListenAndServe(addr,h);err != nil{
		log.Fatalf("could not start server :%v\n",err)
	}
}
