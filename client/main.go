package main

import (
	"context"
	"fmt"
	"github.com/qthang02/protovalidate-example/protobuf/greetPb"
	"google.golang.org/grpc"
	"log"
)

func main() {
	conn, err := grpc.Dial("localhost:8080", grpc.WithInsecure())
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close()

	client := greetPb.NewGreetServiceClient(conn)

	_, err = client.Greet(context.Background(), &greetPb.GreetRequest{
		Name:   "John",
		Age:    34,
		Status: 99,
	})
	if err != nil {
		fmt.Println(err)
	}
}
