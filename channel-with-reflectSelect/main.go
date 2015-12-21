package main

import (
	"fmt"
	"reflect"
)

const nbOfSalves = 1000

func event(c chan<- int) {
	i := 0

	for {
		c <- i
		i++
		if i == 10 {
			close(c)
			return
		}
	}
}

func main() {
	channels := make([]chan int, nbOfSalves)
	cases := make([]reflect.SelectCase, len(channels))

	for i := range channels {
		channels[i] = make(chan int)
		cases[i] = reflect.SelectCase{
			Dir:  reflect.SelectRecv,
			Chan: reflect.ValueOf(channels[i]),
		}
		go event(channels[i])
	}
	for nb := nbOfSalves; nb > 0; {
		id, v, ok := reflect.Select(cases)
		if ok {
			fmt.Printf("%v <- %v\n", id, v.Interface().(int))
		} else { // the channel has been closed
			fmt.Printf("Done %v \n", id)
			cases[id].Chan = reflect.ValueOf(nil)
			nb--
		}
	}
}
