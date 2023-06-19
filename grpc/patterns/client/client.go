package main

import (
	"context"
	"grpc-patterns/proto"

	"github.com/sirupsen/logrus"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
)

func main() {
	logrus.Info("Connecting to grpc server...")
	adderess := "localhost:6001" // os.Getenv("GRPC_ADDRESS")
	local := true                // os.Getenv("LOCAL")

	var opts []grpc.DialOption
	if local {
		opts = append(opts, grpc.WithInsecure())
	} else {
		cred := credentials.NewTLS(nil)
		opts = append(opts, grpc.WithTransportCredentials(cred))
	}
	conn, err := grpc.Dial(adderess, opts...)
	if err != nil {
		logrus.Fatalf("Error connecting to grpc server %v: \n%v\n", adderess, err)
	}
	defer conn.Close()

	client := proto.NewServiceClient(conn)
	logrus.Info("Connected...")

	Test(client)
}

func Test(client proto.ServiceClient) {
	logrus.Info("Calling service...")
	empty := ""
	req := &proto.OptionalRequest{
		ReqValue:   "req value",
		OptValue_1: &empty,
	}
	resp, err := client.UpdateOptional(context.Background(), req)
	if err != nil {
		logrus.Fatalf("Error getting response: %v", err)
	}
	logrus.Infof("Response: %v", resp)
}
