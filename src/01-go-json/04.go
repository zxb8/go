package main

import (
	"bytes"
	"encoding/json"
	"fmt"
)

type Members struct {
	Id       int    `json:"id"`
	Username string `json:"username"`
	Sex      uint   `json:"gender"`
	Email    string `json:"email"`
}


func main() {
	members := []Members{
		{
			Id:1,
			Username:"小明",
			Sex:1,
			Email:"xiaoming@163.com",
		},
		{
			Id:2,
			Username:"小红",
			Sex:1,
			Email:"xiaohong@163.com",
		},
		{
			Id:3,
			Username:"小华",
			Sex:2,
			Email:"xiaohua@163.com",
		},
	}

	b := &bytes.Buffer{}
	encoder := json.NewEncoder(b)
	err := encoder.Encode(members)
	if err != nil{
		panic(err)
	}
	fmt.Println(b.String())
}





