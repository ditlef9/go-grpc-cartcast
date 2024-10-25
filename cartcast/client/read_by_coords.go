// cartcast/client/read_by_coords.go
// Read a CartCast

package main

import (
	"context"
	pb "ekeberg.com/go-grpc-cartcast/cartcast/proto"
	"log"
)

func readByCoordsCartCast(c pb.CartCastServiceClient, x_cord string, y_cord string) *pb.CartCast {
	log.Println("client/read_by_coords.go::readByCoordsCartCast()::Invoked -----------------------")

	req := &pb.CartCastCoords{XCord: x_cord, YCord: y_cord}
	res, err := c.ReadByCoordsCartCast(context.Background(), req)

	if err != nil {
		log.Printf("client/read_by_coords.go::readByCoordsCartCast()::Error while reading: %v\n", err)
	}

	if res != nil {
		log.Printf("client/read_by_coords.go::readByCoordsCartCast()::CartCast read: %v\n", res)
	} else {
		log.Println("client/read_by_coords.go::readByCoordsCartCast()::Received nil response")
	}

	return res
}
