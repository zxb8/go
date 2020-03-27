package main

import "fmt"

func main(){
	var colors = []string{"red","Green","Blue"}
	fmt.Println(colors)

	colors = append(colors,"Purlen")
	fmt.Println(colors)

	colors = append(colors[1:len(colors)])
	fmt.Println(colors)

	colors = append(colors[:len(colors)-1])
	fmt.Println(colors)

	numbers := make([]int,5,5)
	numbers[0] = 1
	numbers[1] = 2
	numbers[2] = 3
	numbers[3] = 4
	fmt.Println(numbers)

	numbers = append(numbers,25)
	fmt.Println(numbers)


}
