package main

import (
	"google.golang.org/grpc"
	"log"
	"net"
)

type server struct {

}

func main() {
	lis, err := net.Listen("tcp", "0.0.0.0:50069")

	if err != nil {
		log.Fatal("Err while create listen %v", err)
	}

	s := grpc.NewServer()

	calculatorpb.RegisterCalculatorServiceServer(s, &server{})

	err = s.Serve(lis)

	if err != nil {
		log.Fatal("Err while create listen %v", err)
	}
}
