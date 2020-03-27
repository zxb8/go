package main

import "fmt"

type X struct{
	a string
	b int
}

func (varx X) method1(n string){
	varx.a = n
}

func (vary *X) method2(num int){
	vary.b  = num
}

func main(){
	m := X{"abc",121}
	fmt.Printf("m.a=%s,m.b=%d\n",m.a,m.b)
	m.method1("xyz")
	fmt.Printf("m.a=%s\n",m.a)
	m.method2(11)
	fmt.Printf("m.b=%d",m.b)
}
