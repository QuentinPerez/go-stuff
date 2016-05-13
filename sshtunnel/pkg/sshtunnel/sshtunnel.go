package sshtunnel

import (
	"fmt"
	"io"
	"net"
	"os"

	"golang.org/x/crypto/ssh"
	"golang.org/x/crypto/ssh/agent"
)

type HostPort struct {
	Host string
	Port string
}

func (hp HostPort) String() string {
	return fmt.Sprintf("%v:%v", hp.Host, hp.Port)
}

type sshTunnel struct {
	Local   HostPort
	Remote  HostPort
	Server  HostPort
	config  ssh.ClientConfig
	errChan chan error
}

func NewTunnel(sshUser string) *sshTunnel {
	sshagent := func() ssh.AuthMethod {
		sshAgent, err := net.Dial("unix", os.Getenv("SSH_AUTH_SOCK"))
		if err != nil {
			return nil
		}
		return ssh.PublicKeysCallback(agent.NewClient(sshAgent).Signers)
	}

	return &sshTunnel{
		config: ssh.ClientConfig{
			User: sshUser,
			Auth: []ssh.AuthMethod{
				sshagent(),
			},
		},
		errChan: make(chan error, 10),
	}
}

func (s *sshTunnel) Start() <-chan error {
	ln, err := net.Listen("tcp", s.Local.String())
	if err != nil {
		s.errChan <- err
		return s.errChan
	}
	go func() {
		defer ln.Close()
		for {
			conn, errConn := ln.Accept()
			if err != nil {
				s.errChan <- errConn
			}
			go s.forward(conn)
		}
	}()
	return s.errChan
}

func (s *sshTunnel) forward(local net.Conn) {
	defer local.Close()

	server, err := ssh.Dial("tcp", s.Server.String(), &s.config)
	if err != nil {
		s.errChan <- err
		return
	}
	defer server.Close()

	remote, err := server.Dial("tcp", s.Remote.String())
	if err != nil {
		s.errChan <- err
		return
	}
	defer remote.Close()

	copy := func(writer, reader net.Conn) {
		_, err := io.Copy(writer, reader)
		if err != nil {
			s.errChan <- err
		}
	}
	go copy(local, remote)
	copy(remote, local)
}
