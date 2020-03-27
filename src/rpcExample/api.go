package main

import (
	"log"
	"net"
	"net/http"
	"net/rpc"
)

type Item struct{
	Title string
	Body string
}

type API int

var database []Item

func (a *API) GetDB(title string,reply *[]Item) error{
	*reply = database
	return nil
}

func (a *API) GetByName(title string,reply *Item) error{
	var getItem Item

	for _,val := range database{
		if val.Title == title{
			getItem = val
		}
	}
	*reply = getItem
	return nil
}

func (a *API)CreateItem(item Item,reply *Item) error{
	database = append(database,item)
	*reply = item
	return nil
}

func (a *API)EditItem(edit Item,reply *Item) error{
	var changed Item
	for idx,val := range database{
		if val.Title == edit.Title{
			database[idx] = Item{edit.Title,edit.Body}
			changed = edit
		}
	}
	*reply = changed
	return nil
}

func (a *API)DeleteItem(item Item,reply *Item) error{
	var del Item
	for idx,val := range database{
		if val.Title == item.Title && val.Body == item.Body{
			database = append(database[:idx],database[idx+1:]...)
			del = item
			break
		}
	}
	*reply = del
	return nil
}

func main(){

	var api = new(API)
	err :=rpc.Register(api)
	if err != nil{
		log.Fatal("error registering API",err)
	}

	rpc.HandleHTTP()

	listener,err :=net.Listen("tcp",":8000")

	if err != nil{
		log.Fatal("Listener error",err)
	}

	log.Printf("serving rpc on port %d",8000)

	err =http.Serve(listener,nil)
	if err != nil{
		log.Fatal("error serving:",err)
	}



	//fmt.Println("initial database:",database)
	//
	//a := Item{"first","a test item"}
	//b := Item{"second","a second item"}
	//c := Item{"third","a third item"}
	//
	//
	//CreateItem(a)
	//CreateItem(b)
	//CreateItem(c)
	//
	//fmt.Println("second database:",database)
	//
	//
	//DeleteItem(b)
	//
	//fmt.Println("third database:",database)
	//
	//EditItem("third",Item{"forth","a new Item"})
	//
	//x :=GetByName("fourth")
	//y := GetByName("first")
	//
	//fmt.Println(x,y)



}