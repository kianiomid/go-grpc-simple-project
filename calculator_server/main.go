package main

import (
	"context"
	"fmt"
	"google.golang.org/grpc"
	"../calculatorpb"
	"log"
	"net"
)

type server struct{}

func main() {
	fmt.Println("Server is running...")

	// Make a listener
	lis, err := net.Listen("tcp", "localhost:50051")
	if err != nil {
		log.Fatalf("Failed to listen: %v", err)
	}

	// Make a gRPC server
	grpcServer := grpc.NewServer()
	calculatorpb.RegisterCalculatorServiceServer(grpcServer, &server{})

	// Run the gRPC server
	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("Failed to serve: %v", err)
	}
}

func (*server) Sum(ctx context.Context, req *calculatorpb.SumRequest) (*calculatorpb.SumResponse, error) {
	fmt.Printf("Received Sum RPC: %v", req)

	firstNumber := req.GetFirstNumber()
	secondNumber := req.GetSecondUmber()

	sum := firstNumber + secondNumber

	res := &calculatorpb.SumResponse{
		SumResult: sum,
	}

	return res, nil
}