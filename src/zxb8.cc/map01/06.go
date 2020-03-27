package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main(){
	rand.Seed(time.Now().UnixNano())
	var a map[string]int = make(map[string]int,1024)

	for i:=0;i<128;i++{
		key := fmt.Sprintf("key%d",i)
		a[key] = i
	}

	fmt.Println(a)
}
