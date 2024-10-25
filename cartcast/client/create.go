// cartcast/client/create.go

package main

import (
	"context"
	pb "ekeberg.com/go-grpc-cartcast/cartcast/proto"
	"log"
)

func createCartCast(c pb.CartCastServiceClient) string {
	log.Println("client/create.go::createCartCast()::Invoked ---------------------")

	cartCast := &pb.CartCast{
		Title:    "Omo Color kun 39,90 i helgen!",
		XCord:    "90",
		YCord:    "20",
		ImageUrl: "https://omo.no/wp-content/uploads/2022/05/Omo-Expert-Color-Pulver-16vask-768x768.png.webp",
		VideoUrl: "https://www.youtube.com/watch?v=-ufMzy_6VMA",
	}

	res, err := c.CreateCartCast(context.Background(), cartCast)
	if err != nil {
		log.Fatalf("client/create.go::createCartCast()::Unexpected error: %v\n", err)
	}

	log.Printf("client/create.go::createCartCast()::CartCast created: %s\n", res.Id)
	return res.Id

}
