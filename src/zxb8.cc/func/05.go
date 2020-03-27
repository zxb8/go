package main

import "fmt"

/**
output:

[]
[abc def ghi]
[abc def ghi klm]

 */
func main(){
	ve()
	ve("abc","def","ghi")
	ve("abc","def","ghi","klm")
}

/**
 可变参数
 */
func ve(s ...string){
	//fmt.Println(s[0])
	//fmt.Println(s[1])
	//fmt.Println(s[2])
	fmt.Println(s)
}
