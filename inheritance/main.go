package main

import "fmt"

type Person interface {
	Init()
	PrintName()
}

type base struct {
	Name string
}

func (b *base) Init() {
	panic("You can't initialize a Base directly")
}

func (b *base) PrintName() {
	panic("You can't do it")
}

type Peon struct {
	base
}

func (p *Peon) Init() {
	p.Name = "Peon"
}

func (p *Peon) PrintName() {
	fmt.Println(p.Name)
}

type Soldier struct {
	base
}

func (s *Soldier) Init() {
	s.Name = "Soldier"
}

func (s *Soldier) PrintName() {
	fmt.Println(s.Name)
}

func main() {
	var people []Person

	people = append(people, &Peon{})
	people = append(people, &Soldier{})
	for i := range people {
		people[i].Init()
	}
	for i := range people {
		people[i].PrintName()
	}
}
