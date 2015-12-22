package main

import (
	"fmt"
	"time"
)

type Writer struct {
	output chan interface{}
}

func NewWriter() *Writer {
	return &Writer{
		output: make(chan interface{}),
	}
}

func (r *Writer) Output(...<-chan interface{}) <-chan interface{} {
	go func() {
		defer fmt.Println("Writer.Output: Exit")
		defer close(r.output)

		r.output <- "Msg From Writer"
		time.Sleep(1 * time.Second)
	}()
	return r.output
}
