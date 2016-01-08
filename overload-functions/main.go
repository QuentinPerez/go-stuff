package main

import (
	"fmt"
	"runtime"
	"unsafe"
)

type Parent struct {
}

func (p Parent) Overload() {
	pc, _, _, _ := runtime.Caller(0)
	fmt.Println("  ->", runtime.FuncForPC(pc).Name())
	pc, _, _, _ = runtime.Caller(1)
	fmt.Println("    ->", runtime.FuncForPC(pc).Name())
}

type Child struct {
	Parent
}

func (c Child) Overload() {
	pc, _, _, _ := runtime.Caller(0)
	fmt.Println("  ->", runtime.FuncForPC(pc).Name())
}

func (c Child) CallAOverloadUnsafe() {
	(*Parent)(unsafe.Pointer(&c)).Overload()
}

func (c Child) CallAOverloadMember() {
	c.Parent.Overload()
}

func main() {
	var p Parent
	var c Child

	fmt.Println("Call parent.Overload()")
	p.Overload()
	fmt.Println()
	fmt.Println("Call child.Overload()")
	c.Overload()
	fmt.Println()
	fmt.Println("Call child.Parent.Overload()")
	c.CallAOverloadMember()
	fmt.Println("Call child.Parent.Overload() unsafe")
	c.CallAOverloadUnsafe()
}
