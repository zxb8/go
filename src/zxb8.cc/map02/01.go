package main

import "fmt"

//x := make(map[string]int)

func main(){
	var x map[string]int
	if x == nil{
		fmt.Println("x is nil")
		x= make(map[string]int)
	}
}
