package main

import (
	"log"
	"net"
	"io"
)

// Echo is a handler function that simply echoes received data.
func echo(conn net.Conn) {
	defer conn.Close()



	//Create buffer to store received data
	b := make([]byte, 512)
	for {
		//REceive data via conn.Read into a buffer
		size, err := conn.Read(b[0:])
		if err == io.EOF {
			log.Println("Client Disconnected")
			break
		}
		if err != nil {

			log.Println("Unexpected error")
			//Send data via conn.Write
			log.Println("Writing data")
			if _,err := conn.Write(b[0:size]); err != nil {
				log.Fatalln("Unable to write data")
			}
		}
	}
}
func main() {
	//Bind to TCP Port 1337 on all interfaces
	listener, err := net.Listen("tcp", ":1337")
	if err != nil {
		log.Fatalln("Unable to bind to port")
	}
	log.Println("Listening on 0.0.0.0:1337")
	for {
		//Wait for connection. Create .net.Conn on conneciton established
		conn, err := listener.Accept()
		log.Println("Received connection")
		if err != nil {
			log.Fatalln("Unable to accept connection")
		}
		// Handle the connection. Using goroutine for concurrency
		go echo(conn)
	}
}
