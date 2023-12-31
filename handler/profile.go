package handler

import (
	"context"
	"implement-grpc/proto/profile"
	"log"
)

type Profile struct {
	profile.UnimplementedProfileServiceServer
}

func (p *Profile) Create(ctx context.Context, req *profile.CreateRequest) (*profile.CreateResponse, error) {
	log.Printf("Receive message body from client %s", req.GetName())
	return &profile.CreateResponse{Message: "Profile Created!"}, nil
}
