// cartcast/client/read.go
// Read a CartCast

package main

import (
	"context"
	pb "ekeberg.com/go-grpc-cartcast/cartcast/proto"
	"log"
)

func readCartCast(c pb.CartCastServiceClient, id string) *pb.CartCast {
	log.Println("client/read.go::readCartCast()::Invoked -----------------------")

	req := &pb.CartCastId{Id: id}
	res, err := c.ReadCartCast(context.Background(), req)

	if err != nil {
		log.Printf("client/read.go::readCartCast()::Error while reading: %v\n", err)
	}

	if res != nil {
		log.Printf("client/read.go::readCartCast()::CartCast read: %v\n", res)
	} else {
		log.Println("client/read.go::readCartCast()::Received nil response")
	}

	return res
}
