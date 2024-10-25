// cartcast/server/main.go
// This file initializes the gRPC server and connects to the MongoDB database for the CartCast service.

package main

import (
	"context"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"log"
	"net"

	pb "ekeberg.com/go-grpc-cartcast/cartcast/proto"
	"google.golang.org/grpc"
)

// collection holds the MongoDB collection for cart items
var collection *mongo.Collection

// addr specifies the address the gRPC server will listen on
var addr string = "0.0.0.0:50051"

// Server struct implements the CartCast service defined in the protocol buffer
type Server struct {
	pb.CartCastServiceServer // Embedding the CartCast service interface
}

// main function initializes the MongoDB connection and starts the gRPC server
func main() {
	// Create a new MongoDB client
	client, err := mongo.NewClient(options.Client().ApplyURI("mongodb://root:root@localhost:27017/"))
	if err != nil {
		log.Fatal(err)
	}

	// Connect to the MongoDB database
	err = client.Connect(context.Background())
	if err != nil {
		log.Fatal(err)
	}
	// Assign the collection to the variable after a successful connection
	collection = client.Database("cartcastdb").Collection("cartcast")

	// Start listening on the specified address
	lis, err := net.Listen("tcp", addr)
	if err != nil {
		log.Fatalf("server/main.go::main()::Failed to listen: %v\n", err)
	}
	log.Printf("server/main.go::main()::Listening at %s\n", addr)

	// Create a new gRPC server
	s := grpc.NewServer()
	// Register the CartCast service with the gRPC server
	pb.RegisterCartCastServiceServer(s, &Server{})

	// Start serving incoming requests
	if err := s.Serve(lis); err != nil {
		log.Fatalf("server/main.go::main()::Failed to serve: %v\n", err)
	}
}
