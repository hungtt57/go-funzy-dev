package main

import (
	"context"
	"fmt"
	"google.golang.org/grpc"
	"log"
	"net"
	"github.com/hungtt57/go-funzy-dev/calculator/calculatorpb"
)

type server struct {}

func (*server) Sum(context context.Context,req *calculatorpb.SumRequest) (*calculatorpb.SumResponse, error) {
	log.Println("Sum called...")
	resp := &calculatorpb.SumResponse{
		Result: req.GetNum1() + req.GetNum2(),
	}
	return resp, nil
}

func main() {
	lis, err := net.Listen("tcp", "0.0.0.0:50069")

	if err != nil {
		log.Fatal("Err while create listen %v", err)
	}

	s := grpc.NewServer()

	calculatorpb.RegisterCalculatorServiceServer(s, &server{})
	fmt.Println("Server running")
	err = s.Serve(lis)

	if err != nil {
		log.Fatal("Err while create listen %v", err)
	}
}
