package main

import (
	"log"
	"net"
	"strconv"
	"time"
)

func main() {
	conn, err := net.ListenUDP("udp4", &net.UDPAddr{
		IP:   net.IPv4zero,
		Port: 5000,
	})
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close()
	buf := make([]byte, 1024)
	for {
		n, addr, err := conn.ReadFromUDP(buf)
		if err != nil {
			log.Fatal(err)
		}
		port, _ := strconv.Atoi(string(buf[:n]))
		go func(p int, a string) {
			tcpa, err := net.ResolveTCPAddr("tcp", a)
			if err != nil {
				log.Fatal(err)
			}

			c, err := net.DialTCP("tcp", nil, &net.TCPAddr{
				IP:   tcpa.IP,
				Port: p,
			})
			if err != nil {
				log.Fatal(err)
			}
			buf := make([]byte, 16)
			for {
				c.Read(buf)
				time.Sleep(time.Second)
				c.Write([]byte("ACK"))
			}
		}(port, addr.String())
		_, err = conn.WriteToUDP([]byte("ACK"), addr)
		if err != nil {
			log.Fatal(err)
		}
	}
}
