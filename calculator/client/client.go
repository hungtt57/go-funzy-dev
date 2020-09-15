package main

import (
	"context"
	"google.golang.org/grpc"

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
	calllSum(client)
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
