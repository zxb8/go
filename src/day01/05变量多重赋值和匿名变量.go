package main

import "fmt"

func main01() {
	a, b, c, d := 10, 3.14, "传智播客", "黑马程序员"
	fmt.Println(a)
	fmt.Println(b)
	fmt.Println(c)
	fmt.Println(d)
}

func main02() {
	var a int = 10
	var b int = 20

	//在一个作用域范围内变量名不能重复
	// var a float64 = 3.14 //a redeclared in this block

	fmt.Println(a)
	fmt.Println(b)

	//a, b := 20, 30 //no new variables on left side of :=

	//如果在多重赋值时有新定义的变量，可以使用自动推到类型
	a, b, c, d := 20, 30, "你好", "hello world"

	fmt.Println(a)
	fmt.Println(b)
	fmt.Println(c)
	fmt.Println(d)

}

func main() {

	// _ 表示匿名变量，不接收数据
	_, c, d := 10, "你好", "朋友"

	// fmt.Println(_)
	fmt.Println(c)
	fmt.Println(d)

}
