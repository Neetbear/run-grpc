package main

import (
	"context"
	"log"
	"net"
	"net/http"
	pb "run-grpc/pkg"

	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"google.golang.org/grpc"
)

const (
	host     = "localhost"
	grpcport = ":8080"
	muxport  = ":8081"
)

type helloServer struct {
	pb.GreetServiceServer
}

func main() {
	go func() {
		// mux
		mux := runtime.NewServeMux()

		// register
		pb.RegisterGreetServiceHandlerServer(context.Background(), mux, &helloServer{})

		// http server
		log.Fatalln(http.ListenAndServe(host+muxport, mux))
	}()

	lis, err := net.Listen("tcp", grpcport)
	if err != nil {
		log.Fatalf("Failed to start the server %v", err)
	}

	grpcServer := grpc.NewServer()
	pb.RegisterGreetServiceServer(grpcServer, &helloServer{})
	log.Printf("Server started at %v", lis.Addr())
	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("Failed to start: %v", err)
	}
}
