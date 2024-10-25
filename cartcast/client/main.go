// cartcast/client/main.go
// This file implements the client for the CartCast gRPC service.
// It establishes a connection to the gRPC server and invokes the CreateCartCast method.

package main

import (
	"log"

	pb "ekeberg.com/go-grpc-cartcast/cartcast/proto" // Importing the CartCast protocol buffer definitions
	"google.golang.org/grpc"                         // Importing gRPC package for client connection
	"google.golang.org/grpc/credentials/insecure"    // Importing insecure credentials for gRPC connection
)

// addr specifies the address of the gRPC server to connect to
var addr string = "0.0.0.0:50051"

func main() {
	// Establish a connection to the gRPC server using insecure credentials
	conn, err := grpc.Dial(addr, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		// Log and terminate the program if the connection fails
		log.Fatalf("client/main.go::main()::Couldn't connect to client: %v\n", err)
	}

	defer conn.Close() // Ensure the connection is closed when the function exits

	// Create a new CartCast service client from the established connection
	c := pb.NewCartCastServiceClient(conn)

	// Call the function to create a new CartCast entry (implementation not shown in this snippet)
	id := createCartCast(c)
	readCartCast(c, id) // Valid
	// readCartCast(c, "invalid id")
	UpdateCartCast(c, id)
	ListCartCast(c)
	DeleteCartCast(c, id)
	readByCoordsCartCast(c, "90", "20") // Valid
	// readByCoordsCartCast(c, "23", "23") // invalid
}
