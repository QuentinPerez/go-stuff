package main

import "reflect"

//
// go version go1.5.2 darwin/amd64
// 4 cores
// 32 Go RAM
//

// TODO use go bench

// 65535 is the limit that reflect.Select can handle.
const nbOfSalves = 10000

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
			// performances:
			// go run main.go  814.63s user 19.88s system 173% cpu 8:00.19 total
			//
			// cases[id].Chan = reflect.ValueOf(nil)

			// 2) remove entry
			// performances:
			// go run main.go  750.54s user 16.91s system 175% cpu 7:16.18 total
			//
			cases = cases[:id+copy(cases[id:], cases[id+1:])]
			nb--
		}
	}
}
