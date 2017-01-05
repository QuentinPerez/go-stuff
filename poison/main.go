package main

import (
	"log"
	"math/rand"
	"sync"
	"time"
)

func openPoison(fname string, poison chan bool) {
	defer func() {
		if r := recover(); r != nil {
			log.Print("simultanous close")
		}
	}()
	log.Print("Open poison: ", fname)
	close(poison)
}

func foo(poison chan bool, wait *sync.WaitGroup) {
	defer wait.Done()
	for {
		select {
		case <-poison:
			log.Println("close foo")
			return
		case <-time.After(time.Duration(rand.Int31n(300)) * time.Microsecond):
			openPoison("foo", poison)
		}
	}
}

func bar(poison chan bool, wait *sync.WaitGroup) {
	defer wait.Done()
	for {
		select {
		case <-poison:
			log.Println("close bar")
			return
		case <-time.After(time.Duration(rand.Int31n(300)) * time.Microsecond):
			openPoison("bar", poison)
		}
	}
}

func main() {
	poison := make(chan bool)
	wg := sync.WaitGroup{}

	wg.Add(1)
	go foo(poison, &wg)
	wg.Add(1)
	go bar(poison, &wg)
	wg.Wait()
}
