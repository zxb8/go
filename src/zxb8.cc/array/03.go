package main

import "fmt"

func main(){
	a :=[5]int{1,2,3,4,5}
	b:= a[1:4]
	c:=a[1:4]

	fmt.Println(b)
	fmt.Println(c)

	for index,_:=range b{
		b[index]++
	}
	fmt.Println(b)
	fmt.Println(c)

}
