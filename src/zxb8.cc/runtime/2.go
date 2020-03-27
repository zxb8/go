package main

import "fmt"

const (
	x= 10
	y = 20
	z  =30
)

func main(){
	const name string = "abc"
	fmt.Println(name)

	const a = 35
	fmt.Println(a)

	fmt.Println(x,y,z)
}
