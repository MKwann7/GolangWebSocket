package main

import (
	"bufio"
	"fmt"
	"net"
)

func main() {
	// listen on port 8080
	ln, _ := net.Listen("tcp", ":8080")

	// accept connection
	conn, _ := ln.Accept()

	// run loop forever (or until ctrl-c)
	for {
		// process data
		// get message, output
		message, _ := bufio.NewReader(conn).ReadString('\n')
		fmt.Print("Message Received:", string(message))
	}
}
