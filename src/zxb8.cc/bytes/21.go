package main

import (
	"fmt"
	"io/ioutil"
)

func main(){
	fileName := "./fromStriing.txt"
	content,err := ioutil.ReadFile(fileName)
	checkError1(err)


	fmt.Println(string(content))
}

func checkError1(err error){
	if err != nil{
		panic(err)
	}
}
