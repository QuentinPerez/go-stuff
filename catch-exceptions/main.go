package main

import (
	"fmt"
	"math/rand"
	"time"
)

func catchException(exception chan<- interface{}) {
	defer func(exc chan<- interface{}) {
		if err := recover(); err != nil {
			exc <- err
		}
	}(exception)
	if rand.Int()%2 == 0 {
		panic("Exception !!!!")
	}
}

func main() {
	exception := make(chan interface{}, 1)

	rand.Seed(time.Now().UTC().UnixNano())
	for catch := 0; catch < 5; {
		catchException(exception)
		select {
		case err := <-exception:
			fmt.Println("Exception catched ", catch, err)
			catch++
		default:
			fmt.Println(`No exception \o/`)
		}
	}
	fmt.Println("Too many panics")
}
