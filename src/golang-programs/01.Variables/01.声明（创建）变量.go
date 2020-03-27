package main

import "fmt"

/**
https://www.golangprograms.com/go-language/variables.html


A variable is a name that represents a data stored in the computer's memory. Here're some important things to know about Golang variables:

Golang is statically typed language, which means that when variables are declared, they either explicitly or implicitly assigned a type even before your program runs.
Golang requires that every variable you declare inside main() get used somewhere in your program.
You can assign new values to existing variables, but they need to be values of the same type.
A variable declared within brace brackets {}, the opening curly brace { introduces a new scope that ends with a closing brace }.
 */


 /**
 Declaring (Creating) Variables
The keyword var is used for declaring variables followed by the desired name and the type of value the variable will hold.

You can declare a variable without assigning the value, and assign the value later.
  */


  func main(){
  	var i int
  	var s string

  	i = 10
  	s = "Canda"

  	fmt.Println(i)
  	fmt.Println(s)
  }