package main

import (
	"com.grpc.tleu/greet/greetpb"
	"google.golang.org/grpc"
	"io"
	"log"
	"net"
)

type Server struct {
	greetpb.UnimplementedGreetServiceServer
}

func (s *Server) LongGreet(stream greetpb.GreetService_LongGreetServer) error {
	var sum float32
	var result float32
	var counter float32 = 0

	for {
		request, err := stream.Recv()
		if err == io.EOF {
			result = sum / counter
			return stream.SendAndClose(&greetpb.LongGreetResponse{
				Result: result,
			})
		}else if err != nil {
			log.Fatalf("Error while reading client stream: %v", err)
		}
		numberFromStream := request.Greeting.GetNumber()
		sum += numberFromStream

		counter += 1
	}
}

func main() {
	l, err := net.Listen("tcp", "0.0.0.0:50051")
	if err != nil {
		log.Fatalf("Failed to listen:%v", err)
	}
	s := grpc.NewServer()
	greetpb.RegisterGreetServiceServer(s, &Server{})
	log.Println("Server is running on port:50051")
	if err := s.Serve(l); err != nil {
		log.Fatalf("failed to serve:%v", err)
	}
}
