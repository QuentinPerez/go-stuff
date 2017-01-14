package main

import (
	"fmt"
	"plugin"
)

type MyInterface interface {
	Start() error
	Stop() error
}

func main() {
	aInterface := make([]MyInterface, 0, 3)

	for i := 1; i < 4; i++ {
		load := fmt.Sprintf("plug%d/plug%d.so", i, i)
		fmt.Println(load)
		so, err := plugin.Open(load)
		if err != nil {
			panic(err)
		}
		fn, err := so.Lookup("NewFoo")
		if err != nil {
			panic(err)
		}
		aInterface = append(aInterface, fn.(func() interface{})().(MyInterface))
	}
	for _, interf := range aInterface {
		interf.Start()
		interf.Stop()
	}
}
