package main

import "fmt"

func main(){
	fmt.Println(add(20,30))
}

func add(x,y int) int{
	total := x + y
	return total
}
