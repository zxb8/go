package main

import (
	"fmt"
	"math"
)

func main0201() {
	//变量的定义和使用
	var sun int = 50

	// sun = sun +
	//表达式
	//变量在程序运行过程中，值可以发生改变
	sun = sun + 25

	fmt.Println("sun:", sun)
}

func main0202() {
	//变量的声明,如果没有赋值，默认值为0
	var sun int

	sun = 50 //为变量赋值

	fmt.Println(sun) //output:0
}

func main() {
	//float64 浮点类型
	var value float64 = 2

	var sum float64 = math.Pow(value, 10)

	fmt.Println(sum)

}
