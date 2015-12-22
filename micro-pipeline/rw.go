package main

import "fmt"

type RW struct {
	output chan interface{}
}

func NewRW() *RW {
	return &RW{
		output: make(chan interface{}),
	}
}

func (r *RW) Output(channels ...<-chan interface{}) <-chan interface{} {
	for id, channel := range channels {
		go func(ch <-chan interface{}, id int) {
			for {
				d, ok := <-ch
				if ok {
					fmt.Println("RW.Input from", id, ":", d)
					r.output <- d
				} else {
					fmt.Println("Reconnect", id)
					if id == 0 {
						ch = NewReader().Output(nil)
					} else if id == 1 {
						ch = NewWriter().Output(nil)
					}
				}
			}
		}(channel, id)
	}
	return r.output
}
