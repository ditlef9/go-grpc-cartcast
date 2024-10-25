// cartcast/client/delete.go

package main

import (
	"context"
	pb "ekeberg.com/go-grpc-cartcast/cartcast/proto"
	"log"
)

func DeleteCartCast(c pb.CartCastServiceClient, id string) {
	log.Println("client/delete.go::DeleteCartCast()::Invoked")

	_, err := c.DeleteCartCast(context.Background(), &pb.CartCastId{Id: id})

	if err != nil {
		log.Fatalf("client/delete.go::DeleteCartCast()::Error while updating: %v\n", err)
	}

	log.Println("client/delete.go::DeleteCartCast()::Delete successfully")

}
