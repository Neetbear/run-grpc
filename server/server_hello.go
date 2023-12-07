package main

import (
	"context"
	"log"
	pb "run-grpc/pkg"
)

func (s *helloServer) SayHelloServer(ctx context.Context, req *pb.HelloRequest) (*pb.HelloResponse, error) {
	log.Printf("got request with name: %v", req.Name)
	if err := req.Validate(); err != nil {
		return nil, err
	}

	return &pb.HelloResponse{
		Message: "Hello " + req.Name,
	}, nil
}
