package main

import "fmt"

type X struct{
	a string
	b int
}

func (varx X) method1(){
	fmt.Printf("a:%s b:%d",varx.a,varx.b)
}

func main(){
	vary := X{
		a:"abc",
		b:10,
	}
	vary.method1()
}
