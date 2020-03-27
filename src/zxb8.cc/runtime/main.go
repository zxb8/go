package main

import (
	"fmt"
	"time"
	"runtime"
)

var i int

func calc(){
	for{
		i++
	}
}

func main(){
	cpu := runtime.NumCPU()
	fmt.Printf("cpu num=%d",cpu)

	runtime.GOMAXPROCS(2)

	for i:=0;i<10;i++{
		go calc()
	}

	time.Sleep(time.Second * 1000)

}
