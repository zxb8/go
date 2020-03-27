package main

import (
	"fmt"
	"time"
)

func main(){
	timer := time.NewTicker(3*time.Second)
	<-timer.C
	fmt.Println("test")

	//fmt.Fprintf()
	//fmt.Println()
}
