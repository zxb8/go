package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

func main(){
	resp,err :=http.Get("http://www.baidu.com")

	if err != nil{
		fmt.Println("err")
	}
	defer resp.Body.Close()

	body,err:=ioutil.ReadAll(resp.Body)

	fmt.Println(string(body))



}


