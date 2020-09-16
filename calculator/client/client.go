package main

import (
	"context"
	"google.golang.org/grpc"
	"io"

	"log"

	"github.com/hungtt57/go-funzy-dev/calculator/calculatorpb"
)

func main() {
	cc, err := grpc.Dial("localhost:50069", grpc.WithInsecure()) //connect den 50069

	if err != nil {
		log.Fatalf(" err while dial %v", err)
	}

	defer cc.Close() //chay cuoi cung main

	client := calculatorpb.NewCalculatorServiceClient(cc)
	//calllSum(client)
	callPND(client)
}

func calllSum(c calculatorpb.CalculatorServiceClient) {
	log.Println("Calling sum api")
	resp, err := c.Sum(context.Background(), &calculatorpb.SumRequest{
		Num1: 5,
		Num2: 8,
	})
	if err != nil {
		log.Fatal("call sum api err %v", err)
	}
	log.Printf("sum api response %v", resp.GetResult())
}

func callPND(c calculatorpb.CalculatorServiceClient) {
	log.Println("Calling callPND")
	stream, err := c.PrimeNumberDecomposition(context.Background(), &calculatorpb.PNDRequest{
		Number: 120,
	})

	if err != nil {
		log.Fatalf("calPND err %v", err)
	}

	for {
		resp, recErr := stream.Recv()
		if recErr == io.EOF {
			log.Println("Server finish streaming")
			return
		}

		log.Printf("prime number %v", resp.GetResult())
	}
}
