package main

import (
	"fmt"
	"net"
)

func main() {
	conn, err := net.Dial("tcp", ":8000")
	if err != nil {
		fmt.Println(err)
		return
	}
	_, err = conn.Write([]byte("Hello Server!"))
	if err != nil {
		fmt.Println(err)
		return
	}

	defer conn.Close()
}
