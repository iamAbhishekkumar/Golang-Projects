package main

import (
	"fmt"
	"net"
)

const bufferSize int = 1024

func main() {
	ln, err := net.ListenUDP("udp", &net.UDPAddr{
		IP:   net.ParseIP("127.0.0.1"),
		Port: 8000,
	})

	if err != nil {
		fmt.Println(err.Error())
		panic("Unable to Listen")
	}

	for {
		var buf [bufferSize]byte
		_, addr, err := ln.ReadFromUDP(buf[:])
		if err != nil {
			fmt.Println(err.Error())
			return
		}

		fmt.Println("Message Read From : ", addr.IP)
		fmt.Println("Message : ", string(buf[:]))

		ln.WriteToUDP([]byte("Message Received !!\n"), addr)
	}
}
