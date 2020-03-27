package main

import "fmt"

func main01(){
	a := [...]string{"abc","def","ghi"}
	b :=a
	b[0] = "xxx"
	fmt.Println(b)
	fmt.Println(a)
}


func myFunc(num [4]int){
	num[0] = 12
	fmt.Println(num)
}

func main(){
	num := [...]int{11,22,33,44}
	fmt.Println(num)
	myFunc(num)
	fmt.Println(num)
}