package main

import "fmt"

type Dog struct{
	Bread string
	Weight int
	Sound string
}


func (d Dog) Speak() {
	fmt.Println(d.Sound)
}

func (d Dog) SpeakThreeTimes(){
	d.Sound = fmt.Sprintf("%v! %v! %v!",d.Sound,d.Sound,d.Sound)
	fmt.Print(d.Sound)
}


func main(){
	poodle := Dog{"poodle",37,"Woof"}
	fmt.Println(poodle)
	poodle.Speak()
	poodle.Sound = "Arf"
	poodle.Speak()
}
