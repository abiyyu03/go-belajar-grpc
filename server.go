package main

import (
	"fmt"
	"implement-grpc/handler"
	"implement-grpc/proto/profile"
	"log"
	"net"

	"google.golang.org/grpc"
)

func main() {
	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", 9000))
	if err != nil {
		log.Fatalf("Failed to listen : %v", err)
	}

	s := grpc.NewServer()

	profile.RegisterProfileServiceServer(s, &handler.Profile{})

	if err := s.Serve(lis); err != nil {
		log.Fatalf("Failed to serve : %s", err)
	}
}
