package reflectSelect

import (
	"log"
	"reflect"
)

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

func Start(nbOfSalves int) {
	if nbOfSalves > 65535 {
		log.Fatal("reflect.Select can't handle more than 65535 descriptor")
	}
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
		// id, v, ok := reflect.Select(cases)
		id, _, ok := reflect.Select(cases)
		if ok {
			// fmt.Printf("%v <- %v\n", id, v.Interface().(int))
		} else { // the channel has been closed
			// fmt.Printf("Done\n")
			//
			// disable the channel
			//
			// 1) with nil
			// cases[id].Chan = reflect.ValueOf(nil)
			// nb--

			// 2) remove entry
			cases = cases[:id+copy(cases[id:], cases[id+1:])]
			nb = len(cases)
		}
	}
}
