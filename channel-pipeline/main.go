package main

import "fmt"

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

func main() {
	for v := range square(gen(10)) {
		fmt.Println(v)
	}
}
