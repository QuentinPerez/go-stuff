package main

import (
	"fmt"
	"sync"
)

func gen(nb int) <-chan int {
	out := make(chan int)
	go func() {
		for i := 1; i <= nb; i++ {
			out <- i
		}
		close(out)
	}()
	return out
}

func square(input <-chan int) <-chan int {
	out := make(chan int)
	go func() {
		for v := range input {
			out <- v * v
		}
		close(out)
	}()
	return out
}

func merge(cs ...<-chan int) <-chan int {
	var wg sync.WaitGroup
	out := make(chan int)

	wg.Add(len(cs))
	for _, c := range cs {
		go func(c <-chan int) {
			for n := range c {
				out <- n
			}
			wg.Done()
		}(c)
	}
	go func() {
		wg.Wait()
		close(out)
	}()
	return out
}

func main() {
	input := gen(100)

	// Fan-out
	c1 := square(input)
	c2 := square(input)
	c3 := square(input)

	// Fan-in
	for n := range merge(c1, c2, c3) {
		fmt.Println(n)
	}
}
