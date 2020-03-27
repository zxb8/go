package main

import "fmt"

func main(){
	a :=[]string{"abc","def"}
	fmt.Println(len(a),cap(a))
	a = append(a,"ghi")
	fmt.Println(len(a),cap(a))
}
