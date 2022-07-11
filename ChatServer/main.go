package main

import (
	"log"
	"net"
)

func main() {
	Server := NewServer()
	go Server.Run()

	Listener, err := net.Listen("tcp", ":8888")
	if err != nil {
		log.Fatalf("Unable to Start Server: %s", err.Error())
	}

	defer Listener.Close()
	log.Printf("Started Server On :8888")

	for {
		Conn, Err := Listener.Accept()
		if Err != nil {
			log.Printf("Unable to Accept Connection: %s", Err.Error())
			continue
		}

		go Server.NewClient(Conn)
	}
}
