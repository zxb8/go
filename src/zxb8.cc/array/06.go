package main

import "fmt"

func main(){
	var names[] string
	if names == nil{
		names = append(names,"abc","def","ghi")
		fmt.Println(names)
	}

	a :=[]string{"abc","def","ghi"}
	b :=[]string{"aaa","bbb","ccc"}

	c := append(a,b...)
	fmt.Println(c)


	e :=[]int{1,2,3}
	fun(e)
	fmt.Println(e)


}


func fun(n []int){
	for i,_:=range n{
		n[i]  = -2
	}
}