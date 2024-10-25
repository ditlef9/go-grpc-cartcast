// cartcast/client/list

package main

import (
	"context"
	pb "ekeberg.com/go-grpc-cartcast/cartcast/proto"
	"google.golang.org/protobuf/types/known/emptypb"
	"io"
	"log"
)

func ListCartCast(c pb.CartCastServiceClient) {
	log.Println("client/list.go::ListCartCast()::Invoked ------------------")

	stream, err := c.ListCartCast(context.Background(), &emptypb.Empty{})

	if err != nil {
		log.Fatalf("client/list.go::ListCartCast()::Error listing CartCasts: %v: ", err)
	}

	for {
		res, err := stream.Recv()

		// Finished
		if err == io.EOF {
			break
		}

		if err != nil {
			log.Fatalf("client/list.go::ListCartCast()::Error in for: %v: ", err)
		}

		log.Println(res)
	}

}
