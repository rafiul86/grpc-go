package main

import (
	"log"
	"net"

	pb "github.com/rafiul86/grpc-go/greet/proto"
	"google.golang.org/grpc"
)

type Server struct {
	pb.GreetServiceServer
}

var addr string = "0.0.0.0:50051"

func main() {
	lis, err := net.Listen("tcp", addr)

	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	log.Printf("listening on %s", addr)

	s := grpc.NewServer()

	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}