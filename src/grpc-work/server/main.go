package main

import (
	"context"
	"fmt"
	"log"
	"google.golang.org/grpc"
	"grpc-work/model"
	"net"
)
var (
	address = "localhost:8080"
)


type (
	MathService struct{}
)

func (m *MathService) Add(ctx context.Context,in *model.MathRequest) (*model.MathResponse,error){
	if m == nil{
		return nil,fmt.Errorf("Add Called on nil object")
	}
	if in == nil{
		return nil,fmt.Errorf("Add called with invalid paramter value nil")
	}

	result := &model.MathResponse{}

	result.Result = in.Operand1 + in.Operand2

	return result,nil
}


func main(){
	var opts []grpc.ServerOption
	server :=grpc.NewServer(opts ...)


	model.RegisterMyMathServiceServer(server,&MathService{})

	lis,err :=net.Listen("tcp",address)

	if err != nil{
		log.Fatal(fmt.Errorf("unable to start"))
	}

	server.Serve(lis)
}
