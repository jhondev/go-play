package server

import (
	"context"
	"fmt"
	"grpc-patterns/proto"

	"github.com/chidiwilliams/flatbson"
	"google.golang.org/grpc"
)

type GrpcService struct{}

func New() *grpc.Server {
	svr := grpc.NewServer()
	proto.RegisterServiceServer(svr, &GrpcService{})
	return svr
}

type DocumentUpdate struct {
	Title                  *string   `bson:"title,omitempty"`
	Description            *string   `bson:"description,omitempty"`
	SubType                *string   `bson:"subType,omitempty"`
	SubSubType             *string   `bson:"subSubType,omitempty"`
	Role                   *string   `bson:"role,omitempty"`
	OwnerConditionRequired *bool     `bson:"ownerConditionRequired,omitempty"`
	AccessLevel            *string   `bson:"accessLevel,omitempty"`
	ClassificationIDs      *[]string `bson:"classificationIds,omitempty"`
}

// UpdateOptional implements proto.ServiceServer.
func (*GrpcService) UpdateOptional(ctx context.Context, req *proto.OptionalRequest) (*proto.OptionalResponse, error) {
	fmt.Println("\n\nREQUEST VALUES")
	fmt.Print(req)

	fmt.Println("\n\nREQUEST BSON")
	doc := &DocumentUpdate{
		Title:       req.OptValue_1,
		Description: req.OptValue_1,
		SubSubType:  req.OptValue_2,
	}
	if req.ClassificationIds != nil {
		doc.ClassificationIDs = &req.ClassificationIds.Values
	}
	fmt.Print(flatbson.Flatten(doc))

	return &proto.OptionalResponse{}, nil
}
