package main

import "fmt"

type PetType int

const (
	Dog PetType = iota
	Cat
)

type Pet interface {
	String() string
}

type cat struct {
}

func (c cat) String() string {
	return "cat"
}

type dog struct {
}

func (d dog) String() string {
	return "dog"
}

func Factory(what PetType) Pet {
	switch what {
	case Dog:
		return &dog{}
	case Cat:
		return &cat{}
	default:
		return nil
	}
}

func main() {
	fmt.Println(Factory(Dog))
	fmt.Println(Factory(Cat))
}
