package main

import (
	"fmt"
	"reflect"
)

/**
The := short variable assignment operator indicates that short variable declaration is being used. There is no need to use the var keyword or declare the variable type.
 */


 func main(){
 	name := "Jone Doe"
 	fmt.Println(reflect.TypeOf(name))
 }