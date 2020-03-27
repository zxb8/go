package main

import (
	"bytes"
	"fmt"
)

/**
bytes.buffer是一个缓冲byte类型的缓冲器存放着都是byte
Buffer 是 bytes 包中的一个 type Buffer struct{…}

A buffer is a variable-sized buffer of bytes with Read and Write methods. The zero value for Buffer is an empty buffer ready to use.
（是一个变长的 buffer，具有 Read 和Write 方法。 Buffer 的 零值 是一个 空的 buffer，但是可以使用）
Buffer 就像一个集装箱容器，可以存东西，取东西（存取数据）
 */

 func main(){
 	//方式1
 	var b bytes.Buffer
 	b.Write([]byte("Hello World"))
 	fmt.Println(b.String())

	 //方式2
 	b1 := new(bytes.Buffer) //直接使用new初始化,可以直接使用
 	b1.Write([]byte("zxb8.cc"))

 	fmt.Println(b1.String())

	 // 其它两种定义方式
	 //func NewBuffer(buf []byte) *Buffer
	 //func NewBufferString(s string) *Buffer

	 buf1:=bytes.NewBufferString("swift")
	 buf2:=bytes.NewBuffer([]byte("swift"))
	 buf3:=bytes.NewBuffer([]byte{'s','w','i','f','t'})
	 fmt.Println("===========以下buf1,buf2,buf3等效=========")
	 fmt.Println("buf1:", buf1)
	 fmt.Println("buf2:", buf2)
	 fmt.Println("buf3:", buf3)

	 fmt.Println("===========以下创建空的缓冲器等效=========")
	 buf4:=bytes.NewBufferString("")
	 buf5:=bytes.NewBuffer([]byte{})
	 fmt.Println("buf4:", buf4)
	 fmt.Println("buf5:", buf5)


	 //向 Buffer 中写入数据

	 fmt.Println("===========以下通过Write把swift写入Learning缓冲器尾部=========")
	 newBytes := []byte("swift")
	 //创建一个内容Learning的缓冲器
	 buf := bytes.NewBuffer([]byte("Learning"))
	 //打印为Learning
	 fmt.Println(buf.String())
	 //将newBytes这个slice写到buf的尾部
	 buf.Write(newBytes)
	 fmt.Println(buf.String())

	 //WriteString
	 //使用WriteString方法，将一个字符串放到缓冲器的尾部

	 fmt.Println("===========以下通过WriteString把swift写入Learning缓冲器尾部=========")
	 newString := "swift"
	 //创建一个string内容Learning的缓冲器
	 s1 :=bytes.NewBufferString("Learning")
	 //打印为Learning
	 fmt.Println(s1.String())
	 //将newString这个string写到buf的尾部
	 s1.WriteString(newString)
	 fmt.Println(s1.String())

	 //WriteByte
	 //将一个byte类型的数据放到缓冲器的尾部

	 fmt.Println("===========以下通过WriteByte把!写入Learning缓冲器尾部=========")
	 var newByte byte = '!'
	 //创建一个string内容Learning的缓冲器
	 s2 := bytes.NewBufferString("Learning")
	 //打印为Learning
	 fmt.Println(s2.String())
	 //将newString这个string写到buf的尾部
	 s2.WriteByte(newByte)
	 fmt.Println(s2.String())
 }

