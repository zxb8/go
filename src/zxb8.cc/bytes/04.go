package main

import (
	"fmt"
	"time"
)

func main(){
	t := time.Date(2009,time.November,10,23,0,0,0,time.UTC)
	fmt.Printf("Go launced at %s\n",t)

	now := time.Now()
	fmt.Printf("The time is %s\n",now)

	fmt.Println("The month is",t.Year())
	fmt.Println("The day is",t.Day())
	fmt.Println("The weekday is",t.Weekday())

	tomorrow := t.AddDate(0,0,1)
	fmt.Printf("Tomorrow is %v ,%v,%v %v\n", tomorrow.Weekday(),tomorrow.Month(),tomorrow.Day(),tomorrow.Year())

}
