package main

import (
	"fmt"
	"reflect"
)

/**
Variable Declaration Omit Types
You can omit the variable type from the declaration, when you are assigning a value to a variable at the time of declaration. The type of the value assigned to the variable will be used as the type of that variable.

变量声明省略类型

在声明时为变量赋值时，可以从声明中省略变量类型。分配给变量的值的类型将用作该变量的类型。

 */


 func main(){
 	var i =10
 	var s = "Canda"

 	fmt.Println(reflect.TypeOf(i))
 	fmt.Println(reflect.TypeOf(s))

 	//int
	 //string
 }