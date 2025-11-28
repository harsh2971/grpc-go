package main

import (
	"log"
	"net"
	pb "github.com/harsh2971/grpc-go/proto"
	"google.golang.org/grpc"
)
// port for the server
const (
	port= ":8080"
)

// helloServer is the server implementation for the greet service
type helloServer struct {
	pb.GreetServiceServer
}

func main() {
	lis, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalf("Failed to start the server: %v", err)
	}

	//create a new grpc server
	grpcServer := grpc.NewServer()
	// register the service 
	pb.RegisterGreetServiceServer(grpcServer, &helloServer{})
	//start the server
	log.Printf("sever started at %v", lis.Addr())
	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("failed to start %v", err)
	}

}
