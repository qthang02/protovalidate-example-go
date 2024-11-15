package main

import (
	"fmt"
	"github.com/bufbuild/protovalidate-go"
	"github.com/qthang02/protovalidate-example/protobuf/greetPb"
)

func main() {
	validator, err := protovalidate.New()
	if err != nil {
		panic(err)
	}

	req := &greetPb.GreetRequest{
		Name: "John",
		Age:  25,
	}

	if err := validator.Validate(req); err != nil {
		fmt.Println("Validation failed:", err)
	} else {
		fmt.Println("Validation passed!")
	}

	invalidReq := &greetPb.GreetRequest{
		Name: "Jo",
		Age:  16,
	}

	if err := validator.Validate(invalidReq); err != nil {
		fmt.Println("Validation failed for invalidReq:", err)
	}
}
