// cartcast/server/create.go
// This file contains the implementation of the CreateCartCast method
// which handles the creation of new CartCast entries in the MongoDB database.

package main

import (
	"context"
	pb "ekeberg.com/go-grpc-cartcast/cartcast/proto" // Importing the CartCast protocol buffer definitions
	"fmt"
	"go.mongodb.org/mongo-driver/bson/primitive" // Importing BSON primitive types for MongoDB
	"google.golang.org/grpc/codes"               // Importing gRPC error codes
	"google.golang.org/grpc/status"              // Importing gRPC status handling
	"log"
)

// CreateCartCast creates a new CartCast entry in the database.
// It receives a CartCast message and returns the ID of the created entry.
func (*Server) CreateCartCast(ctx context.Context, in *pb.CartCast) (*pb.CartCastId, error) {
	log.Printf("server/create.go::CreateCartCast()::CreateCartCast invoked %v\n", in)

	// Create an instance of CartCastItem from the input CartCast message
	data := CartCastItem{
		Title:    in.Title,
		XCord:    in.XCord,
		YCord:    in.YCord,
		ImageURL: in.ImageUrl,
		VideoURL: in.VideoUrl,
	}

	// Insert the new CartCastItem into the MongoDB collection
	res, err := collection.InsertOne(ctx, data)
	if err != nil {
		return nil, status.Errorf(
			codes.Internal,
			fmt.Sprintf("server/create.go::CreateCartCast()::Internal error: %v", err),
		)
	}

	// Extract the inserted ID from the result
	oid, ok := res.InsertedID.(primitive.ObjectID)
	if !ok {
		return nil, status.Errorf(
			codes.Internal,
			"server/create.go::CreateCartCast()::Cannot convert inserted ID to ObjectID", // Improved error message
		)
	}

	// Return the ID of the newly created CartCast entry
	return &pb.CartCastId{
		Id: oid.Hex(),
	}, nil
}
