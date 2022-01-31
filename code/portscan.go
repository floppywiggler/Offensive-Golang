// A simple portscanner implemented in Golang!

package main

import (
	"fmt"
	"sort"
	"net"
)

func worker(ports chan int, results chan int){ //Creating a function for the threads with ports and results parameters
	for p := range ports {
		address := fmt.Sprintf("scanme.nmap.org:%d", p)
		fmt.Printf("Scanning port %d", p)
		conn, err := net.Dial("tcp", address)
		if err != nil {
			results <- 0
			continue
		}
		conn.Close()
		results <- p
	}
}

func main() {
	ports := make (chan int, 100) // Define data type for ports
	results := make(chan int) // Define data type for results
	var openports []int // Variable table for open ports

	for i := 0; i < cap(ports); i++ { // For each ports in list
		go worker (ports, results)

	}

	go func(){

		for i := 1; i <= 1024; i++ {
		ports <- i
		}

	}()
	for i := 0; i < 1024; i++ {
		port := <-results
		if port != 0{ // If port is not 0,
			openports = append(openports, port) // Add it to openports
		}
	}
	close(ports)
	close(results)
	sort.Ints(openports)
	for _, port := range openports {
		fmt.Printf("%d open\n", port)
	}
}
