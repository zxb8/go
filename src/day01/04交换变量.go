package main

import "fmt"

func main01() {
	a := 10
	b := 20
	var c int
	c = a
	a = b
	b = c

	fmt.Println(a, b)
}

func main02() {
	a := 10
	b := 20
	a, b = b, a

	fmt.Println(a, b)

}

func main() {
	a := 10
	b := 20

	a = a + b
	b = a - b
	a = a - b

	fmt.Println(a, b)
}
