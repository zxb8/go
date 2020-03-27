package main

import "fmt"

/**
Golang uses lexical scoping based on code blocks to determine the scope of variables. Inner block can access its outer block defined variables, but outer block cannot access inner block defined variables.
 */


 var s = "China"

 func main(){
 	fmt.Println(s)
 	x:=true

 	if x{
 		y :=1
 		if x!= false{
 			fmt.Println(s)
 			fmt.Println(x)
 			fmt.Println(y)
		}
	}
 	fmt.Println(x)

 }