package main

import (
	"context"
	"fmt"

	pb "github.com/rafiul86/grpc-go/greet/proto"
)

func (s *Server) Greet(ctx context.Context, in *pb.GreetRequest) (*pb.GreetResponse, error) {
	fmt.Printf("Greet function was invoked with %v\n", in)
	firstName := in.FirstName
	result := "Hello " + firstName
	resp := &pb.GreetResponse{
		Result: result,
	}
	return resp, nil
}