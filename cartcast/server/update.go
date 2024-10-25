// cartcast/server/update.go
package main

import (
	"context"
	pb "ekeberg.com/go-grpc-cartcast/cartcast/proto"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/emptypb"
	"log"
)

func (s *Server) UpdateCartCast(ctx context.Context, in *pb.CartCast) (*emptypb.Empty, error) {
	log.Printf("server/update.go::UpdateCartCast()::Invoked id: %v\n", in)

	// Transform to object ID
	oid, err := primitive.ObjectIDFromHex(in.Id)
	if err != nil {
		return nil, status.Errorf(codes.InvalidArgument, "server/update.go::UpdateCartCast()::Can not cast id to ObjectID")
	}

	data := &CartCastItem{
		Title:    in.Title,
		XCord:    in.XCord,
		YCord:    in.YCord,
		ImageURL: in.ImageUrl,
		VideoURL: in.VideoUrl,
	}
	res, err := collection.UpdateOne(ctx, bson.M{"_id": oid}, bson.M{"$set": data})
	if err != nil {
		return nil, status.Errorf(
			codes.Internal, "server/update.go::UpdateCartCast()::Can not update object")
	}

	if res.MatchedCount == 0 {
		return nil, status.Errorf(
			codes.NotFound, "server/update.go::UpdateCartCast()::Can not find CartCast with Id")
	}

	return &emptypb.Empty{}, nil
}
