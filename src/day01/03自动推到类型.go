package main

import "fmt"

func main() {
	//var a int =10
	// var b int  =20
	a := 10     //int 整形
	b := 20.80  //float64 浮点型
	c := "自学吧学习社区" //string 字符串类型

	fmt.Println(a)
	fmt.Println(b)
	fmt.Println(c)

	//不同的数据类型不能计算
	//fmt.Println(a + b) //invalid operation: a + b (mismatched types int and float64)

}
