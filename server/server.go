package main

import (
	"context"
	"fmt"
	"github.com/hvn2001/grpc_demo/calculatorpb"
	"google.golang.org/grpc"
	"log"
	"net"
)

type server struct{}

func (*server) Sum(ctx context.Context, req *calculatorpb.SumRequest) (*calculatorpb.SumResponse, error) {
	log.Println("sum called...")
	resp := &calculatorpb.SumResponse{
		Result: req.GetNum1() + req.GetNum2(),
	}

	return resp, nil
}

func main() {

	lis, err := net.Listen("tcp", "0.0.0.0:50069")
	if err != nil {
		log.Fatalf("err while create listen %v", err)
	}

	s := grpc.NewServer()

	calculatorpb.RegisterCalculatorServiceServer(s, &server{})

	fmt.Println("calculator is running...")
	err = s.Serve(lis)

	if err != nil {
		log.Fatalf("err while serve %v", err)
	}
}
