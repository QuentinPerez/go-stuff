package main

import (
	"log"
	"net"
	"net/http"
	"net/rpc"

	"github.com/QuentinPerez/go-stuff/rpc/rpcdefinition"
)

func main() {
	call := new(rpcdefinition.Call)
	err := rpc.Register(call)
	if err != nil {
		log.Fatalf("Format of service call isn't correct. %v", err)
	}
	rpc.HandleHTTP()
	//start listening for messages on port 1234
	l, e := net.Listen("tcp", ":1234")
	if e != nil {
		log.Fatalf("Couldn't start listening on port 1234. %v", e)
	}
	log.Println("Serving RPC handler")
	err = http.Serve(l, nil)
	if err != nil {
		log.Fatalf("Error serving: %v", err)
	}
}
