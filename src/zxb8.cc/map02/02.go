package main

import "fmt"

func main01(){
	x := make(map[string]int)
	x["abc"] = 100
	x["def"] = 200
	fmt.Println("map x contents:",x)

	fmt.Println("map x content:",x["abc"],x["def"])
}

func main(){
	x := map[string]int{
		"abc":100,
		"def":200,
	}
	fmt.Println("x:=",x)
}