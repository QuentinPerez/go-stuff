package main

import (
	"fmt"
	"sync/atomic"
	"time"
)

type RW struct {
	output chan interface{}
	input  chan interface{}
}

func NewRW() *RW {
	return &RW{
		output: make(chan interface{}),
		input:  make(chan interface{}),
	}
}

func NewRWBuffered() *RW {
	return &RW{
		output: make(chan interface{}, 100000),
		input:  make(chan interface{}, 100000),
	}
}

func (r *RW) Output(channels ...<-chan interface{}) <-chan interface{} {
	for id, channel := range channels {
		go func(ch <-chan interface{}, id int) {
			for {
				d, ok := <-ch
				if ok {
					// fmt.Println("RW.Input from", id, ":", d)
					r.output <- d
				} else {
					// fmt.Println("Reconnect", id)
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

func (r *RW) Input(channels ...chan<- interface{}) chan<- interface{} {
	go func() {
		var inc uint64

		go func() {
			var (
				last uint64
				tmp  uint64
			)

			for {
				<-time.After(1 * time.Second)
				tmp = atomic.LoadUint64(&inc)
				fmt.Println(tmp-last, "msgs/s")
				last = tmp
			}
		}()

		for {
			<-r.input
			atomic.AddUint64(&inc, 1)
			// for i := range channels {
			// 	channels[i] <- v
			// }
		}
	}()
	return r.input
}
