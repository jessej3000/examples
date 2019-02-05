package main

import (
	"context"
	"fmt"
	"log"

	"github.com/jessej3000/examples/testgrpc/greet/greetpb"
	"google.golang.org/grpc"
)

func main() {
	fmt.Println("Start client.")

	con, err := grpc.Dial("localhost:50051", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("Coould not connect: %v", err)
	}
	defer con.Close()

	c := greetpb.NewGreetServiceClient(con)

	doUnary(c)

}

func doUnary(c greetpb.GreetServiceClient) {
	req := &greetpb.GreetRequest{
		Greeting: &greetpb.Greeting{
			FirstName: "Jesse",
			LastName:  "Junsay",
		},
	}
	res, err := c.Greet(context.Background(), req)
	if err != nil {
		log.Fatalf("Error while calling Greet %v", err)
	}
	log.Printf("Response from Greet: %v", res.Result)
}
