// cartcast/client/update

package main

import (
	"context"
	pb "ekeberg.com/go-grpc-cartcast/cartcast/proto"
	"log"
)

func UpdateCartCast(c pb.CartCastServiceClient, id string) {
	log.Println("client/update.go::UpdateCartCast()::Invoked")

	newCartCast := &pb.CartCast{
		Id:       id,
		Title:    "Omo Color kun 39,90 hele uken!",
		XCord:    "90",
		YCord:    "20",
		ImageUrl: "https://omo.no/wp-content/uploads/2022/05/Omo-Expert-Color-Pulver-16vask-768x768.png.webp",
		VideoUrl: "https://www.youtube.com/watch?v=-ufMzy_6VMA",
	}

	_, err := c.UpdateCartCast(context.Background(), newCartCast)

	if err != nil {
		log.Fatalf("client/update.go::UpdateCartCast()::Error while updating: %v\n", err)
	}

	log.Println("client/update.go::UpdateCartCast()::Update successfully")

}
