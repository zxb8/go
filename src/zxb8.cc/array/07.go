package main

import "fmt"

func main(){
	a :=[][]string{
		{"abc","def"},
		{"xxx","yyy"},
	}

	for _,j := range a{
		for _,k :=range j{
			fmt.Printf("%s\t",k)
		}
		fmt.Printf("\n")
	}
}
