package main

import "fmt"

func testInterface(){
	var a interface{}
	var b int  = 100
	var c float32= 1.2
	var d string  = "hello"

	a = b
	fmt.Printf("a=%v",a)
	fmt.Printf("b=%#v",b)
	fmt.Printf("c=%#v",c)
	fmt.Printf("d=%#v",d)
}

func studentStore(){
	var stuMap map[int]map[string]interface{}
	stuMap = make(map[int]map[string]interface{},16)
	var id =1
	var name = "stu01"
	var score = 78.2
	var age  =18
	value,ok := stuMap[id]
	if !ok{
		value = make(map[string]interface{},8)
	}
	value["name"] = name
	value["id"] = id
	value["score"] = score
	value["age"] = age
	stuMap[id] = value

}

func main(){
	testInterface()
	studentStore()
}
