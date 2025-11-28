package main

import (
	"io"
	"log"
	"time"
	pb "github.com/harsh2971/grpc-go/proto"
)


func (s *helloServer) SayHelloBidirectionalStreaming(stream pb.GreetService_SayHelloBidirectionalStreamingServer) error {
	log.Printf("Bidirectional streaming started")

	for {
		req, err := stream.Recv()
		if err == io.EOF{
			break
		}
		if err != nil {
			log.Fatalf("error while receiving request: %v", err)
		}
		log.Printf("Got request with name: %s", req.Name)
		res := &pb.HelloResponse{
			Message: "Hello " + req.Name,
		}
		if err := stream.Send(res); err != nil {
			log.Fatalf("error while sending response: %v", err)
		}
		log.Printf("Sent response with message: %s", res.Message)
		time.Sleep(2*time.Second)
	}
	log.Printf("Bidirectional streaming finished")
	return nil
}