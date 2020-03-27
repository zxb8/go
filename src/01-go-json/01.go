package main

import (
	"bytes"
	"encoding/json"
	"io"
	"os"
)

type Person struct{
	Username string
	Password string
	Email string
}

func main(){
	user := Person{
		Username:"jack",
		Password:"123",
		Email:"hyrxb@163.com",
	}
	//fmt.Println(user)

	var buf = new(bytes.Buffer)
	enc := json.NewEncoder(buf)
	enc.Encode(user)
	io.Copy(os.Stdout,buf)
}
