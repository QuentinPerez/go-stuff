package main

import (
	"fmt"
	"reflect"
)

type A struct {
	A1 int
	A2 string
}

type B struct {
	B1 struct{}
	B2 []byte
}

type C struct {
	B1 int
	B2 int
}

type D struct {
	_ int
}

func StructEqual(src interface{}, dest interface{}) bool {
	elemSrc := reflect.TypeOf(src).Elem()
	elemDest := reflect.TypeOf(dest).Elem()
	if elemSrc.NumField() != elemDest.NumField() {
		return false
	}
	for i := 0; i < elemSrc.NumField(); i++ {
		fieldSrc := elemSrc.Field(i)
		fieldDest := elemDest.Field(i)
		if fieldDest.Name != fieldSrc.Name {
			return false
		}
		if fieldDest.PkgPath != fieldSrc.PkgPath {
			return false
		}
		if fieldDest.Type != fieldSrc.Type {
			return false
		}
		if fieldDest.Tag != fieldSrc.Tag {
			return false
		}
		if fieldDest.Offset != fieldSrc.Offset {
			return false
		}
		if len(fieldDest.Index) != len(fieldSrc.Index) {
			return false
		}
		for i := range fieldDest.Index {
			if fieldDest.Index[i] != fieldSrc.Index[i] {
				return false
			}
		}
		if fieldDest.Anonymous != fieldSrc.Anonymous {
			return false
		}
	}
	return true
}

func main() {
	var ia interface{}

	ia = &A{}
	if StructEqual(ia, &A{}) {
		fmt.Println("A == A")
	} else {
		fmt.Println("A != A")
	}
	if StructEqual(ia, &B{}) {
		fmt.Println("A == B")
	} else {
		fmt.Println("A != B")
	}
	if StructEqual(ia, &C{}) {
		fmt.Println("A == C")
	} else {
		fmt.Println("A != C")
	}
	if StructEqual(ia, &D{}) {
		fmt.Println("A == D")
	} else {
		fmt.Println("A != D")
	}
	fmt.Println()

	var ib interface{}

	ib = &B{}
	if StructEqual(ib, &A{}) {
		fmt.Println("B == A")
	} else {
		fmt.Println("B != A")
	}
	if StructEqual(ib, &B{}) {
		fmt.Println("B == B")
	} else {
		fmt.Println("B != B")
	}
	if StructEqual(ib, &C{}) {
		fmt.Println("B == C")
	} else {
		fmt.Println("B != C")
	}
	if StructEqual(ib, &D{}) {
		fmt.Println("B == D")
	} else {
		fmt.Println("B != D")
	}
	fmt.Println()

	var ic interface{}

	ic = &C{}
	if StructEqual(ic, &A{}) {
		fmt.Println("C == A")
	} else {
		fmt.Println("C != A")
	}
	if StructEqual(ic, &B{}) {
		fmt.Println("C == B")
	} else {
		fmt.Println("B != B")
	}
	if StructEqual(ic, &C{}) {
		fmt.Println("C == C")
	} else {
		fmt.Println("C != C")
	}
	if StructEqual(ic, &D{}) {
		fmt.Println("C == D")
	} else {
		fmt.Println("C != D")
	}
	fmt.Println()

	var id interface{}

	id = &D{}
	if StructEqual(id, &A{}) {
		fmt.Println("D == A")
	} else {
		fmt.Println("D != A")
	}
	if StructEqual(id, &B{}) {
		fmt.Println("D == B")
	} else {
		fmt.Println("D != B")
	}
	if StructEqual(id, &C{}) {
		fmt.Println("D == C")
	} else {
		fmt.Println("D != C")
	}
	if StructEqual(id, &D{}) {
		fmt.Println("D == D")
	} else {
		fmt.Println("D != D")
	}
	fmt.Println()
}
