package main

import (
	"net"
	"time"

	"github.com/QuentinPerez/go-stuff/auto-reconnect/autosock"
	"github.com/Sirupsen/logrus"
)

func main() {
	ready := make(chan struct{})
	// setup server
	go func() {
		ln, err := net.Listen("tcp", ":4242")
		if err != nil {
			logrus.Fatal("main.listen:", err)
		}
		ready <- struct{}{}
		defer ln.Close()
		cl, err := ln.Accept()
		if err != nil {
			logrus.Fatal("main.Accept:", err)
		}
		defer cl.Close()
	}()
	<-ready
	auto := autosock.New(func() (net.Conn, error) {
		co, err := net.Dial("tcp", ":4242")
		if err != nil {
			return nil, err
		}
		return co, nil
	})
	time.Sleep(1 * time.Second)
	auto.Close()
}
