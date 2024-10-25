// cartcast/server/list.go

package main

import (
	"context"
	pb "ekeberg.com/go-grpc-cartcast/cartcast/proto"
	"fmt"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/emptypb"
	"log"
)

func (s *Server) ListCartCast(in *emptypb.Empty, stream pb.CartCastService_ListCartCastServer) error {
	log.Println("server/list.go::ListCartCast()::Invoked")

	cur, err := collection.Find(context.Background(), primitive.D{})

	if err != nil {
		return status.Error(
			codes.Internal,
			fmt.Sprintf("server/list.go::ListCartCast()::Unknown internal error %v", err),
		)
	}
	defer cur.Close(context.Background())

	for cur.Next(context.Background()) {
		data := &CartCastItem{}
		err := cur.Decode(data)
		if err != nil {
			return status.Errorf(
				codes.Internal,
				fmt.Sprintf("server/list.go::ListCartCast()::Decode error: %v", err),
			)
		}

		stream.Send(documentToCartCast(data))
	}

	if err := cur.Err(); err != nil {
		return status.Errorf(
			codes.Internal,
			fmt.Sprintf("server/list.go::ListCartCast()::Unknown internal error: %v", err),
		)
	}

	return nil
}
