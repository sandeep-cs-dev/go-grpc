package main

import (
	"context"
	"fmt"
	"log"

	"github.com/sandeep-cs-dev/grpc/server/proto"
	"google.golang.org/grpc"
)

func client() {

	conn, err := grpc.Dial("localhost:8080", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("failed to connect: %v", err)
	}

	defer conn.Close()
	c := proto.NewMathServiceClient(conn)

	// numbers to add
	num := &proto.AddRequest{
		Number1: 10,
		Number2: 10,
	}

	res, err := c.Add(context.Background(), num)

	if err != nil {
		log.Fatalf("error calling MathService Add function %v", err)
	}
	fmt.Println(res.GetSum())

}
