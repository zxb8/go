package main

import "fmt"

func main() {
	a := 10
	b := 20

	c := a < b && false
	fmt.Println(c)
}
