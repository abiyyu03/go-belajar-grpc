package main

import (
	"context"
	"implement-grpc/proto/profile"
	"log"

	"google.golang.org/grpc"
)

func main() {
	var conn *grpc.ClientConn
	conn, err := grpc.Dial(":9000", grpc.WithInsecure())

	if err != nil {
		log.Fatalf("did not connect: %s", err)
	}

	defer conn.Close()

	c := profile.NewProfileServiceClient(conn)

	createReq := profile.CreateRequest{Name: "Abiyyu Cakra", Id: 123, IsVerified: true}
	response, err := c.Create(context.Background(), &createReq)
	if err != nil {
		log.Printf("Error creating profile: %s", err)
	}
	log.Printf("Response from Server: %s", response.Message)

}
