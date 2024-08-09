package main

import (
	"fmt"
	"net"
)

const bufferSize int = 1024

func main() {
	ln, err := net.Listen("tcp", ":8000")
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	for {
		conn, err := ln.Accept()
		if err != nil {
			fmt.Println(err.Error())
			continue
		}

		go handleConnection(conn)
	}
}

func handleConnection(conn net.Conn) {
	defer conn.Close()

	buf := make([]byte, bufferSize)
	_, err := conn.Read(buf)
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	fmt.Printf("Recevied: %s ", buf)
}
