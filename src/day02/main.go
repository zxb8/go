package main

import (
	"log"
	"net"
	"net/http"
	"net/rpc"
)

type Item struct{
	title string
	body string
}

type API int

var database []Item

func (a *API) GetByName1(title string,reply *Item) error{
	var getItem Item
	for _,val := range database{
		if val.title == title{
			getItem = val
		}
	}
	*reply = getItem
	return nil
}

func (a *API) AddItem1(item Item,reply *Item) error{
	database = append(database,item)
	*reply = item
	return nil
}

func (a *API) EditItem1(edit Item,reply *Item) error{
	var changed Item
	for idx,val := range database{
		if val.title == edit.title{
			database[idx] = Item{edit.title,edit.body}
			changed = database[idx]
		}
	}
	*reply = changed
	return nil
}

func (a *API) DeleteItem1(item Item,reply *Item) error{
	var del Item
	for idx,val := range database{
		if val.title == item.title && val.body == item.body{
			database = append(database[:idx],database[idx+1:]...)
			del = item
			break
		}
	}
	* reply = del
	return nil
}


/*
func GetByName(title string) Item{
	var getItem Item
	for _,val := range database{
		if val.title == title{
			getItem = val
		}
	}
	return getItem
}

func AddItem(item Item) Item{
	database = append(database,item)
	return item
}

func EditItem(title string ,edit Item) Item{
	var changed Item
	for idx,val := range database{
		if val.title == edit.title{
			database[idx] = edit
			changed = edit
		}
	}
	return changed
}

func DeleteItem(item Item) Item{
	var del Item
	for idx,val := range database{
		if val.title == item.title && val.body == item.body{
			database = append(database[:idx],database[idx+1:]...)
			del = item
			break
		}
	}
	return del
}*/

func main(){
 /*	fmt.Println("initial database:",database)
 	a := Item{"first","a test item"}
 	b := Item{"second","a second item"}
 	c := Item{"third","a third item"}

 	AddItem(a)
 	AddItem(b)
 	AddItem(c)

 	fmt.Println("second database:",database)

 	DeleteItem(b)

 	fmt.Println("third database:",database)

 	EditItem("third",Item{"four","a new item"})

 	fmt.Println("fourth database:",database)

 	x := GetByName("fourth")
 	y := GetByName("first")

 	fmt.Println(x,y)*/

 	var api = new (API)
 	err := rpc.Register(api)

 	if err != nil{
 		log.Fatal("error registering API",err)
	}
 	rpc.HandleHTTP()

 	listener,err :=net.Listen("tcp",":4040")

 	if err != nil{
 		log.Fatal("Listener error",err)
	}

 	log.Printf("serving rpx on port %d",4040)
 	err = http.Serve(listener,nil)

 	if err != nil{
 		log.Fatal("error serving:",err)
	}



}