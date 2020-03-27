package main

import (
	"fmt"
)

func main(){
	i1,i2,i3 :=12,45,68
	intSum := i1 + i2 + i3
	fmt.Println("integer sum:",intSum)

	f1,f2,f3 := 23.5,65.1,76.3
	floatSum := f1 + f2 + f3
	fmt.Println("float sum:",floatSum)

	//var b1,b2,b3,bigSum big.Float
	//b1.SetFloat64(23.5)
	//b2.SetFloat64(65.1)
	//b3.SetFloat64(76.3)
}
