package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main(){
	rand.Seed(time.Now().UnixNano())
	var a map[string]int  = make(map[string]int,1024)

	for i:=0;i<128;i++{
		key := fmt.Sprintf("stu%d",i)
		value := rand.Intn(1000)
		a[key] = value
	}

	//fmt.Println(a)

	for key,value := range a{
		fmt.Printf("key=%s,value=%d",key,value)
	}
}
