package server

import (
	"fmt"
	"net"
	"os"
	"os/signal"

	"github.com/sirupsen/logrus"
	"google.golang.org/grpc"
)

func ServeGrpc(server *grpc.Server, grpcAddress string) {
	listener, err := net.Listen("tcp", grpcAddress)
	if err != nil {
		logrus.Fatalf("GrpcServer failed to listen %v", err)
	}

	logrus.Infof("Preparing grpc server on %v...\n", grpcAddress)

	go func() {
		logrus.Infof("Starting grpc server on %v\n", grpcAddress)
		if err := server.Serve(listener); err != nil {
			logrus.Fatalf("Failed to serve: %v", err)
		}
	}()

	// wait for ctrl c to exit
	ch := make(chan os.Signal, 1)
	signal.Notify(ch, os.Interrupt)

	// Block until signal is received
	<-ch
	fmt.Println("Stopping the grpc server")
	server.Stop()
	fmt.Println("Closing the listener")
	_ = listener.Close()
	fmt.Println("End of program")
}
