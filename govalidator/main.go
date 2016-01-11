package main

import (
	"fmt"

	"github.com/asaskevich/govalidator"
)

type Foo struct {
	Name   string `valid:"alpha"`
	Custom string `valid:"custom"`
}

func main() {
	f := &Foo{
		Name:   "bar",
		Custom: "custom",
	}
	govalidator.TagMap["custom"] = govalidator.Validator(func(str string) bool {
		fmt.Println("Custom call")
		return str == "custom"
	})
	if _, err := govalidator.ValidateStruct(f); err != nil {
		panic(err)
	}
	f.Custom = "fail"
	if _, err := govalidator.ValidateStruct(f); err != nil {
		panic(err)
	}
}
