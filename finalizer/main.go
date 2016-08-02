package main

import (
	"fmt"
	"runtime"
)

type foo struct {
	name string
}

func pScope(w chan struct{}) {
	pFoo := &foo{"pScope"}

	fmt.Println("  SetFinalizer")
	runtime.SetFinalizer(pFoo, func(p *foo) { fmt.Println("finalizer ", p.name); close(w) })
	fmt.Println("   Force GC")
	runtime.GC() // force GC
	// <-w           // all goroutines are asleep
	pFoo.name = "10" // without this line the function SetFinalizer is called, and we can uncomment the line above
	fmt.Println("   Exit pScope")
	fmt.Println("   Release the reference on pFoo")
}

func main() {
	runtime.GOMAXPROCS(1)
	wait := make(chan struct{})
	fmt.Println("--Call pScope")
	pScope(wait)
	fmt.Println("--Force GC")
	runtime.GC() // force GC
	<-wait
	fmt.Println("--Exit")

}
