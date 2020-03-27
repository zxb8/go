package main

import "fmt"

func main(){
	var a map[string]int = map[string]int{
		"stu01":1000,
		"stu02":1000,
		"stu03":1000,
	}
	fmt.Println(a)

	a["stu01"] = 8888
	a["stu02"] = 38333
	fmt.Println(a)

	var key string = "stu02"

	fmt.Printf("the value of key[%s] is :%d\n",key,a[key])
}
