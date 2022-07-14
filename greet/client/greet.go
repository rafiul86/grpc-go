package main

import (
	"context"
	"log"

	pb "github.com/rafiul86/grpc-go/greet/proto"
)

 func doGreet(c pb.GreetServiceClient) {
	log.Printf("Starting to do a Greet RPC...")
	res, err := c.Greet(context.Background(), &pb.GreetRequest{
			FirstName: "Rafiul",
		})

	if err != nil {
		log.Fatalf("Error while calling Greet RPC: %v", err)
	}
	log.Printf("Response from Greet: %v", res.Result)
 }

 