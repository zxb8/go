package handler

import (
	"context"
	 addservice "micro/proto"
)


type AddServiceServer1 struct{}

func (s *AddServiceServer1)Add(ctx context.Context, in *addservice.AddRequest) (*addservice.AddResponse, error){

	a,b := in.A,in.B

	return &addservice.AddResponse{Result:a+b},nil
}
