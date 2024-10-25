// cartcast/server/read.go
// This file contains the implementation of the ReadCartCast method,
// which retrieves a CartCast entry from the MongoDB database based on its ID.

package main

import (
	"context"
	pb "ekeberg.com/go-grpc-cartcast/cartcast/proto" // Importing the CartCast protocol buffer definitions
	"go.mongodb.org/mongo-driver/bson"               // Importing BSON types for MongoDB queries
	"go.mongodb.org/mongo-driver/bson/primitive"     // Importing BSON primitive types for ObjectID handling
	"google.golang.org/grpc/codes"                   // Importing gRPC error codes
	"google.golang.org/grpc/status"                  // Importing gRPC status handling
	"log"
)

// ReadCartCast retrieves a CartCast entry from the database using the provided ID.
// It receives a CartCastId message and returns the corresponding CartCast entry or an error.
func (s *Server) ReadCartCast(ctx context.Context, in *pb.CartCastId) (*pb.CartCast, error) {
	log.Printf("server/read.go::ReadCartCast()::Invoked id: %v\n", in)

	// Convert the ID from the CartCastId message to a primitive.ObjectID.
	oid, err := primitive.ObjectIDFromHex(in.Id)
	if err != nil {
		return nil, status.Errorf(
			codes.InvalidArgument,
			"server/read.go::ReadCartCast()::Cannot parse ID", // Return error if the ID is invalid
		)
	}

	// Prepare to read the CartCastItem from the database.
	data := &CartCastItem{}      // Create an empty CartCastItem to hold the result
	filter := bson.M{"_id": oid} // Create a filter to search for the item by ID

	// Execute the MongoDB query to find the CartCastItem
	res := collection.FindOne(ctx, filter)

	// Decode the result into the data variable, handling the case where no document is found.
	if err := res.Decode(data); err != nil {
		return nil, status.Errorf(
			codes.NotFound,
			"server/read.go::ReadCartCast()::CartCast not found", // Return error if the item is not found
		)
	}

	// Convert the CartCastItem to a CartCast protocol buffer message and return it
	return documentToCartCast(data), nil
}
