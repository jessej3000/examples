package main

import (
	"log"
	"net"

	"golang.org/x/net/context"

	"github.com/jessej3000/examples/testgrpc/math/mathpb"
	"google.golang.org/grpc"
)

type server struct{}

func (*server) Compute(ctx context.Context, req *mathpb.NumbersRequest) (*mathpb.NumbersResponse, error) {
	num1 := req.GetNumbers().GetNum1()
	num2 := req.GetNumbers().GetNum2()
	res := &mathpb.NumbersResponse{
		Result: num1 + num2,
	}

	return res, nil
}

func main() {
	lis, err := net.Listen("tcp", "localhost:50051")
	if err != nil {
		log.Fatal(err)
	}

	rpc := grpc.NewServer()

	mathpb.RegisterNumbersServiceServer(rpc, &server{})

	if err := rpc.Serve(lis); err != nil {
		log.Fatalf("Error to serve. %v", err)
	}

}
