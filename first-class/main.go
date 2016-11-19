package main

import (
	"fmt"
	"io"
	"log"
	"net"
	"strconv"
	"strings"
)

type Server struct {
	ops chan func(map[int]net.Conn)
}

func NewServer() *Server {
	return &Server{
		ops: make(chan func(map[int]net.Conn), 5),
	}
}

func (s *Server) HandleDisconnect(why error, key int) {
	s.ops <- func(m map[int]net.Conn) {
		m[key].Close()
		delete(m, key)
		s.BroacastMsg(fmt.Sprintf("Client %d disconneted: %v\n", key, why))
	}
}

func (s *Server) ListUser(key int) {
	s.ops <- func(m map[int]net.Conn) {
		for k := range m {
			if k == key {
				continue
			}
			io.WriteString(m[key], fmt.Sprintf("clientID: %d", k))
		}
	}
}

func (s *Server) SendPrivateMSG(mykey, keypmgs int, msg string) {
	s.ops <- func(m map[int]net.Conn) {
		if peer, ok := m[keypmgs]; ok {
			io.WriteString(peer, msg)
		} else {
			io.WriteString(m[mykey], fmt.Sprintf("User %d not found", keypmgs))
		}
	}
}

func (s *Server) HandleCommands(conn net.Conn, key int, cmd string) error {
	tab := strings.Split(cmd, " ")
	switch tab[0] {
	case "bdt":
		if len(tab) < 2 {
			io.WriteString(conn, "command format error\n")
			return nil
		}
		msg := cmd[len(tab[0]):]
		s.BroacastMsg(msg)
	case "list":
		s.ListUser(key)
	case "pmsg":
		if len(tab) < 3 {
			io.WriteString(conn, "command format error\n")
			return nil
		}
		keypeer, err := strconv.Atoi(tab[1])
		if err != nil {
			io.WriteString(conn, err.Error())
			return nil
		}
		msg := strings.Join(tab[2:], " ")
		s.SendPrivateMSG(key, keypeer, msg)
	case "quit":
		return fmt.Errorf("Goodbye")
	default:
		io.WriteString(conn, fmt.Sprintf("Invalid command [%v]\n", cmd))
	}
	return nil
}

func (s *Server) HandleClient(conn net.Conn, key int) {
	buf := make([]byte, 1024)

	for {
		n, err := conn.Read(buf)
		if err != nil {
			s.HandleDisconnect(err, key)
			return
		}
		if n > 1 {
			err = s.HandleCommands(conn, key, string(buf[:n-1]))
		}
		if err != nil {
			s.HandleDisconnect(err, key)
			return
		}
		io.WriteString(conn, "#> ")
	}
}

func (s *Server) BroacastMsg(msg string) {
	s.ops <- func(m map[int]net.Conn) {
		for _, cl := range m {
			_, _ = io.WriteString(cl, msg)
		}
	}
}

func (s *Server) HandleConnections(ln net.Listener) {
	go func() {
		for {
			conn, _ := ln.Accept()
			s.ops <- func(m map[int]net.Conn) {
				fmt.Println("new connection", len(m))
				key := len(m)
				m[key] = conn
				_, _ = io.WriteString(conn, `Welcome
Commands available:
bdt  <msg>           # broadcast a message
list                 # lists the users
pmsg <id> <msg>      # send msg to an user
quit                 # just quit
#> `)
				go s.HandleClient(conn, key)
			}
		}
	}()
}

func (s *Server) Loop() {
	conns := make(map[int]net.Conn)
	for op := range s.ops {
		op(conns)
	}
}

func main() {
	serv := NewServer()
	ln, err := net.Listen("tcp", "localhost:4242")
	if err != nil {
		log.Fatal(err)
	}
	serv.HandleConnections(ln)
	serv.Loop()
}
