package main

import (
	"grpc-patterns/server"
)

func main() {
	server.ServeGrpc(server.New(), "localhost:6001")
}
