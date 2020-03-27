package main

import "fmt"

//匿名结构体
func main(){
	x := struct{
		fruits,vegetables string
	}{
		fruits:"apple",
		vegetables:"beans",
	}

	fmt.Println("X:",x)
}
