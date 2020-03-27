package main

import "fmt"

func main01() {
	var score int
	fmt.Scan(&score)

	switch score / 10 {
	case 10:
		fmt.Println("A")
	case 9:
		fmt.Println("A")
	case 8:
		fmt.Println("B")
	case 7:
		fmt.Println("C")
	default:
		fmt.Println("E")
	}
}

func main() {
	m := map[int]string{1: "Luffy", 130: "Sanji", 1301: "Zero"}
}
