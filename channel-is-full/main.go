package main

import "fmt"

// How to know is a channel is full
func main() {
	c := make(chan bool, 1)

	c <- true

	// 1) using select method
	select {
	case c <- false:
	default:
		fmt.Println("Channel is full")
	}

	// 2) check the length and the capacity
	if len(c) == cap(c) {
		fmt.Println("Channel is full")
	}

	// get an element to flush the channel
	_ = <-c

	select {
	case c <- false:
		_ = <-c
		fmt.Println("Channel is not full")
	default:
		fmt.Println("Channel is full")
	}

	if len(c) == cap(c) {
		fmt.Println("Channel is full")
	} else {
		fmt.Println("Channel is not full")
	}
}
