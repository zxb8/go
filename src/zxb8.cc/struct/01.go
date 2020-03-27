package main

import "fmt"

type Basket struct{
	fruits,vegetables string
}

func main(){
	x:= Basket{
		fruits:"apple",
		vegetables:"beans",
	}

	y := Basket{"pine","sss"}
	fmt.Println(x)
	fmt.Println(y)
}
