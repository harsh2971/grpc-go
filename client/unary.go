package main

import (
	"context"
	"log"
	"time"
	pb "github.com/harsh2971/grpc-go/proto"
)

func callSayHello(client pb.GreetServiceClient) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	res,err:=client.SayHello(ctx, &pb.NoParam{})
	if err != nil {
		log.Fatalf("Error while calling SayHello: %v", err)
	}
	log.Printf("Response from SayHello: %s", res.Message)

}
