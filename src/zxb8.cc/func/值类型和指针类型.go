package main

import "fmt"

func testPoint(){
	var a int = 100
	var b *int = &a
	fmt.Println("b指向的地址存储的值为:%d\n",*b)
}

func modify(a *int){
	*a = 100
}
func testPoint3(){
	var b int  = 10
	 p := &b
	modify(p)
	 fmt.Printf("b:%d\n",b)
}

func testPoint5(){
	var a *int = new(int)
	*a = 100
	fmt.Printf("*a=%d\n",*a)

	var b *[]int = new([]int)
	fmt.Printf("*b=%v\n",*b)
}


func main(){
	var a int32
	a = 1000
	fmt.Printf("the addr of a:%p,a:%d",&a,a)

	var b * int32
	fmt.Printf("the addr of b:%p,b:%v\n",&b,b)

	b = &a
	fmt.Printf("the addr of b:%p,b:%v",&b,b)
	testPoint()
	testPoint3()
	testPoint5()
}
