package main

import (
	"google.golang.org/grpc"
	"grpc-service-01/metadata"
	"log"
	"net"
)

func main() {
	lis, err := net.Listen("tcp", ":9000")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	s := metadata.Server{}

	grpcServer := grpc.NewServer()

	metadata.RegisterMetadataServer(grpcServer, &s)

	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %s", err)
	}
}