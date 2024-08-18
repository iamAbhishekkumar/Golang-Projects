package main

import (
	"bufio"
	"fmt"
	"net"
)

func main() {
	ln, err := net.DialUDP("udp", nil, &net.UDPAddr{
		IP:   net.ParseIP("127.0.0.1:8000"),
		Port: 8000,
	})

	if err != nil {
		fmt.Println(err.Error())
		return
	}

	linesWritten, err := ln.Write([]byte("Hello My name is Abhishek!!!"))
	fmt.Println("No of lines written : ", linesWritten)

	if err != nil {
		fmt.Println(err.Error())
		return
	}

	data, err := bufio.NewReader(ln).ReadString('\n')
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	fmt.Println("Message Recived From Server : ", data)

}
