package main

import (
	"fmt"
	"context"
	"google.golang.org/grpc"
	addservice "micro/proto"
)

func main(){
	conn,err :=grpc.Dial(":8000",grpc.WithInsecure())

	if err != nil{
		fmt.Println(err)
	}

	defer conn.Close()

	addServ := addservice.NewAddServiceClient(conn)
	res,err := addServ.Add(context.Background(),&addservice.AddRequest{A:1,B:2})

	if err != nil{
		fmt.Println(err)
	}else{
		fmt.Println("result1:",res.Result)
	}
}
