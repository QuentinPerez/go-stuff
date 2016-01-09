package main

import (
	"fmt"
	"reflect"
	"unsafe"
)

type Empty struct{}

func (e Empty) String() string {
	val := reflect.Indirect(reflect.ValueOf(e))
	return val.Type().String()
}

type MakeStringer struct {
	Empty
}

func main() {
	var e Empty
	var tab [10]Empty
	var m MakeStringer

	fmt.Println("string():", e)
	fmt.Println("empty == struct{}{}:", e == struct{}{})
	fmt.Println("empty.sizeof:", unsafe.Sizeof(e))
	fmt.Println("tab.sizeof:", unsafe.Sizeof(tab))
	fmt.Println(m)
}
