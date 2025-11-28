package main

import (
	"context"
	"io"
	"log"
	"time"

	pb "github.com/harsh2971/grpc-go/proto"
)

func callSayHelloBidirectionalStreaming(client pb.GreetServiceClient, names *pb.NamesList) {
	log.Printf("Bidirectional streaming started")

	ctx, cancel := context.WithTimeout(context.Background(), 20*time.Second)
	defer cancel()

	stream, err :=client.SayHelloBidirectionalStreaming(ctx)
	if err != nil {
		log.Fatalf("error while calling SayHelloBidirectionalStreaming: %v", err)
	}

	waitc := make(chan struct{})//channel to wait for the streaming to finish

	go func() {
		for {
			// receive a stream of responses from the server
			message, err := stream.Recv()
			if err == io.EOF {
				break
			}
			if err != nil {
				log.Fatalf("error while receiving message: %v", err)
			}
			log.Printf("Received: %s", message.Message)
		}
		close(waitc)//close the channel to signal the main goroutine that the streaming is finished
	}()

	// send a stream of requests to the server
	for _, name := range names.Names {
		req:= &pb.HelloRequest{
			Name: name,
		}
		if err := stream.Send(req); err!= nil {
			log.Fatalf("error while sending request: %v", err)
		}
		log.Printf("Sent request with name: %s", name)
		time.Sleep(2*time.Second)
	}
	// close the send side of the stream to signal the server that the client has finished sending requests
	stream.CloseSend()
	// wait for the streaming to finish
	<-waitc
	log.Printf("Bidirectional streaming finished")

}