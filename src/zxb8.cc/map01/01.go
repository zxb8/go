package main

import "fmt"

func main(){
	var a map[string]int
	fmt.Printf("a:%v\n",a)

	if a == nil{
		a = make(map[string]int)
		a["stu01"] = 1000
		a["stu02"] = 1000
		a["stu03"] = 1000
		fmt.Printf("a=%v\n",a)
		fmt.Printf("a=%#v\n",a)
	}
}
