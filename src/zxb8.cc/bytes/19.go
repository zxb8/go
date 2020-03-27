package main

import "fmt"

func main(){
	defer fmt.Println("Close the file!")
	fmt.Println("Open the file!")

	defer fmt.Println("statement 1")
	defer fmt.Println("statement 2")
	defer fmt.Println("statement 3")
	defer fmt.Println("statement 4")
}
