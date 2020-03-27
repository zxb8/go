package main

import (
	"fmt"
	"time"
)

func testTime(){
	now := time.Now()
	fmt.Printf("current time:%v\n",now)

	year := now.Year()
	month := now.Month()
	day  := now.Day()
	hour := now.Hour()
	minute := now.Minute()
	send := now.Second()

	fmt.Printf("%02d-%02d-%02d-%02d-%02d-%02d",year,month,day,hour,minute,send)

}

func testFor1(){
	var i int
	for i=1;i<=10;i++{
		if i % 2== 0{
			continue
		}
		fmt.Printf("i=%d\n",i)
	}
	fmt.Printf("i=%d",i)
}


func testSwitch(){
	switch a:=2;a{
	case 1:
		fmt.Printf("a=1")
	case 2:
		fmt.Printf("a=2")
	case 3:
		fmt.Printf("a=3")
	case 4:
		fmt.Printf("a=4")
	case 5:
		fmt.Printf("a=5")
	}
}


func main(){
	//testTime()
	//testFor1()
	testSwitch()


}
