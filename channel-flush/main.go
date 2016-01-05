package main

import (
	"fmt"
	"runtime"
)

func blockAndFlushBruteForce() {
	buffered := make(chan bool, 2)
	done := make(chan struct{})

	go func() {
		for {
			for len(buffered) > 0 {
				<-buffered
				fmt.Printf("cap: %d, len: %d\n", cap(buffered), len(buffered))
			}
			runtime.Gosched()
			if len(buffered) == 0 {
				close(done)
				fmt.Println("Flushed")
				break
			}
		}
	}()
	buffered <- true
	buffered <- true
	buffered <- true
	buffered <- true
	buffered <- true
	buffered <- true
	buffered <- true
	buffered <- true
	buffered <- true
	buffered <- true
	buffered <- true
	buffered <- true
	buffered <- true
	buffered <- true
	buffered <- true
	buffered <- true
	buffered <- true
	buffered <- true
	buffered <- true
	buffered <- true
	buffered <- true
	buffered <- true
	buffered <- true
	buffered <- true
	buffered <- true
	<-done
}

func blockAndFlush() {
	buffered := make(chan bool, 5)

	go func() {
		for {
			if len(buffered) == 5 {
				fmt.Println("Full")
				break
			}
		}
		for len(buffered) > 0 {
			<-buffered
			fmt.Printf("cap: %d, len: %d\n", cap(buffered), len(buffered))
		}
	}()
	buffered <- true
	buffered <- true
	buffered <- true
	buffered <- true
	buffered <- true
	buffered <- true
	fmt.Println("Unblock")
}

func basicFlush() {
	buffered := make(chan bool, 5)

	buffered <- true
	buffered <- true
	buffered <- true
	buffered <- true
	for len(buffered) > 0 {
		<-buffered
		fmt.Printf("cap: %d, len: %d\n", cap(buffered), len(buffered))
	}
}

func main() {
	fmt.Println("=== basicFlush ===")
	basicFlush()
	fmt.Println("\n=== blockAndFlush ===")
	blockAndFlush()
	fmt.Println("\n=== blockAndFlushBruteForce ===")
	blockAndFlushBruteForce()
}
