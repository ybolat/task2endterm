package main

import (
	"context"
	"fmt"
	"log"
	"time"

	"com.grpc.tleu/greet/greetpb"
	"google.golang.org/grpc"
)

func main() {
	conn, err := grpc.Dial("localhost:50051", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("could not connect: %v", err)
	}
	defer conn.Close()

	c := greetpb.NewGreetServiceClient(conn)
	doLongGreet(c)
}

func doLongGreet(c greetpb.GreetServiceClient) {

	//var size int
	//fmt.Print("Enter count of numbers:")
	//fmt.Scan(&size)

	//	for i := 0; i < size; i++{
	//	fmt.Print("Enter the number:")
	//	fmt.Scan(&number)
	//	requests = append(requests, i)
	//}

	var number1 float32
	var number2 float32
	var number3 float32
	var number4 float32

	fmt.Print("Enter 4 numbers:")
	fmt.Scan(&number1, &number2, &number3, &number4)

	requests := []*greetpb.LongGreetRequest{
		{
			Greeting: &greetpb.Greeting{
				Number: number1,
			},
		},
		{
			Greeting: &greetpb.Greeting{
				Number: number2,
			},
		},
		{
			Greeting: &greetpb.Greeting{
				Number: number3,
			},
		},
		{
			Greeting: &greetpb.Greeting{
				Number: number4,
			},
		},
	}

	ctx := context.Background()
	stream, err := c.LongGreet(ctx)
	if err != nil {
		log.Fatalf("error while calling LongGreet: %v", err)
	}

	for _, req := range requests {
		fmt.Printf("Sending number: %v\n", req)
		stream.Send(req)
		time.Sleep(1500 * time.Millisecond)
	}

	response, err := stream.CloseAndRecv()
	if err != nil {
		log.Fatalf("error while receiving response from LongGreet: %v", err)
	}
	fmt.Printf("Avarage of numbers is: %v\n", response)
}
