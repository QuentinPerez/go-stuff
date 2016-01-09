package main

import (
	"fmt"
	"unsafe"
)

type Empty struct{}

func (e Empty) String() string {
	return "Empty string"
}

func main() {
	var e Empty
	var tab [10]Empty

	fmt.Println("string():", e)
	fmt.Println("empty == struct{}{}:", e == struct{}{})
	fmt.Println("empty.sizeof:", unsafe.Sizeof(e))
	fmt.Println("tab.sizeof:", unsafe.Sizeof(tab))
}
