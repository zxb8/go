package main

import (
	"io"
	"log"
	"net/http"
)

/**
 net/http 包涵盖了与 HTTP 请求发送和处理的相关代码。虽然包中定义了大量类型、函数，但最重要、最基础的概念只有两个：ServeMux 和 Handler。
ServeMux 是 HTTP 请求多路复用器（即路由器，HTTP request router），记录着请求路由表。对于每一个到来的请求
，它都会比较请求路径和路由表中定义的 URL 路径，并调用指定的 handler 来处理请求。

Handler 的任务是返回一个 response 消息，并把 header 和 body 写入消息中。
任何对象只要实现了 http.Handler 接口的接口方法 ServeHTTP 就可以成为 handler。ServeHTTP 的方法签名是：ServeHTTP(ResponseWriter, *Request)。
**/


type myHandler struct{}

func(h *myHandler) ServeHTTP(w http.ResponseWriter,r *http.Request){
	io.WriteString(w,"Hello world!\n")
}

func main(){
	mux := http.NewServeMux()
	mux.Handle("/hello",new(myHandler))

	log.Println("Listening...")
	http.ListenAndServe(":3000",mux)
}
