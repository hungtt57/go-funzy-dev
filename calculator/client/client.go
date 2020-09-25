package main

import (
	"context"
	"google.golang.org/grpc"
	"io"
	"time"

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
	//callPND(client)
	//callAverage(client)
	callFindMax(client)
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

func callAverage(c calculatorpb.CalculatorServiceClient) {
	log.Println("Calling callAverage")
	stream, err := c.Average(context.Background())

	listReqs := []calculatorpb.AverageRequest{
		calculatorpb.AverageRequest{
			Num: 5,
		},
		calculatorpb.AverageRequest{
			Num: 6.2,
		},
		calculatorpb.AverageRequest{
			Num: 7.2,
		},
	}
	if err != nil {
		log.Fatalf("call average err %v", err)
	}

	for _, req := range listReqs {
		err := stream.Send(&req)
		if err != nil {
			log.Fatalf("Send average request err %v", err)
		}
		time.Sleep(500 * time.Millisecond)
	}

	resp, err := stream.CloseAndRecv()

	if err != nil {
		log.Fatalf("Receive average resp err %v", err)

	}

	log.Printf("Receive average resp %v", resp)
}

func callFindMax(c calculatorpb.CalculatorServiceClient) {
	log.Println("Calling callFindMax.................")
	stream, err := c.FindMax(context.Background())
	if err != nil {
		log.Fatalf("call find max err %v", err)
	}

	waitc := make(chan struct{})

	go func() {
		listReqs := []calculatorpb.FindMaxRequest{
			calculatorpb.FindMaxRequest{
				Num: 5,
			},
			calculatorpb.FindMaxRequest{
				Num: 6,
			},
			calculatorpb.FindMaxRequest{
				Num: 7,
			},
			calculatorpb.FindMaxRequest{
				Num: 4,
			},
		}

		//gui nhieu request
		for _, req := range listReqs {
			err := stream.Send(&req)
			if err != nil {
				log.Fatalf("Send max request err %v", err)
				break
			}
			time.Sleep(500 * time.Millisecond)
		}

		stream.CloseSend()
	}()

	go func() {
		for {
			resp, err := stream.Recv()
			if err == io.EOF {
				log.Println("ending find max api...")
				break
			}
			if err != nil {
				log.Fatalf("resp find max err %v", err)
			}
			log.Printf("print MAX %v", resp.GetMax())
		}
		close(waitc)
	}()

	<-waitc
}
