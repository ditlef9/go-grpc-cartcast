// cartcast/server/read_by_coords.go
// This file contains the implementation of the ReadCartCast method,
// which retrieves a CartCast entry from the MongoDB database based on its ID.

package main

import (
	"context"
	pb "ekeberg.com/go-grpc-cartcast/cartcast/proto" // Importing the CartCast protocol buffer definitions
	"go.mongodb.org/mongo-driver/bson"               // Importing BSON types for MongoDB queries
	"google.golang.org/grpc/codes"                   // Importing gRPC error codes
	"google.golang.org/grpc/status"                  // Importing gRPC status handling
	"log"
)

// ReadCartCast retrieves a CartCast entry from the database using the provided x and y coordinates
func (s *Server) ReadByCoordsCartCast(ctx context.Context, in *pb.CartCastCoords) (*pb.CartCast, error) {
	log.Printf("server/read_by_coords.go::ReadByCoordsCartCast()::Invoked coordinates: %v\n", in)

	// Update filter to use coordinates instead of ObjectID
	filter := bson.M{"x_cord": in.XCord, "y_cord": in.YCord}

	// Prepare an empty CartCastItem to store the query result
	data := &CartCastItem{}

	// Execute the MongoDB query to find the CartCastItem by coordinates
	res := collection.FindOne(ctx, filter)

	// Decode the result into the data variable, handling the case where no document is found.
	if err := res.Decode(data); err != nil {
		return nil, status.Errorf(
			codes.NotFound,
			"server/read_by_coords.go::ReadByCoordsCartCast()::CartCast not found", // Return error if the item is not found
		)
	}

	// Convert the CartCastItem to a CartCast protocol buffer message and return it
	return documentToCartCast(data), nil
}
