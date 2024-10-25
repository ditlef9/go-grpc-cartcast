// cartcast/server/cartcast_item.go
// This package provides functionality to manage cart items for the CartCast service,
// including the definition of the CartCastItem structure and conversion functions.

package main

import (
	pb "ekeberg.com/go-grpc-cartcast/cartcast/proto" // Importing the CartCast protocol buffer definitions
	"go.mongodb.org/mongo-driver/bson/primitive"     // Importing BSON primitive types for MongoDB
)

// CartCastItem represents an item in the CartCast service.
// It includes fields for item identification and multimedia content.
type CartCastItem struct {
	ID       primitive.ObjectID `bson:"_id,omitempty"` // Unique identifier for the item, generated by MongoDB
	Title    string             `bson:"title"`
	XCord    string             `bson:"x_cord"`
	YCord    string             `bson:"y_cord"`
	ImageURL string             `bson:"image_url"`
	VideoURL string             `bson:"video_url"`
}

// documentToCartCast converts a CartCastItem struct to a CartCast protocol buffer message.
// This is used for sending CartCast items over gRPC.
// Fields exists in cartcast/proto/cartcast.pb.go
func documentToCartCast(data *CartCastItem) *pb.CartCast {
	return &pb.CartCast{
		Id:       data.ID.Hex(),
		Title:    data.Title,
		XCord:    data.XCord,
		YCord:    data.YCord,
		ImageUrl: data.ImageURL,
		VideoUrl: data.VideoURL,
	}
}