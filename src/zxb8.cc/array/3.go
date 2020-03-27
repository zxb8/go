package main

import "fmt"

func main(){
	a :=[3][2]string{
		{"ab","ddd"},
		{"ss","ddd"},
		{"sss","eee"},
	}

	fmt.Println(a)

	var b[2][2] string
	b[0][1] = "aa"
	b[0][0] = "000"
	b[1][0] = "222"
	b[1][1] = "111"
	fmt.Println(b)

	for _,x := range a{
		for _,y := range x{
			fmt.Printf("%s\t",y)
		}
		fmt.Printf("\n")
	}
}
