package main

import(
	"io"
	"net"
	"log"
)

func handle(src net.Conn){
	dst, err := net.Dial("tcp", "google.com:80")
	if err != nil {
		log.Fatalln("Unable to connect to our unreachable host")
	}
defer dst.Close()

// Run in goroutine to prevent io.Copy from blocking
go func(){
	// Copy our source's output to the destination
	if _, err := io.Copy(dst,src); err != nil {
		log.Fatalln(err)
	}
}()
// Copy our destination's output back to our source
if _, err := io.Copy(src, dst); err != nil{
	log.Fatalln(err)
	}
}
func main(){
	// Listen on local port 80
	listener, err := net.Listen("tcp", ":8000")
	if err != nil {
		log.Fatalln("Unable to bind to port")
	}
	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Fatalln("Unable to accept connection")
		}
		go handle(conn)
	}
}
