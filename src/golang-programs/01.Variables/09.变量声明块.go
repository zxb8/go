package main

import "fmt"

/**
Variable Declaration Block
Variables declaration can be grouped together into blocks for greater readability and code quality.
变量声明可以组合成块以提高可读性和代码质量。
 */


 var (
 	product = "Mobile"
 	quantity = 50
 	price  = 50.50
 	inStock = true
 )

func main(){
	fmt.Println(quantity)
	fmt.Println(price)
	fmt.Println(product)
	fmt.Println(inStock)
}