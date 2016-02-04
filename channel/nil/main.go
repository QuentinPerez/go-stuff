package main

import (
	"fmt"
	"time"
)

func main() {
	var running chan int
	var fetch <-chan time.Time
	deadline := time.After(15 * time.Second)
	id := 0

	running = nil
	for {
		if running == nil {
			fetch = time.After(0 * time.Second)
		}
		select {
		case data := <-running:
			running = nil
			fmt.Println("Data", data)
		case <-deadline:
			// check running
			if running != nil {
				fmt.Println("Wait running")
				<-running
			}
			return
		case <-fetch:
			running = make(chan int, 1)
			go func() {
				<-time.After(3 * time.Second)
				id++
				running <- id
			}()
		}
	}
}
