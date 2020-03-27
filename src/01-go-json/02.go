package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"encoding/xml"
	"log"
	"io"
	"os"
)

type User struct{
	Username string
	Password string
	Email string
}

type UserDb struct{
	Users []User
	Type string
}

func main01(){
	users :=[]User{
		{Username:"Jane Doe",Password:"jack",Email:"hy@163.com"},
		{Username:"John Doe",Password:"change me",Email:"hy@163.com"},
	}
	db := UserDb{Users:users,Type:"Simple"}

	var buf = new(bytes.Buffer)
	enc :=json.NewEncoder(buf)
	enc.Encode(db)
	io.Copy(os.Stdout,buf)
}

func main02(){
	users :=[]User{
		{Username:"Jane Doe",Password:"jack",Email:"hy@163.com"},
		{Username:"John Doe",Password:"change me",Email:"hy@163.com"},
	}
	db := UserDb{Users:users,Type:"Simple"}

	var buf = new(bytes.Buffer)
	enc :=json.NewEncoder(buf)
	enc.Encode(db)

	f,err :=os.Create("./user.db.json")
	if err != nil{
		log.Fatal(err)
	}

	defer f.Close()

	io.Copy(f,buf)

	db1,err := readJson("./user.db.json")

	if err != nil{
		log.Println(err)
	}
	f,err = os.Create("user.xml")
	if err != nil{
		fmt.Println(err)
	}

	xmlD :=xml.NewEncoder(f)
	xmlD.Encode(db1)
}

/**
生成xml
 */
func main(){

	//读取json数据
	db1,err := readJson("./user.db.json")

	if err != nil{
		log.Println(err)
	}
	f,err := os.Create("./user.db.xml")
	if err != nil{
		fmt.Println(err)
	}

	xmlD :=xml.NewEncoder(f)
	xmlD.Encode(db1)
}


func readJson(path string) (db *UserDb,err error){
	f,err :=os.Open(path)
	if err != nil{
		return nil,err
	}
	defer f.Close()
	dec :=json.NewDecoder(f)
	db = new(UserDb)
	dec.Decode(db)
	return
}