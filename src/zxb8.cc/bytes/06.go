package main

import "fmt"

func main(){
	var colors [3]string
	colors[0] = "Red"
	colors[1] = "Blue"
	colors[2] = "Green"

	fmt.Println(colors)
	fmt.Println(colors[1])

	var number = [5]int{
		5,3,2,1,2,
	}
	fmt.Println(number)
	fmt.Println("Number of nubers:",len(number))
}
