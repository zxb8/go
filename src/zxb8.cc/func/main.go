package main

import (
	"fmt"
	"math/rand"
	"time"
)

func add(a,b int) int{
	return a + b
}

func sayHello(){
	fmt.Printf("hello world\n")
}

func calc(a,b int)(int,int){
	sum := a +b
	sub := a-b
	return sum,sub
}

func calc_v1(b ...int) int{
	sum := 0
	for _,val:= range b{
		sum  +=val
	}
	return sum
}

func testFunc3(){
	var i int = 0
	defer fmt.Printf("defer  i=%d ",i)
}

func sumArray(a [10]int) int{
	var sum int = 0
	for i:=0;i<len(a);i++{
		sum = sum + a[i]
	}
	return sum
}

func sumArray1(a [10] int) int{
	var sum int = 0
	for _,val :=range a{
		sum = sum + val
	}
	return sum
}

func testArraySum(){
	var b[10] int
	for i:=0;i <len(b);i++{
		b[i] = i
	}
	sum := sumArray(b)
	fmt.Printf("sum=%d\n",sum)
}

func testArraySum1(){
	rand.Seed(time.Now().Unix())
	var b[10]int
	for i:=0;i<len(b);i++{
		b[i] = rand.Intn(1000)
	}
	sum := sumArray(b)
	fmt.Printf("sum=%d\n",sum)
}


func main(){
	b :=add(1,22)
	fmt.Printf("b=%d\n",b)
	sayHello()

	sum,sub :=calc(10,20)
	fmt.Printf("sum=%d,sub=%d\n",sum,sub)

	s :=calc_v1(1,2,3,4,5,6)
	fmt.Printf("s=%d\n",s)

	testArraySum()
	testArraySum1()
}
