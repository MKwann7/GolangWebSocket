package main

import (
	"bufio"
	"fmt"
	"net"
	"strconv"
)

func main() {

	// configuration
	network := "tcp"
	port := 8080

	// listen on port 8080
	ln, _ := net.Listen(network, ":"+strconv.Itoa(port))

	// accept connection
	conn, _ := ln.Accept()

	// run loop forever (or until ctrl-c)
	for {
		message, _ := bufio.NewReader(conn).ReadString('\n')

		fmt.Print("Message Received:", string(message))
	}
}
