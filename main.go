package main

import (
	"context"
	"errors"
	"github.com/bufbuild/protovalidate-go"
	"github.com/qthang02/protovalidate-example/protobuf/greetPb"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
	"log"
	"net"
)

type server struct {
	*greetPb.UnimplementedGreetServiceServer
}

func (s *server) Greet(ctx context.Context, req *greetPb.GreetRequest) (*greetPb.GreetResponse, error) {
	validator, err := protovalidate.New()
	if err != nil {
		panic(err)
	}

	if err := validator.Validate(req); err != nil {
		return nil, errors.New("validation failed: " + err.Error())
	} else {
		return nil, err
	}

	return nil, nil
}

func main() {

	lis, err := net.Listen("tcp", ":8080")
	if err != nil {
		log.Fatalln("Failed to listen:", err)
	}

	s := grpc.NewServer()

	greetPb.RegisterGreetServiceServer(s, &server{})
	reflection.Register(s)

	log.Println("Serving gRPC on 0.0.0.0:8080")
	log.Fatal(s.Serve(lis))
}
