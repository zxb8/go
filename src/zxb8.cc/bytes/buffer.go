package main

import (
	"bytes"
	"fmt"
)

func main(){
	var chars [6]string
	chars[0] = "a"
	chars[1] = "b"
	chars[2] = "c"
	chars[3] = "d"
	chars[4] = "o"
	chars[5] = "p"

	fmt.Println("A Palindrome")
	fmt.Println("****************")

	var buffer bytes.Buffer

	for i:=0;i<len(chars);i++{
		buffer.WriteString(chars[i])
	}

	for i:=len(chars) -2;i>=0;i--{
		buffer.WriteString(chars[i])
	}

	fmt.Println(buffer.String())
}
