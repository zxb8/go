package main

import (
	"fmt"
	"strings"
)

func statWordCount(str string)map[string]int{
	words :=strings.Split(str," ")
	var result map[string]int = make(map[string]int,128)
	for _,v :=range words{
		count,ok :=result[v]
		if !ok{
			result[v] = 1
		}else{
			result[v] =count +1
		}
	}
	return result
}

func main(){
	var str = "how do you do ? do you like me?"
	result := statWordCount(str)
	fmt.Printf("result:%#v\n",result)
}
