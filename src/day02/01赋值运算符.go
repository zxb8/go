package main

import "fmt"

func main01() {
	a := 10
	b := 20
	c := a + b
	c += 20
	fmt.Println("c", c)
}

func main() {
	c := 20
	c %= 3

	fmt.Println(c)
}
