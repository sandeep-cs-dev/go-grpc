package main

import (
	"fmt"
	"log"
	"net"

	"github.com/sandeep-cs-dev/grpc/server/proto"
	"github.com/sandeep-cs-dev/grpc/server/services"
	"google.golang.org/grpc"
)

const (
	port = ":8080"
)

func main() {
	fmt.Println("running main.go")
	lis, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalf("tcp dail failed at 8080: %v", err)
	}
	grpcServer := grpc.NewServer()
	proto.RegisterMathServiceServer(grpcServer, &services.MathServiceServer{})
	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("grpc server serve failed %v", err)
	}
}
