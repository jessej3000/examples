package main

import (
	"context"
	"fmt"
	"log"

	"github.com/jessej3000/examples/testgrpc/math/mathpb"
	"google.golang.org/grpc"
)

func main() {
	fmt.Println("Start client.")

	con, err := grpc.Dial("localhost:50051", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("Coould not connect: %v", err)
	}
	defer con.Close()

	c := mathpb.NewNumbersServiceClient(con)

	doUnary(c)

}

func doUnary(c mathpb.NumbersServiceClient) {
	req := &mathpb.NumbersRequest{
		Numbers: &mathpb.Numbers{
			Num1: 1,
			Num2: 2,
		},
	}
	res, err := c.Compute(context.Background(), req)
	if err != nil {
		log.Fatalf("Error while calling Greet %v", err)
	}
	log.Printf("Response from Greet: %v", res.Result)
}
