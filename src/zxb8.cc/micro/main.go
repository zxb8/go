package main

import (
	"fmt"
	"micro/handler"
	"net"
	addservice "micro/proto"
	"google.golang.org/grpc"
)

func main(){
	lis,err :=net.Listen("tcp",":8000")
	if err != nil{
		fmt.Println(err)
	}
	defer lis.Close()

	addServ := handler.AddServiceServer1{}

	grpcServer := grpc.NewServer()
	addservice.RegisterAddServiceServer(grpcServer,&addServ)

	if err := grpcServer.Serve(lis);err !=nil{
		fmt.Println(err)
	}
}
