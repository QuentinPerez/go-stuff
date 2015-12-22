package main

import (
	"fmt"
	"time"
)

type Reader struct {
	output chan interface{}
}

func NewReader() *Reader {
	return &Reader{
		output: make(chan interface{}),
	}
}

func (r *Reader) Output(...<-chan interface{}) <-chan interface{} {
	go func() {
		defer fmt.Println("Read.Output: Exit")
		defer close(r.output)
		i := 0

		for {
			r.output <- "Msg From Reader"
			time.Sleep(1 * time.Second)
			i++
			if i == 2 {
				return
			}
		}
	}()
	return r.output
}
