package main

import "C"
import (
	"fmt"
	"log"

	"github.com/rainycape/dl"
)

var installed bool
var origRand func() *C.int

//export rand
func rand() *C.int {
	if installed {
		return origRand()
	}
	installed = true
	lib, err := dl.Open("libc", 0)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Hooked :) from golang")
	defer lib.Close()
	lib.Sym("rand", &origRand)
	return origRand()
}

func main() {
	// We need the main function to make possible
	// CGO compiler to compile the package as C shared library
}
