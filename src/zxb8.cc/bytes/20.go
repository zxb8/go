package main

import (
	"fmt"
	"io"
	"os"
)

func main(){
	content := "Hello from Go!"

	file,err :=os.Create("./fromStriing.txt")

	checkError(err)

	defer file.Close()

	ln,err := io.WriteString(file,content)

	checkError(err)

	fmt.Printf("All done with file of %v characters",ln)


}

func checkError(err error){
	if err != nil{
		panic(err)
	}
}
