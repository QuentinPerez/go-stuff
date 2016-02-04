package main

import "github.com/QuentinPerez/go-stuff/channel-with-reflectSelect/reflectSelect"

//
// go version go1.5.2 darwin/amd64
// 4 cores
// 32 Go RAM
//

func main() {
	reflectSelect.Start(5000)
}
