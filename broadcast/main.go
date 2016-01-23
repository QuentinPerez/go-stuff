package main

import (
	"fmt"
	"time"

	"github.com/QuentinPerez/go-stuff/broadcast/pkg/broadcast"
)

func main() {
	brd := broadcast.New()
	closed := make(chan struct{})
	defer func() {
		for i := 0; i < 10; i++ {
			<-closed
		}
	}()
	defer brd.Close()

	for i := 0; i < 10; i++ {
		go func() {
			id, output := brd.SubscribeOuput()

			for {
				value, ok := <-output
				if !ok {
					fmt.Println(id, "closed")
					closed <- struct{}{}
					return
				}
				fmt.Println(id, "-->", value)
			}
		}()
	}
	input := brd.SubscribeInput()
	input2 := brd.SubscribeInput()
	input3 := brd.SubscribeInput()

	for brd.NbOutputChannel() != 10 {
		time.Sleep(10 * time.Millisecond)
	}
	input <- "message 1 from input 1"
	input <- "message 99 from input 1"
	input2 <- "message 21 from input 2"
	input2 <- "message 299 from input 2"
	input3 <- "message 31 from input 3"
	input3 <- "message 399 from input 3"
}
