package main

import (
	"io"
	"log"

	pb "github.com/harsh2971/grpc-go/proto"
)


func (s *helloServer) SayHelloClientStreaming(stream pb.GreetService_SayHelloClientStreamingServer) error {
	messages := []string{}

	for {
		// receive a stream of requests from the client
		req,err := stream.Recv()
		if err == io.EOF {
			return stream.SendAndClose(&pb.MessagesList{Messages: messages})
		}
		if err != nil {
			return err
		}
		log.Printf("Got request with name: %s", req.Name)
		messages = append(messages, "Hello "+req.Name)
	}
}