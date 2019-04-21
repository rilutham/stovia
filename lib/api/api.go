package api

import (
	"log"

	pb "rilutham/stovia/lib/rpc"

	"golang.org/x/net/context"
)

// API :nodoc:
type API struct{}

func NewAPI() *API {
	return &API{}
}

func (a *API) GetFutureValue(ctx context.Context, in *pb.FutureValueRequest) (*pb.FutureValueResponse, error) {
	log.Printf("Received: %v", in.Code)
	return &pb.FutureValueResponse{
		FutureValue:      100,
		RecommendedToBuy: true,
	}, nil
}
