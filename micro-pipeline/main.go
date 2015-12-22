package main

import "fmt"

func main() {
	for d := range NewRW().Output(NewReader().Output(nil), NewWriter().Output(nil)) {
		fmt.Println("Main", d)
	}
}
