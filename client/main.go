package main

import (
	"log"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure" // insecure credentials for the connection
	pb "github.com/harsh2971/grpc-go/proto"
)
const (
	port = ":8080"
)

// main function to call the greet service
func main() {
	conn, err := grpc.Dial("localhost"+port, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("Failed to connect: %v", err)
	}
	defer conn.Close()

	client := pb.NewGreetServiceClient(conn)

	names:= &pb.NamesList{
		Names: []string{"Harsh", "Sweta", "John", "Jane", "Jim", "Jill" },
	}
	// log.Printf("NamesList: %v", names)
 
	//callSayHello(client)
	//callSayHelloServerStreaming(client, names)
	//callSayHelloClientStreaming(client, names)
	callSayHelloBidirectionalStreaming(client, names)
}
