package main

import "fmt"

func main(){
	var score int
	fmt.Scan(&score)

	if score > 700{
		fmt.Println("我要上清华")
	}else{
		fmt.Println("我要上北大")
	}

}