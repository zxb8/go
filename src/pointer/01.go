package main

import "fmt"

func main(){
	var me string = "zxb8.cc"
	var u *string

	u = &me
	fmt.Println(&me,&u)

	me = "123"
	fmt.Println(&me,&u)
}
