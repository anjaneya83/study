// This is a TCP client program
package main

import (
	"fmt"
	"net"
)

func main() {
	//connect to a server
	conn, err := net.Dial("tcp", "127.0.0.1:12001")
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("Connection established with the server")
	fmt.Println("Say hello")
	_, err = conn.Write([]byte("Hello my Friend ,The Server, From the Client"))
	if err != nil {
		fmt.Println(err)
		return
	}
	buf := make([]byte, 1024) //create a 1kb byte array
	fmt.Println("going to receive response from server")
	_, err = conn.Read(buf)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Printf("Response Received: %s", buf)

}
