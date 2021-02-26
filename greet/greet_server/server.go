package main

import (
	"com.grpc.tleu/greet/greetpb"
	"context"
	"fmt"
	"google.golang.org/grpc"
	"log"
	"net"
	"strconv"
)

type Server struct {
	greetpb.UnimplementedGreetServiceServer
}

func (s *Server) Greet(ctx context.Context, req *greetpb.GreetRequest) (*greetpb.GreetResponse, error) {
	fmt.Printf("Greet function was invoked with %v \n", req)

	var x= int(req.GetGreeting().GetNumber())
	var result string

	for i := 2;x > i; i++  {
		for x%i == 0 {
			var a = strconv.Itoa(int(i))
			result += a + " "
			x /= i
		}
	}
	if x >2 {
		//last is last prime number
		var last = strconv.Itoa(x)
		result += last
	}

	res := &greetpb.GreetResponse{
		Result: result,
	}

	return res, nil
}

func main() {
	l, err := net.Listen("tcp", "0.0.0.0:50051")
	if err != nil {
		log.Fatalf("Failed to listen:%v", err)
	}
	s := grpc.NewServer()
	greetpb.RegisterGreetServiceServer(s, &Server{})
	log.Println("Server is running on port:50051")
	if err := s.Serve(l); err != nil {
		log.Fatalf("failed to serve:%v", err)
	}
}
