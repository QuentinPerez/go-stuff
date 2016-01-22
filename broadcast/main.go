package main

import (
	"fmt"

	"github.com/QuentinPerez/go-stuff/broadcast/pkg/broadcast"
)

func main() {
	brd := broadcast.New()
	defer brd.Close()

	o1 := brd.SubscribeOuput()
	o2 := brd.SubscribeOuput()
	i := brd.SubscribeInput()

	id := 0
	done := 2
	for {
		select {
		case value := <-o1:
			_ = value
			done++
			fmt.Println("o1 <-", value)
		case value := <-o2:
			_ = value
			done++
			fmt.Println("o2 <-", value)
		default:
			if done == 2 {
				if id == 10 {
					return
				}
				i <- id
				id++
				done = 0
			}
		}
	}
}
