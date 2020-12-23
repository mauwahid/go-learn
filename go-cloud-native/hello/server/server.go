package main

import (
	"context"
	"fmt"
	"log"
	"net"

	"github.com/mauwahid/hello/hellopb"
	"google.golang.org/grpc"
)

type Server struct{}

func (*Server) Hello(ctx context.Context, request *hellopb.HelloRequest) (*hellopb.HelloResponse, error) {
	name := request.Name
	response := &hellopb.HelloResponse{
		Greeting: "Hello " + name,
	}

	return response, nil
}

func main() {
	address := "0.0.0.0:50051"
	lis, err := net.Listen("tcp", address)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("Server is listening on %v ...", address)

	s := grpc.NewServer()
	hellopb.RegisterHelloServiceServer(s, &Server{})
	s.Serve(lis)
}
