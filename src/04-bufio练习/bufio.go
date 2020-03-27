package main

import (
	"bufio"
	"fmt"
	"os"
)

func main(){
	w :=bufio.NewWriter(os.Stdout)
	w.Write([]byte("hahhasdassssssssssssssssssssssssssssssssssssssssssssssssssssss"))
	w.Flush()
    //bufio 默认缓冲区大小是4096，如果不强制刷新，写入的内容不到缓冲区大小，不会输出，可以强制刷新
	//不要忘了刷新


	w1 := bufio.NewWriter(os.Stdout)
	//w1 := bufio.NewReaderSize(os.Stdout,3) 可以设置缓冲区大小
	fmt.Fprint(w1, "Hello, ")
	fmt.Fprint(w1, "world!\n")
	w1.Flush() // Don't forget to flush!
}
