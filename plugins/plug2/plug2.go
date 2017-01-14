package main

import "fmt"

type Foo struct{}

func (f *Foo) Start() error {
	fmt.Println("coucou 2")
	return nil
}

func (f *Foo) Stop() error {
	fmt.Println("coucou 2")
	return nil
}

func NewFoo() interface{} {
	return &Foo{}
}
