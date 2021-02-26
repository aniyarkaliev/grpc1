package main

import (
	"com.grpc.tleu/greet/greetpb"
	"context"
	"google.golang.org/grpc"
	"log"
)

func main() {

	conn, err := grpc.Dial("localhost:50051", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("could not connect: %v", err)
	}
	defer conn.Close()

	c := greetpb.NewGreetServiceClient(conn)
	doUnary(c)
}

func doUnary(c greetpb.GreetServiceClient) {
	ctx := context.Background()

	request := &greetpb.GreetRequest{Greeting: &greetpb.Greeting{
		Number: 120,
	}}

	response, err := c.Greet(ctx, request)
	if err != nil {
		log.Fatalf("error while calling Greet RPC %v", err)
	}
	log.Printf("response from Greet:%v", response.Result)
}