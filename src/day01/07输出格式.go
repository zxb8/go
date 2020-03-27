package main

import "fmt"

func main() {
	// fmt.Println("hello world")
	// fmt.Print("性感荷官在线发牌")
	// fmt.Print("澳门在线赌场上线了")

	//format

	a := 10
	b := 3.14

	//%d 是一个占位符，表示输出一个整形数据
	//%f 是一个占位符，表示一个浮点型数据
	//%.2f 保留小数位数为2位，会第三位小数进行四舍五入
	fmt.Printf("a=%d, b= %.2f\n", a, b)

	c := "你愁啥"

	//%s 是一个占位符， 表示输出一个字符串类型
	fmt.Printf("%s", c)

}
