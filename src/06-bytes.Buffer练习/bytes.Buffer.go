package main

import (
	"bytes"
	"os"
)
//https://blog.csdn.net/txj236/article/details/52130439
func main(){
	var name1 bytes.Buffer
	name1.Write([]byte("aaaa"))
	name1.WriteTo(os.Stdout)
}