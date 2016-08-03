package main

import (
	"fmt"

	"golang.org/x/sync/errgroup"
)

func main() {
	var g errgroup.Group

	for i := 0; i < 10; i++ {
		i := i // closure tricks
		g.Go(func() (err error) {
			fmt.Println("->", i)
			if i == 0 {
				err = fmt.Errorf("Sorry")
			}
			return
		})
	}
	if err := g.Wait(); err != nil {
		fmt.Println(err)
	}
}
