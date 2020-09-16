package main

import (
	"context"
	"fmt"
	"github.com/hungtt57/go-funzy-dev/calculator/calculatorpb"
	"google.golang.org/grpc"
	"log"
	"net"
	"time"
)

type server struct{}

func (*server) Sum(context context.Context, req *calculatorpb.SumRequest) (*calculatorpb.SumResponse, error) {
	log.Println("Sum called...")
	resp := &calculatorpb.SumResponse{
		Result: req.GetNum1() + req.GetNum2(),
	}
	return resp, nil
}

func (*server) PrimeNumberDecomposition(req *calculatorpb.PNDRequest,
	stream calculatorpb.CalculatorService_PrimeNumberDecompositionServer) error {
	log.Println("PrimeNumberDecomposition called...")
	k := int32(2)
	N := req.GetNumber()
	for N > 1 {
		if N%k == 0 {
			N = N / k
			//send to client
			stream.Send(&calculatorpb.PDNResponse{
				Result: k,
			})
			time.Sleep(1000 * time.Millisecond)
		} else {
			k++
			log.Printf("k increase to %v", k)
		}
	}
	return nil
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
