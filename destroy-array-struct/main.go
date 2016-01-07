package main

import "fmt"

type fake struct {
	id int
}

type ArrayFake []fake

func (a *ArrayFake) Destroy() {
REDO:
	for i := range *a {
		(*a)[i] = (*a)[len(*a)-1]
		*a = (*a)[:len(*a)-1]
		goto REDO
	}
}

func (a *ArrayFake) DestroySecond() {
	(*a)[1] = (*a)[len(*a)-1]
	*a = (*a)[:len(*a)-1]
}

func (a *ArrayFake) DestroyFirst() {
	(*a)[0] = (*a)[len(*a)-1]
	*a = (*a)[:len(*a)-1]
}

func (a *ArrayFake) DestroyFour() {
	(*a)[3] = (*a)[len(*a)-1]
	*a = (*a)[:len(*a)-1]
}

func main() {
	array := ArrayFake{}

	fmt.Println("Before:", len(array), array)
	array.Destroy()
	fmt.Println("After:", len(array), array)

	array = append(array, fake{id: 0})
	fmt.Println("Before:", len(array), array)
	array.Destroy()
	fmt.Println("After:", len(array), array)

	array = append(array, fake{id: 0})
	array = append(array, fake{id: 1})
	fmt.Println("Before:", len(array), array)
	array.Destroy()
	fmt.Println("After:", len(array), array)

	array = append(array, fake{id: 0})
	array = append(array, fake{id: 1})
	array = append(array, fake{id: 2})
	fmt.Println("Before:", len(array), array)
	array.Destroy()
	fmt.Println("After:", len(array), array)

	array = append(array, fake{id: 0})
	array = append(array, fake{id: 1})
	array = append(array, fake{id: 2})
	array = append(array, fake{id: 3})
	fmt.Println("Before:", len(array), array)
	array.Destroy()
	fmt.Println("After:", len(array), array)

	array = append(array, fake{id: 0})
	array = append(array, fake{id: 1})
	fmt.Println("Before:", len(array), array)
	array.DestroySecond()
	fmt.Println("After:", len(array), array)

	fmt.Println("Before:", len(array), array)
	array.DestroyFirst()
	fmt.Println("After:", len(array), array)

	array = append(array, fake{id: 0})
	array = append(array, fake{id: 1})
	array = append(array, fake{id: 2})
	array = append(array, fake{id: 3})
	fmt.Println("Before:", len(array), array)
	array.DestroyFour()
	fmt.Println("After:", len(array), array)
}
