package main

import "fmt"

func main(){
	var x,y,z int
	x = 75
	y = 25
	z = x & y
	fmt.Println(z)
	z= x|y
	fmt.Println(z)
}
