package main

import "fmt"

type Dog struct{
	Breed string
	Weight int
}

func main(){
	poodle := Dog{"poodle",34}
	fmt.Println(poodle)

	fmt.Printf("%+v\n",poodle)


}
