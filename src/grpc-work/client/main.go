package main

import (
	"fmt"
	"golang.org/x/net/context"
	"log"
	"google.golang.org/grpc"
	"grpc-work/model"
)

var (
	servel = "localhost:8080"
)


func main(){
	var opts []grpc.DialOption

	opts = append(opts,grpc.WithInsecure())

	conn,err :=grpc.Dial(servel,opts...)
	if err != nil{
		log.Fatal(fmt.Errorf("Unable to connect to grpc service ：%v",err))
	}
	defer conn.Close()
	client := model.NewMyMathServiceClient(conn)
	ctx := context.Background()

	in := &model.MathRequest{Operand1:11,Operand2:12}


	result,err := client.Add(ctx,in)

	if err != nil{
		log.Fatal(fmt.Errorf("request add error ：%v",err))
		return
	}

	fmt.Println(result)

}


