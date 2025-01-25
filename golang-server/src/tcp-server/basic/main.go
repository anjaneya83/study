// To create a tcp server
// Import the net package

package main

import (
	"fmt"
	"net"
)

func main() {
	//Listen for incoming connections on port 8080
	fmt.Println("This is a TCP server program")
	ln, err := net.Listen("tcp", "127.0.0.1:12001")
	if err != nil {
		fmt.Println(err)
		return
	}
	defer ln.Close()
	fmt.Println("Server is now listening on port 12001")
	//now we listen for incoming connection and spawn a new goroutine to handle per connection request
	for {
		conn, err := ln.Accept()
		if err != nil {
			fmt.Println(err)
			return
		}

		fmt.Println("Server has been approached by a client")
		go handleConection(conn)

	}

}

func handleConection(conn net.Conn) {
	//close the connection when done
	defer conn.Close()
	// Read incoming data
	buf := make([]byte, 1024) //create a 1kb byte array
	fmt.Println("going to read data from client")
	no_of_bytes, err := conn.Read(buf)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Printf("Received: %s", buf[:no_of_bytes])
	_, err = conn.Write([]byte("Hello my Friend ,The Client, From Server"))
	if err != nil {
		fmt.Println(err)
		return
	}
}
