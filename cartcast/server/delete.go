// cartcast/server/delete.go
package main

import (
	"context"
	pb "ekeberg.com/go-grpc-cartcast/cartcast/proto"
	"fmt"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/emptypb"
	"log"
)

func (*Server) DeleteCartCast(ctx context.Context, in *pb.CartCastId) (*emptypb.Empty, error) {
	log.Printf("server/delete.go::DeleteCartCast()::Invoked id: %v\n", in)

	// Transform to object ID
	oid, err := primitive.ObjectIDFromHex(in.Id)
	if err != nil {
		return nil, status.Errorf(codes.InvalidArgument, "server/delete.go::DeleteCartCast()::Can not cast id to ObjectID")
	}

	res, err := collection.DeleteOne(ctx, bson.M{"_id": oid})
	if err != nil {
		return nil, status.Errorf(
			codes.Internal,
			fmt.Sprintf("server/delete.go::DeleteCartCast()::Can not delete object"),
		)
	}

	if res.DeletedCount == 0 {
		return nil, status.Errorf(
			codes.NotFound,
			fmt.Sprintf("server/delete.go::DeleteCartCast()::CartCast not found"),
		)
	}

	return &emptypb.Empty{}, nil
}
