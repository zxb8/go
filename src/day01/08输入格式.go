
package main

import "fmt"

func main01() {
	var a int

	//通过键盘为变量赋值
	fmt.Scan(&a)

	fmt.Printf("a = %d", a)
	fmt.Println(&a)

}

func main02() {
	var a, b int

	//空格和回车表示一个输入接收结束
	fmt.Scan(&a, &b)

	fmt.Println(a)
	fmt.Println(b)

}

func main03() {
	var a int
	var b string

	//在接收字符串时要使用空格作为分隔
	fmt.Scanf("%d%s", &a, &b)

	fmt.Println(a)
	fmt.Println(b)
	fmt.Printf("===%s===", b) //output:======
}

//通过键盘输入学生三门成绩，计算总成绩和平均成绩

func main() {
	var c, m, e int
	fmt.Scan(&c, &m, &e)

	sum := c + m + e

	fmt.Println("总成绩,", sum)
	fmt.Println("平均成绩,", sum/3)
}
