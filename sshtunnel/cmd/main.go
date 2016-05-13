package main

import (
	"log"

	"github.com/QuentinPerez/go-stuff/sshtunnel/pkg/sshtunnel"
	"github.com/Sirupsen/logrus"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

func main() {
	tunnel := sshtunnel.NewTunnel("qperez")

	tunnel.Local = sshtunnel.HostPort{Host: "localhost", Port: "9999"}
	tunnel.Server = sshtunnel.HostPort{Host: "HOSTDB", Port: "PORTSSH(22)"}
	tunnel.Remote = sshtunnel.HostPort{Host: "HOSTDB", Port: "PORTDB"}

	errChan := tunnel.Start()
	select {
	case err := <-errChan:
		logrus.Fatal(err)
	default:
		break
	}
	_, err := gorm.Open("postgres", "host=localhost user=USER dbname=test sslmode=disable port=9999")
	if err != nil {
		log.Fatalf("%v", err)
	}
	select {}
}
