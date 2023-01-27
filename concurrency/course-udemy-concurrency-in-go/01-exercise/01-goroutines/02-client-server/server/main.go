package main

import (
	"io"
	"log"
	"net"
	"time"
)

func main() {
	listener, err := net.Listen("tcp", "localhost:8000")
	if err != nil {
		log.Fatal(err)
	}

	var counter int
	for {
		log.Println("waiting for new connection")
		conn, err := listener.Accept()

		if err != nil {
			log.Printf("\nerror in connection %v\n", err)
			continue
		}
		counter++
		log.Printf("connection number %d established\n", counter)

		go handleConn(conn)
	}
}

// handleConn - utility function
func handleConn(c net.Conn) {
	defer c.Close()
	for {
		_, err := io.WriteString(c, "response from server\n")
		if err != nil {
			return
		}
		time.Sleep(time.Second)
	}
}
