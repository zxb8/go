package main

import (
	"fmt"
	"log"
	"net/rpc"
)

type Item struct{
	Title string
	Body string
}

func main(){
	var reply Item
	var db []Item

	client,err :=rpc.DialHTTP("tcp","localhost:8000")
	if err != nil{
		log.Fatal("connection error:",err)
	}

	a := Item{"First","A first item"}
	b := Item{"Second","A second item"}
	c := Item{"Third","A third item"}

	client.Call("API.CreateItem",a,&reply)
	client.Call("API.CreateItem",b,&reply)
	client.Call("API.CreateItem",c,&reply)


	client.Call("API.GetDB","",&db)

	fmt.Println(db)

	client.Call("API.EditItem",Item{"Second","A new body"},&reply)

	client.Call("API.DeleteItem",c,&reply)

	client.Call("API.DeleteItem",c,&reply)
	client.Call("API.GetDB","",&db)
	fmt.Println("Database:",db)


}
