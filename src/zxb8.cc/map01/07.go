package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main(){
	rand.Seed(time.Now().UnixNano())
	var s []map[string]int
	s = make([]map[string]int,5,16)
	for index,val := range s{
		fmt.Printf("slice[%d]=%v\n",index,val)
	}
	s[0] = make(map[string]int,16)
	s[0]["stu01"] = 1000
	s[0]["stu02"] = 1000
	s[0]["stu03"] = 1000

	for index,value := range s{
		fmt.Printf("slice[%d]=%v\n",index,value)
	}
}