package main

import (
	"fmt"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
	"log"
	"../calculatorpb"
)

func main()  {
	fmt.Println("Client is running...")

	cc, err := grpc.Dial("localhost:50051", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("could not connect to server: %v", err)
	}
	defer cc.Close()

	c:= calculatorpb.NewCalculatorServiceClient(cc)

	doSum(c)
}

func doSum(c calculatorpb.CalculatorServiceClient)  {
	fmt.Println("Starting to do a sum RPC")

	req := &calculatorpb.SumRequest{
		FirstNumber: 10,
		SecondUmber: 15,
	}

	res, err := c.Sum(context.Background(), req)
	if err != nil {
		log.Fatalf("Error while calling sum RPC: %v", err)
	}

	log.Printf("Response from server: %v", res.SumResult)
}
