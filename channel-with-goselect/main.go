package main

import "github.com/QuentinPerez/go-stuff/channel-with-goselect/goselect"

func main() {
	// no limit
	goselect.Start(10000)
}
