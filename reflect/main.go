package main

import (
	"fmt"
	"reflect"

	"github.com/Sirupsen/logrus"
)

type RefectMe struct {
	A string
	B uint32
	c map[int]string
	t string `foo:"bar"`
}

func main() {
	e := reflect.TypeOf(&RefectMe{}).Elem()

	for i := 0; i < e.NumField(); i++ {
		field := e.Field(i)
		fmt.Printf(`Name: %v
PkgPath: %v
Type: %v
Tag: %v
Offset: %v
Index: %v

`,
			field.Name, field.PkgPath, field.Type, field.Tag, field.Offset, field.Index)
	}

	if !reflect.ValueOf(&RefectMe{}).Type().Implements(reflect.TypeOf((*fmt.Stringer)(nil)).Elem()) {
		logrus.Warn("Stringer interface is not implemented")
	}
}
