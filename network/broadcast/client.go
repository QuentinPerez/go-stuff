package main

import (
	"fmt"
	"log"
	"net"
)

func TCPserver() {
	ln, err := net.ListenTCP("tcp", &net.TCPAddr{
		IP:   net.IPv4zero,
		Port: 1999,
	})
	if err != nil {
		log.Fatal(err)
	}

	conn, err := ln.Accept()
	if err != nil {
		log.Fatal(err)
	}
	buf := make([]byte, 16)

	for {
		conn.Write([]byte("SYNC"))
		n, err := conn.Read(buf)
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println(string(buf[:n]))
	}
}

func main() {
	conn, err := net.DialUDP("udp4", &net.UDPAddr{
		IP:   net.IPv4zero,
		Port: 4000,
	}, &net.UDPAddr{
		IP:   net.IPv4bcast,
		Port: 5000,
	})
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close()
	conn.Write([]byte("1999"))
	TCPserver()
}
